package recovery

import (
	"fmt"

	"github.com/evanxg852000/simpledb/internal/file"
	walog "github.com/evanxg852000/simpledb/internal/log"
)

type SetIntRecord struct {
	txNum   int64
	offset  int64
	value   int64
	blockId file.BlockId
}

func NewSetIntRecord(page file.Page) (SetIntRecord, error) {
	pos := int64(8)
	txNum, err := page.ReadInt(pos)
	if err != nil {
		return SetIntRecord{}, err
	}

	pos += 8
	fileName, err := page.ReadString(pos)
	if err != nil {
		return SetIntRecord{}, err
	}

	pos = pos + file.GetEncodingLength(int64(len(fileName)))
	blockNum, err := page.ReadInt(pos)
	if err != nil {
		return SetIntRecord{}, err
	}
	blockId := file.NewBlockId(fileName, blockNum)

	pos = pos + 8
	offset, err := page.ReadInt(pos)
	if err != nil {
		return SetIntRecord{}, err
	}

	pos = pos + 8
	value, err := page.ReadInt(pos)
	if err != nil {
		return SetIntRecord{}, err
	}

	return SetIntRecord{txNum, offset, value, blockId}, nil
}

func (sir SetIntRecord) Operation() int {
	return SETINT
}

func (sir SetIntRecord) TxNumber() int64 {
	return sir.txNum
}

func (sir SetIntRecord) ToString() string {
	return fmt.Sprintf("<SETINT %d %v %d %d>", sir.txNum, sir.blockId, sir.offset, sir.value)
}

// Replace the specified data value with the value saved in the log record.
// The method pins a buffer to the specified block,
// calls setInt to restore the saved value,
// and unpins the buffer.
func (sir SetIntRecord) Undo(tx *Transaction) {
	tx.Pin(sir.blockId)
	tx.SetInt(sir.blockId, sir.offset, sir.value, false) // don't log the undo
	tx.Unpin(sir.blockId)
}

// A static method to write a setInt record to the log.
// This log record contains the SETINT operator,
// followed by the transaction id, the filename, number,
// and offset of the modified block, and the previous
// integer value at that offset.
func (SetIntRecord) WriteToLog(lm *walog.LogManager, txNum int64, blockId file.BlockId, offset int64, value int64) (int64, error) {
	buffer := file.NewByteBuffer()
	err := buffer.WriteInt(SETINT)
	if err != nil {
		return -1, err
	}

	err = buffer.WriteInt(txNum)
	if err != nil {
		return -1, err
	}

	err = buffer.WriteString(blockId.FileName)
	if err != nil {
		return -1, err
	}

	err = buffer.WriteInt(blockId.BlockNum)
	if err != nil {
		return -1, err
	}

	err = buffer.WriteInt(offset)
	if err != nil {
		return -1, err
	}

	err = buffer.WriteInt(value)
	if err != nil {
		return -1, err
	}

	return lm.Append(buffer.Data())
}
