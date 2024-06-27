package recovery

import (
	"github.com/evanxg852000/simpledb/internal/file"
	walog "github.com/evanxg852000/simpledb/internal/log"
)

type CheckpointRecord struct {
}

func NewCheckpointRecord() (CheckpointRecord, error) {
	return CheckpointRecord{}, nil
}

func (cr CheckpointRecord) Operation() int {
	return CHECKPOINT
}

func (cr CheckpointRecord) TxNumber() int64 {
	return -1
}

func (cr CheckpointRecord) ToString() string {
	return "<CHECKPOINT>"
}

// Does nothing, because a checkpoint record
// contains no undo information.
func (cr CheckpointRecord) Undo(tx *Transaction) {}

// A static method to write a checkpoint record to the log.
// This log record contains the CHECKPOINT operator,
// and nothing else.
func (cr CheckpointRecord) WriteToLog(lm *walog.LogManager) (int64, error) {
	buffer := file.NewByteBuffer()
	err := buffer.WriteInt(CHECKPOINT)
	if err != nil {
		return -1, err
	}

	return lm.Append(buffer.Data())
}
