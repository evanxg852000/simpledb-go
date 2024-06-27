package record

import (
	"log"

	"github.com/evanxg852000/simpledb/internal/file"
	"github.com/evanxg852000/simpledb/internal/tx/recovery"
)

const (
	EMPTY = iota
	USED
)

// Store a record at a given location in a block.
type RecordPage struct {
	tx      *recovery.Transaction
	blockId file.BlockId
	layout  *Layout
}

func NewRecordPage(tx *recovery.Transaction, blockId file.BlockId, layout *Layout) *RecordPage {
	tx.Pin(blockId)
	return &RecordPage{tx, blockId, layout}
}

// Return the integer value stored for the
// specified field of a specified slot.
func (rp *RecordPage) GetInt(slot int64, fldName string) (int64, error) {
	fldPosition := rp.offset(slot) + rp.layout.Offset(fldName)
	return rp.tx.GetInt(rp.blockId, fldPosition)
}

// Return the string value stored for the
// specified field of the specified slot.
func (rp *RecordPage) GetString(slot int64, fldName string) (string, error) {
	fldPosition := rp.offset(slot) + rp.layout.Offset(fldName)
	return rp.tx.GetString(rp.blockId, fldPosition)
}

// Store an integer at the specified field
// of the specified slot.
func (rp *RecordPage) SetInt(slot int64, fldName string, value int64) error {
	fldPosition := rp.offset(slot) + rp.layout.Offset(fldName)
	return rp.tx.SetInt(rp.blockId, fldPosition, value, true)
}

// Store a string at the specified field
// of the specified slot.
func (rp *RecordPage) SetString(slot int64, fldName string, value string) error {
	fldPosition := rp.offset(slot) + rp.layout.Offset(fldName)
	return rp.tx.SetString(rp.blockId, fldPosition, value, true)
}

func (rp *RecordPage) Delete(slot int64) {
	rp.setFlag(slot, EMPTY)
}

// Use the layout to format a new block of records.
// These values should not be logged
// (because the old values are meaningless).
func (rp *RecordPage) Format() error {
	slot := int64(0)
	for rp.isValidSlot(slot) {
		rp.tx.SetInt(rp.blockId, rp.offset(slot), EMPTY, false)
		schema := rp.layout.Schema
		for _, fldName := range schema.Fields() {
			fldPosition := rp.offset(slot) + rp.layout.Offset(fldName)
			if schema.FieldType(fldName) == INTEGER_TYPE {
				err := rp.tx.SetInt(rp.blockId, fldPosition, 0, false)
				if err != nil {
					return err
				}
			} else {
				err := rp.tx.SetString(rp.blockId, fldPosition, "", false)
				if err != nil {
					return err
				}
			}
		}
		slot += 1
	}
	return nil
}

func (rp *RecordPage) NextAfter(slot int64) int64 {
	return rp.searchAfter(slot, USED)
}

func (rp *RecordPage) InsertAfter(slot int64) int64 {
	newSlot := rp.searchAfter(slot, EMPTY)
	if newSlot >= 0 {
		rp.setFlag(newSlot, USED)
	}
	return newSlot
}

func (rp *RecordPage) BlockId() file.BlockId {
	return rp.blockId
}

func (rp *RecordPage) setFlag(slot int64, flag int64) {
	rp.tx.SetInt(rp.blockId, rp.offset(slot), flag, true)
}

func (rp *RecordPage) searchAfter(slot int64, flag int64) int64 {
	slot += 1
	for rp.isValidSlot(slot) {
		storedFlag, err := rp.tx.GetInt(rp.blockId, rp.offset(slot))
		if err != nil {
			log.Fatalf("error check slot status: %v", err)
		}
		if storedFlag == flag {
			return slot
		}
		slot += 1
	}
	return -1
}

func (rp *RecordPage) isValidSlot(slot int64) bool {
	return rp.offset(slot+1) <= rp.tx.BlockSize()
}

func (rp *RecordPage) offset(slot int64) int64 {
	return slot * rp.layout.slotSize
}
