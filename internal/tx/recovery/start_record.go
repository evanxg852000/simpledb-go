package recovery

import (
	"fmt"

	"github.com/evanxg852000/simpledb/internal/file"
	walog "github.com/evanxg852000/simpledb/internal/log"
)

type StartRecord struct {
	txNum int64
}

func NewStartRecord(page file.Page) (StartRecord, error) {
	txNum, err := page.ReadInt(8)
	if err != nil {
		return StartRecord{}, err
	}
	return StartRecord{txNum}, nil
}

func (sr StartRecord) Operation() int {
	return START
}

func (sr StartRecord) TxNumber() int64 {
	return sr.txNum
}

func (sr StartRecord) ToString() string {
	return fmt.Sprintf("<START %d>", sr.txNum)
}

func (sr StartRecord) Undo(tx *Transaction) {}

// A static method to write a start record to the log.
// This log record contains the START operator,
// followed by the transaction id.
func (sr StartRecord) WriteToLog(lm *walog.LogManager, txNum int64) (int64, error) {
	buffer := file.NewByteBuffer()
	err := buffer.WriteInt(START)
	if err != nil {
		return -1, err
	}

	err = buffer.WriteInt(txNum)
	if err != nil {
		return -1, err
	}

	return lm.Append(buffer.Data())
}
