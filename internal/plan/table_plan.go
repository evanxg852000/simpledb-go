package plan

import (
	"fmt"

	"github.com/evanxg852000/simpledb/internal/metadata"
	"github.com/evanxg852000/simpledb/internal/query"
	"github.com/evanxg852000/simpledb/internal/record"
	"github.com/evanxg852000/simpledb/internal/tx/recovery"
)

// The Plan class corresponding to a table.
type TablePlan struct {
	tableName string
	tx        *recovery.Transaction
	layout    *record.Layout
	statInfo  metadata.StatInfo
}

// Creates a leaf node in the query tree corresponding
// to the specified table.
func NewTablePlan(tx *recovery.Transaction, tableName string, mdtManager *metadata.MetadataManager) *TablePlan {
	layout, err := mdtManager.GetLayout(tableName, tx)
	if err != nil {
		panic(fmt.Sprintf("could not retrieve table `%s` layout", tableName))
	}
	statInfo := mdtManager.GetStatInfo(tableName, layout, tx)
	return &TablePlan{
		tableName,
		tx,
		layout,
		statInfo,
	}
}

// Creates a table scan for this query.
func (tblPlan *TablePlan) Open() query.Scan {
	scan, err := record.NewTableScan(tblPlan.tx, tblPlan.tableName, tblPlan.layout)
	if err != nil {
		panic(fmt.Sprint("could not create table plan", err))
	}
	return any(scan).(query.Scan)
}

// Estimates the number of block accesses for the table,
// which is obtainable from the statistics manager.
func (tblPlan *TablePlan) BlockAccessed() int64 {
	return tblPlan.statInfo.BlockAccessed()
}

// Estimates the number of records in the table,
// which is obtainable from the statistics manager.
func (tblPlan *TablePlan) RecordsOutput() int64 {
	return tblPlan.statInfo.RecordsOutput()
}

// Estimates the number of distinct field values in the table,
// which is obtainable from the statistics manager.
func (tblPlan *TablePlan) DistinctValues(fieldName string) int64 {
	return tblPlan.statInfo.DistinctValues(fieldName)
}

// Determines the schema of the table,
// which is obtainable from the catalog manager.
func (tblPlan *TablePlan) Schema() record.Schema {
	return *tblPlan.layout.Schema
}
