package recovery

import (
	"slices"

	"github.com/evanxg852000/simpledb/internal/buffer"
	walog "github.com/evanxg852000/simpledb/internal/log"
)

// Each transaction has its own recovery manager.
type RecoveryManager struct {
	logManager    *walog.LogManager
	bufferManager *buffer.BufferManager
	tx            *Transaction
	txNum         int64
}

// Create a recovery manager for the specified transaction.
func NewRecoveryManager(tx *Transaction, txNum int64, lm *walog.LogManager, bm *buffer.BufferManager) *RecoveryManager {
	StartRecord.WriteToLog(StartRecord{}, lm, txNum)
	return &RecoveryManager{
		tx:            tx,
		txNum:         txNum,
		logManager:    lm,
		bufferManager: bm,
	}
}

// Write a commit record to the log, and flushes it to disk.
func (rm *RecoveryManager) Commit() error {
	err := rm.bufferManager.FlushAll(rm.txNum)
	if err != nil {
		return err
	}
	lsn, err := CommitRecord.WriteToLog(CommitRecord{}, rm.logManager, rm.txNum)
	if err != nil {
		return err
	}
	return rm.logManager.Flush(lsn)
}

// Write a rollback record to the log and flush it to disk.
func (rm *RecoveryManager) Rollback() error {
	err := rm.doRollback()
	if err != nil {
		return err
	}
	rm.bufferManager.FlushAll(rm.txNum)
	lsn, err := RollbackRecord.WriteToLog(RollbackRecord{}, rm.logManager, rm.txNum)
	if err != nil {
		return err
	}
	rm.logManager.Flush(lsn)
	return nil
}

// Recover uncompleted transactions from the log
// and then write a quiescent checkpoint record to the log and flush it.
func (rm *RecoveryManager) Recover() error {
	err := rm.doRecover()
	if err != nil {
		return err
	}
	rm.bufferManager.FlushAll(rm.txNum)
	lsn, err := CheckpointRecord.WriteToLog(CheckpointRecord{}, rm.logManager)
	if err != nil {
		return err
	}
	rm.logManager.Flush(lsn)
	return nil
}

// Write a setint record to the log and return its lsn.
func (rm *RecoveryManager) SetInt(buffer *buffer.Buffer, offset int64, value int64) (int64, error) {
	oldValue, err := buffer.Content().ReadInt(offset)
	if err != nil {
		return 0, err
	}
	blockId := buffer.Block()
	return SetIntRecord.WriteToLog(SetIntRecord{}, rm.logManager, rm.txNum, blockId, offset, oldValue)
}

// Write a setstring record to the log and return its lsn.
func (rm *RecoveryManager) SetString(buffer *buffer.Buffer, offset int64, value string) (int64, error) {
	oldValue, err := buffer.Content().ReadString(offset)
	if err != nil {
		return 0, err
	}
	blockId := buffer.Block()
	return SetStringRecord.WriteToLog(SetStringRecord{}, rm.logManager, rm.txNum, blockId, offset, oldValue)
}

// Rollback the transaction, by iterating
// through the log records until it finds
// the transaction's START record,
// calling undo() for each of the transaction's
// log records.
func (rm *RecoveryManager) doRollback() error {
	iter, err := rm.logManager.Iterator()
	if err != nil {
		return err
	}

	for iter.HasNext() {
		data, err := iter.Next()
		if err != nil {
			return err
		}
		record, err := NewLogRecord(data)
		if err != nil {
			return err
		}

		if record.TxNumber() == rm.txNum {
			if record.Operation() == START {
				return nil
			}
			record.Undo(rm.tx)
		}

	}
	return nil
}

// Do a complete database recovery.
// The method iterates through the log records.
// Whenever it finds a log record for an unfinished
// transaction, it calls undo() on that record.
// The method stops when it encounters a CHECKPOINT record
// or the end of the log.
func (rm *RecoveryManager) doRecover() error {
	finishedTxs := make([]int64, 0)
	iter, err := rm.logManager.Iterator()
	if err != nil {
		return err
	}

	for iter.HasNext() {
		data, err := iter.Next()
		if err != nil {
			return err
		}
		record, err := NewLogRecord(data)
		if err != nil {
			return err
		}

		if record.Operation() == CHECKPOINT {
			return nil
		}

		if record.Operation() == COMMIT || record.Operation() == ROLLBACK {
			finishedTxs = append(finishedTxs, record.TxNumber())
		}

		if !slices.Contains(finishedTxs, record.TxNumber()) {
			record.Undo(rm.tx)
		}
	}
	return nil
}
