package buffer

import (
	"github.com/evanxg852000/simpledb/internal/file"
	walog "github.com/evanxg852000/simpledb/internal/log"
)

// A data buffer that wraps a page ans stores information about its status.
// such as the associated disk block, the number of times the buffer has been
// pinned, whether its contents has been modified, and if so, the id and lsn
// of the modifying transaction.
type Buffer struct {
	fileManager *file.FileManager
	logManager  *walog.LogManager
	content     file.Page
	blockId     file.BlockId
	pins        int64
	txNum       int64
	lsn         int64
}

func NewBuffer(fileManager *file.FileManager, logManager *walog.LogManager) Buffer {
	return Buffer{
		fileManager: fileManager,
		logManager:  logManager,
		content:     file.NewPage(fileManager.BlockSize()),
		blockId:     file.BlockId{},
		pins:        0,
		txNum:       -1,
		lsn:         -1,
	}
}

func (buf *Buffer) Content() *file.Page {
	return &buf.content
}

func (buf *Buffer) Block() file.BlockId {
	return buf.blockId
}

// Negative lsn denotes a transaction withouts corresponding
// log record
func (buf *Buffer) Modify(txNum int64, lsn int64) {
	buf.txNum = txNum
	if lsn >= 0 {
		buf.lsn = lsn
	}
}

func (buf *Buffer) IsPinned() bool {
	return buf.pins > 0
}

func (buf *Buffer) ModifyingTx() int64 {
	return buf.txNum
}

// Reads the content of the specified block into the content of the buffer.
// if the buffer was dirty, then its previous content
// is first written to disk.
func (buf *Buffer) AssignToBlock(blockId file.BlockId) error {
	buf.flush()
	buf.blockId = blockId
	err := buf.fileManager.Read(blockId, &buf.content)
	if err != nil {
		return err
	}
	buf.pins = 0
	return nil
}

// Write the buffer content to its disk block if it is dirty
func (buf *Buffer) flush() error {
	if buf.txNum > 0 {
		err := buf.logManager.Flush(buf.lsn)
		if err != nil {
			return err
		}

		err = buf.fileManager.Write(buf.blockId, &buf.content)
		if err != nil {
			return err
		}
		buf.txNum = -1
	}
	return nil
}

func (buf *Buffer) Pin() {
	buf.pins += 1
}

func (buf *Buffer) Unpin() {
	buf.pins -= 1
}
