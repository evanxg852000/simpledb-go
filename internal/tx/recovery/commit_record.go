package recovery

import (
	"fmt"

	"github.com/evanxg852000/simpledb/internal/file"
	walog "github.com/evanxg852000/simpledb/internal/log"
)

type CommitRecord struct {
	txNum int64
}

func NewCommitRecord(page file.Page) (CommitRecord, error) {
	txNum, err := page.ReadInt(8)
	if err != nil {
		return CommitRecord{}, err
	}
	return CommitRecord{txNum}, nil
}

func (cr CommitRecord) Operation() int {
	return COMMIT
}

func (cr CommitRecord) TxNumber() int64 {
	return cr.txNum
}

func (cr CommitRecord) ToString() string {
	return fmt.Sprintf("<COMMIT %d>", cr.txNum)
}

// Does nothing, because a commit record
// contains no undo information.
func (cr CommitRecord) Undo(tx *Transaction) {}

// A static method to write a commit record to the log.
// This log record contains the COMMIT operator,
// followed by the transaction id.
func (cr CommitRecord) WriteToLog(lm *walog.LogManager, txNum int64) (int64, error) {
	buffer := file.NewByteBuffer()
	err := buffer.WriteInt(COMMIT)
	if err != nil {
		return -1, err
	}

	err = buffer.WriteInt(txNum)
	if err != nil {
		return -1, err
	}

	return lm.Append(buffer.Data())
}
