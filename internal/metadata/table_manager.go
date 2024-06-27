package metadata

import (
	"github.com/evanxg852000/simpledb/internal/record"
	"github.com/evanxg852000/simpledb/internal/tx/recovery"
)

const (
	//The max characters a tablename or fieldname can have.
	MAX_NAME_LENGTH = 32

	TABLE_CATALOG = "table_catalog"
	FIELD_CATALOG = "field_catalog"
)

// The table manager.
// There are methods to create a table, save the metadata
// in the catalog, and obtain the metadata of a
// previously-created table.
type TableManager struct {
	tableCatLayout *record.Layout
	fieldCatLayout *record.Layout
}

// Create a new table manager for the database system.
// If the database is new, the two catalog tables (table, field)
// are created.
func NewTableManager(isNew bool, tx *recovery.Transaction) *TableManager {
	tableCatSchema := record.NewSchema()
	tableCatSchema.AddStringField("table_name", MAX_NAME_LENGTH)
	tableCatSchema.AddIntField("slot_size")
	tableCatLayout := record.NewLayout(tableCatSchema)

	fieldCatSchema := record.NewSchema()
	fieldCatSchema.AddStringField("table_name", MAX_NAME_LENGTH)
	fieldCatSchema.AddStringField("field_name", MAX_NAME_LENGTH)
	fieldCatSchema.AddIntField("type")
	fieldCatSchema.AddIntField("length")
	fieldCatSchema.AddIntField("offset")
	fieldCatLayout := record.NewLayout(fieldCatSchema)

	tableManager := &TableManager{tableCatLayout, fieldCatLayout}
	if isNew {
		tableManager.CreateTable(TABLE_CATALOG, tableCatSchema, tx)
		tableManager.CreateTable(FIELD_CATALOG, fieldCatSchema, tx)
	}
	return tableManager
}

// Create a new table having the specified name and schema
func (tableManager *TableManager) CreateTable(tblName string, schema *record.Schema, tx *recovery.Transaction) error {
	layout := record.NewLayout(schema)

	// insert one record into table_catalog
	tableScan, err := record.NewTableScan(tx, TABLE_CATALOG, tableManager.tableCatLayout)
	if err != nil {
		return err
	}
	tableScan.Insert()
	tableScan.SetString("table_name", tblName)
	tableScan.SetInt("slot_size", layout.SlotSize())
	tableScan.Close()

	// insert one record into field_catalog
	tableScan, err = record.NewTableScan(tx, FIELD_CATALOG, tableManager.fieldCatLayout)
	if err != nil {
		return err
	}
	for _, fldName := range schema.Fields() {
		tableScan.Insert()
		tableScan.SetString("table_name", tblName)
		tableScan.SetString("field_name", fldName)
		tableScan.SetInt("type", int64(schema.FieldType(fldName)))
		tableScan.SetInt("length", int64(schema.FieldLength(fldName)))
		tableScan.SetInt("offset", layout.Offset(fldName))
	}
	tableScan.Close()
	return nil
}

func (tableManager *TableManager) GetLayout(tblName string, tx *recovery.Transaction) (*record.Layout, error) {
	size := int64(-1)
	tableScan, err := record.NewTableScan(tx, TABLE_CATALOG, tableManager.tableCatLayout)
	if err != nil {
		return nil, err
	}
	for tableScan.Next() {
		storedTblName := tableScan.GetString("table_name")
		if storedTblName == tblName {
			size = tableScan.GetInt("slot_size")
			break
		}
	}
	tableScan.Close()

	schema := record.NewSchema()
	offsets := make(map[string]int64)
	tableScan, err = record.NewTableScan(tx, FIELD_CATALOG, tableManager.fieldCatLayout)
	if err != nil {
		return nil, err
	}

	for tableScan.Next() {
		storedTblName := tableScan.GetString("table_name")
		if storedTblName == tblName {
			fldName := tableScan.GetString("field_name")
			fldType := tableScan.GetInt("type")
			fldLength := tableScan.GetInt("length")
			offset := tableScan.GetInt("offset")
			offsets[fldName] = offset
			schema.AddField(fldName, fldType, fldLength)
		}
	}
	tableScan.Close()
	return record.NewLayoutFromMetadata(schema, offsets, size), nil
}
