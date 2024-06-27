package recovery

import (
	"fmt"

	"github.com/evanxg852000/simpledb/internal/file"
	walog "github.com/evanxg852000/simpledb/internal/log"
)

type SetStringRecord struct {
	txNum   int64
	offset  int64
	value   string
	blockId file.BlockId
}

func NewSetStringRecord(page file.Page) (SetStringRecord, error) {
	pos := int64(8)
	txNum, err := page.ReadInt(pos)
	if err != nil {
		return SetStringRecord{}, err
	}

	pos += 8
	fileName, err := page.ReadString(pos)
	if err != nil {
		return SetStringRecord{}, err
	}

	pos = pos + file.GetEncodingLength(int64(len(fileName)))
	blockNum, err := page.ReadInt(pos)
	if err != nil {
		return SetStringRecord{}, err
	}
	blockId := file.NewBlockId(fileName, blockNum)

	pos = pos + 8
	offset, err := page.ReadInt(pos)
	if err != nil {
		return SetStringRecord{}, err
	}

	pos = pos + 8
	value, err := page.ReadString(pos)
	if err != nil {
		return SetStringRecord{}, err
	}

	return SetStringRecord{txNum, offset, value, blockId}, nil
}

func (ssr SetStringRecord) Operation() int {
	return SETSTRING
}

func (ssr SetStringRecord) TxNumber() int64 {
	return ssr.txNum
}

func (ssr SetStringRecord) ToString() string {
	return fmt.Sprintf("<SETSTRING %d %v %d %s>", ssr.txNum, ssr.blockId, ssr.offset, ssr.value)
}

// Replace the specified data value with the value saved in the log record.
// The method pins a buffer to the specified block,
// calls setInt to restore the saved value,
// and unpins the buffer.
func (ssr SetStringRecord) Undo(tx *Transaction) {
	tx.Pin(ssr.blockId)
	tx.SetString(ssr.blockId, ssr.offset, ssr.value, false) // don't log the undo
	tx.Unpin(ssr.blockId)
}

// A static method to write a SetString record to the log.
// This log record contains the SETSTRING operator,
// followed by the transaction id, the filename, number,
// and offset of the modified block, and the previous
// integer value at that offset.
func (SetStringRecord) WriteToLog(lm *walog.LogManager, txNum int64, blockId file.BlockId, offset int64, value string) (int64, error) {
	buffer := file.NewByteBuffer()
	err := buffer.WriteInt(SETSTRING)
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

	err = buffer.WriteString(value)
	if err != nil {
		return -1, err
	}

	return lm.Append(buffer.Data())
}
