package record

import (
	"fmt"

	"github.com/evanxg852000/simpledb/internal/file"
	"github.com/evanxg852000/simpledb/internal/query"
	"github.com/evanxg852000/simpledb/internal/tx/recovery"
)

type TableScan struct {
	tx          *recovery.Transaction
	layout      *Layout
	recordPage  *RecordPage
	fileName    string
	currentSlot int64
}

func NewTableScan(tx *recovery.Transaction, tblName string, layout *Layout) (*TableScan, error) {
	fileName := fmt.Sprintf("%s.tbl", tblName)
	tableScan := &TableScan{
		tx:          tx,
		layout:      layout,
		recordPage:  nil,
		fileName:    fileName,
		currentSlot: -1,
	}
	fileSize, err := tx.Size(fileName)
	if err != nil {
		return nil, err
	}

	if fileSize == 0 {
		err := tableScan.moveToNewBlock()
		if err != nil {
			return nil, err
		}
	} else {
		tableScan.moveToBlock(0)
	}
	return tableScan, nil
}

// Methods that implement Scan

func (tblScan *TableScan) BeforeFirst() {
	tblScan.moveToBlock(0)
}

func (tblScan *TableScan) Next() bool {
	tblScan.currentSlot = tblScan.recordPage.NextAfter(tblScan.currentSlot)
	for tblScan.currentSlot < 0 {
		isAtLast, err := tblScan.atLastBlock()
		if err != nil {
			panic(err)
		}

		if isAtLast {
			return false
		}
		tblScan.moveToBlock(tblScan.recordPage.BlockId().BlockNum + 1)
		tblScan.currentSlot = tblScan.recordPage.NextAfter(tblScan.currentSlot)
	}
	return true
}

func (tblScan *TableScan) GetInt(fldName string) int64 {
	value, err := tblScan.recordPage.GetInt(tblScan.currentSlot, fldName)
	if err != nil {
		panic(err)
	}
	return value
}

func (tblScan *TableScan) GetString(fldName string) string {
	value, err := tblScan.recordPage.GetString(tblScan.currentSlot, fldName)
	if err != nil {
		panic(err)
	}
	return value
}

func (tblScan *TableScan) GetValue(fldName string) query.Constant {
	if tblScan.layout.Schema.FieldType(fldName) == INTEGER_TYPE {
		iVal := tblScan.GetInt(fldName)
		return query.NewConstant(iVal)
	}

	sVal := tblScan.GetString(fldName)
	return query.NewConstant(sVal)
}

func (tblScan *TableScan) HasField(fldName string) bool {
	return tblScan.layout.Schema.HasField(fldName)
}

func (tblScan *TableScan) Close() {
	if tblScan.recordPage != nil {
		tblScan.tx.Unpin(tblScan.recordPage.BlockId())
	}
}

// Methods that implement UpdateScan

func (tblScan *TableScan) SetInt(fldName string, value int64) {
	err := tblScan.recordPage.SetInt(tblScan.currentSlot, fldName, value)
	if err != nil {
		panic(err)
	}
}

func (tblScan *TableScan) SetString(fldName string, value string) {
	err := tblScan.recordPage.SetString(tblScan.currentSlot, fldName, value)
	if err != nil {
		panic(err)
	}
}

func (tblScan *TableScan) SetValue(fldName string, value query.Constant) {
	if tblScan.layout.Schema.FieldType(fldName) == INTEGER_TYPE {
		tblScan.SetInt(fldName, value.AsInt())
		return
	}
	tblScan.SetString(fldName, value.AsString())
}

func (tblScan *TableScan) Insert() {
	tblScan.currentSlot = tblScan.recordPage.InsertAfter(tblScan.currentSlot)
	for tblScan.currentSlot < 0 {
		isAtLast, err := tblScan.atLastBlock()
		if err != nil {
			panic(err)
		}
		if isAtLast {
			err := tblScan.moveToNewBlock()
			if err != nil {
				panic(err)
			}
		} else {
			tblScan.moveToBlock(tblScan.recordPage.BlockId().BlockNum + 1)
		}

		tblScan.currentSlot = tblScan.recordPage.InsertAfter(tblScan.currentSlot)
	}
}

func (tblScan *TableScan) Delete() {
	tblScan.recordPage.Delete(tblScan.currentSlot)
}

func (tblScan *TableScan) MoveToRID(rId query.RID) {
	tblScan.Close()
	blockId := file.NewBlockId(tblScan.fileName, rId.BlockNum)
	tblScan.recordPage = NewRecordPage(tblScan.tx, blockId, tblScan.layout)
	tblScan.currentSlot = rId.Slot
}

func (tblScan *TableScan) GetRID() query.RID {
	return query.NewRID(tblScan.recordPage.BlockId().BlockNum, tblScan.currentSlot)
}

// Private auxiliary methods

func (tblScan *TableScan) moveToBlock(blockNum int64) {
	tblScan.Close()
	blockId := file.NewBlockId(tblScan.fileName, blockNum)
	tblScan.recordPage = NewRecordPage(tblScan.tx, blockId, tblScan.layout)
	tblScan.currentSlot = -1
}

func (tblScan *TableScan) moveToNewBlock() error {
	tblScan.Close()
	blockId, err := tblScan.tx.Append(tblScan.fileName)
	if err != nil {
		return err
	}

	tblScan.recordPage = NewRecordPage(tblScan.tx, blockId, tblScan.layout)
	tblScan.recordPage.Format()
	tblScan.currentSlot = -1
	return nil
}

func (tblScan *TableScan) atLastBlock() (bool, error) {
	fileSize, err := tblScan.tx.Size(tblScan.fileName)
	if err != nil {
		return false, err
	}
	return tblScan.recordPage.BlockId().BlockNum == fileSize-1, nil
}
