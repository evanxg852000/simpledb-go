package metadata

import (
	"fmt"

	"github.com/evanxg852000/simpledb/internal/record"
	"github.com/evanxg852000/simpledb/internal/tx/recovery"
)

// The index manager.
// The index manager has similar functionality to the table manager.
type IndexManager struct {
	layout       *record.Layout
	tableManager *TableManager
	statsManager *StatsManager
}

// Create the index manager.
// This constructor is called during system startup.
// If the database is new, then the index catalog table is created.
func NewIndexManager(isNew bool, tableManager *TableManager, statsManager *StatsManager, tx *recovery.Transaction) *IndexManager {
	if isNew {
		schema := record.NewSchema()
		schema.AddStringField("index_name", MAX_NAME_LENGTH)
		schema.AddStringField("table_name", MAX_NAME_LENGTH)
		schema.AddStringField("field_name", MAX_NAME_LENGTH)
		tableManager.CreateTable("index_catalog", schema, tx)
	}

	layout, err := tableManager.GetLayout("index_catalog", tx)
	if err != nil {
		fmt.Println("err: ", err)
	}

	return &IndexManager{layout, tableManager, statsManager}
}

// Create an index of the specified type for the specified field.
// A unique ID is assigned to this index, and its information
// is stored in the idxcat table.
func (indexManager *IndexManager) CreateIndex(idxName, tblName, fldName string, tx *recovery.Transaction) error {
	tableScan, err := record.NewTableScan(tx, "index_catalog", indexManager.layout)
	if err != nil {
		return err
	}
	tableScan.Insert()
	tableScan.SetString("index_name", idxName)
	tableScan.SetString("table_name", tblName)
	tableScan.SetString("field_name", fldName)
	tableScan.Close()
	return nil
}

// Return a map containing the index info for all indexes
// on the specified table.
func (indexManager *IndexManager) GetIndexInfo(tblName string, tx *recovery.Transaction) (map[string]IndexInfo, error) {
	result := map[string]IndexInfo{}
	tableScan, err := record.NewTableScan(tx, "index_catalog", indexManager.layout)
	if err != nil {
		return result, err
	}

	for tableScan.Next() {
		storedTableName := tableScan.GetString("table_name")
		if storedTableName == tblName {
			idxName := tableScan.GetString("index_name")
			fldName := tableScan.GetString("field_name")
			tableLayout, _ := indexManager.tableManager.GetLayout(tblName, tx)
			statsInfo := indexManager.statsManager.GetStatInfo(tblName, tableLayout, tx)
			idxInfo := NewIndexInfo(idxName, fldName, tableLayout.Schema, tx, statsInfo)
			result[fldName] = *idxInfo
		}
	}
	tableScan.Close()
	return result, nil
}
