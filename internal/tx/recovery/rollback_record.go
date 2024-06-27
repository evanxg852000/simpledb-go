package recovery

import (
	"fmt"

	"github.com/evanxg852000/simpledb/internal/file"
	walog "github.com/evanxg852000/simpledb/internal/log"
)

type RollbackRecord struct {
	txNum int64
}

func NewRollbackRecord(page file.Page) (RollbackRecord, error) {
	txNum, err := page.ReadInt(8)
	if err != nil {
		return RollbackRecord{}, err
	}
	return RollbackRecord{txNum}, nil
}

func (rr RollbackRecord) Operation() int {
	return START
}

func (rr RollbackRecord) TxNumber() int64 {
	return rr.txNum
}

func (rr RollbackRecord) ToString() string {
	return fmt.Sprintf("<ROLLBACK %d>", rr.txNum)
}

// Does nothing, because a rollback record
// contains no undo information.
func (rr RollbackRecord) Undo(tx *Transaction) {}

// A static method to write a rollback record to the log.
// This log record contains the ROLLBACK operator,
// followed by the transaction id.
func (sr RollbackRecord) WriteToLog(lm *walog.LogManager, txNum int64) (int64, error) {
	buffer := file.NewByteBuffer()
	err := buffer.WriteInt(ROLLBACK)
	if err != nil {
		return -1, err
	}

	err = buffer.WriteInt(txNum)
	if err != nil {
		return -1, err
	}

	return lm.Append(buffer.Data())
}
