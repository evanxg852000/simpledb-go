package recovery

import (
	"github.com/evanxg852000/simpledb/internal/file"
)

const (
	CHECKPOINT = iota
	START
	COMMIT
	ROLLBACK
	SETINT
	SETSTRING
)

// The interface implemented by each type of log record.
type LogRecord interface {
	// Returns the log record's type.
	Operation() int

	// Returns the transaction id stored with the log record.
	TxNumber() int64

	// Undoes the operation encoded by this log record.
	// The only log record types for which this method
	// does anything interesting are SETINT and SETSTRING.
	Undo(tx *Transaction)

	// Return string representation
	ToString() string
}

func NewLogRecord(data []byte) (LogRecord, error) {
	page := file.NewPageWithData(data)
	op, err := page.ReadInt(0)
	if err != nil {
		return nil, err
	}
	switch op {
	case CHECKPOINT:
		return NewCheckpointRecord()
	case START:
		return NewStartRecord(page)
	case COMMIT:
		return NewCommitRecord(page)
	case ROLLBACK:
		return NewRollbackRecord(page)
	case SETINT:
		return NewSetIntRecord(page)
	case SETSTRING:
		return NewSetStringRecord(page)
	}
	return nil, nil
}
