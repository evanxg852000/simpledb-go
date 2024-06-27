package recovery

import (
	"fmt"
	"log"
	"sync/atomic"

	"github.com/evanxg852000/simpledb/internal/buffer"
	"github.com/evanxg852000/simpledb/internal/file"
	walog "github.com/evanxg852000/simpledb/internal/log"
	"github.com/evanxg852000/simpledb/internal/tx/concurrency"
	// "github.com/evanxg852000/simpledb/internal/tx/recovery"
)

const END_OF_FILE int = -1

var nextTxNum = atomic.Int64{}

type Transaction struct {
	recoveryManager    *RecoveryManager
	concurrencyManager *concurrency.ConcurrencyManager
	bufferManager      *buffer.BufferManager
	fileManager        *file.FileManager
	txNum              int64
	buffers            *BufferList
}

func NewTransaction(fileManager *file.FileManager, logManager *walog.LogManager, bufferManager *buffer.BufferManager) *Transaction {
	newTxNum := nextTxNum.Add(1)
	tx := &Transaction{
		concurrencyManager: concurrency.NewConcurrencyManager(),
		bufferManager:      bufferManager,
		fileManager:        fileManager,
		txNum:              newTxNum,
		buffers:            NewBufferList(bufferManager),
	}

	tx.recoveryManager = NewRecoveryManager(tx, newTxNum, logManager, bufferManager)
	return tx
}

// Commit the current transaction.
// Flush all modified buffers (and their log records),
// write and flush a commit record to the log,
// release all locks, and unpin any pinned buffers.
func (tx *Transaction) Commit() {
	tx.recoveryManager.Commit()
	log.Printf("transaction %d committed\n", tx.txNum)
	tx.concurrencyManager.Release()
	tx.buffers.UnpinAll()
}

// Rollback the current transaction.
// Undo any modified values,
// flush those buffers,
// write and flush a rollback record to the log,
// release all locks, and unpin any pinned buffers.
func (tx *Transaction) Rollback() {
	tx.recoveryManager.Rollback()
	fmt.Printf("transaction %d rolled back\n", tx.txNum)
	tx.concurrencyManager.Release()
	tx.buffers.UnpinAll()
}

// Flush all modified buffers.
// Then go through the log, rolling back all
// uncommitted transactions.  Finally,
// write a quiescent checkpoint record to the log.
// This method is called during system startup,
// before user transactions begin.
func (tx *Transaction) Recover() {
	tx.bufferManager.FlushAll(tx.txNum)
	tx.recoveryManager.Recover()
}

// Pin the specified block.
// The transaction manages the buffer for the client.
func (tx *Transaction) Pin(blockId file.BlockId) {
	tx.buffers.Pin(blockId)
}

// Unpin the specified block.
// The transaction looks up the buffer pinned to this block,
// and unpins it.
func (tx *Transaction) Unpin(blockId file.BlockId) {
	tx.buffers.Unpin(blockId)
}

// Return the integer value stored at the
// specified offset of the specified block.
// The method first obtains an SLock on the block,
// then it calls the buffer to retrieve the value.
// returns the integer stored at that offset
func (tx *Transaction) GetInt(blockId file.BlockId, offset int64) (int64, error) {
	tx.concurrencyManager.SLock(blockId)
	buffer := tx.buffers.GetBuffer(blockId)
	return buffer.Content().ReadInt(offset)
}

// Return the string value stored at the
// specified offset of the specified block.
// The method first obtains an SLock on the block,
// then it calls the buffer to retrieve the value.
// returns the string stored at that offset
func (tx *Transaction) GetString(blockId file.BlockId, offset int64) (string, error) {
	tx.concurrencyManager.SLock(blockId)
	buffer := tx.buffers.GetBuffer(blockId)
	return buffer.Content().ReadString(offset)
}

// Store an integer at the specified offset
// of the specified block.
// The method first obtains an XLock on the block.
// It then reads the current value at that offset,
// puts it into an update log record, and
// writes that record to the log.
// Finally, it calls the buffer to store the value,
// passing in the LSN of the log record and the transaction's id.
func (tx *Transaction) SetInt(blockId file.BlockId, offset int64, value int64, okToLog bool) error {
	tx.concurrencyManager.XLock(blockId)
	buffer := tx.buffers.GetBuffer(blockId)
	lsn := int64(-1)
	if okToLog {
		lsn, err := tx.recoveryManager.SetInt(buffer, offset, value)
		_ = lsn
		if err != nil {
			return err
		}
	}
	err := buffer.Content().WriteInt(offset, value)
	if err != nil {
		return err
	}
	buffer.Modify(tx.txNum, lsn)
	return nil
}

// Store a string at the specified offset
// of the specified block.
// The method first obtains an XLock on the block.
// It then reads the current value at that offset,
// puts it into an update log record, and
// writes that record to the log.
// Finally, it calls the buffer to store the value,
// passing in the LSN of the log record and the transaction's id.
func (tx *Transaction) SetString(blockId file.BlockId, offset int64, value string, okToLog bool) error {
	tx.concurrencyManager.XLock(blockId)
	buffer := tx.buffers.GetBuffer(blockId)
	lsn := int64(-1)
	if okToLog {
		lsn, err := tx.recoveryManager.SetString(buffer, offset, value)
		_ = lsn
		if err != nil {
			return err
		}
	}
	_, err := buffer.Content().WriteString(offset, value)
	if err != nil {
		return err
	}
	buffer.Modify(tx.txNum, lsn)
	return nil
}

// Return the number of blocks in the specified file.
// This method first obtains an SLock on the
// "end of the file", before asking the file manager
// to return the file size.
func (tx *Transaction) Size(fileName string) (int64, error) {
	blockId := file.NewBlockId(fileName, int64(END_OF_FILE))
	tx.concurrencyManager.SLock(blockId)
	return tx.fileManager.BlockCount(fileName)
}

// Append a new block to the end of the specified file
// and returns a reference to it.
// This method first obtains an XLock on the
// "end of the file", before performing the append.
func (tx *Transaction) Append(fileName string) (file.BlockId, error) {
	blockId := file.NewBlockId(fileName, int64(END_OF_FILE))
	tx.concurrencyManager.XLock(blockId)
	return tx.fileManager.Append(fileName)
}

func (tx *Transaction) BlockSize() int64 {
	return tx.fileManager.BlockSize()
}

func (tx *Transaction) AvailableBuffers() int64 {
	return tx.bufferManager.Available()
}
