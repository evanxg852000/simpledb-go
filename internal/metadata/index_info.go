package metadata

import (
	"github.com/evanxg852000/simpledb/internal/index"
	"github.com/evanxg852000/simpledb/internal/index/hash"
	"github.com/evanxg852000/simpledb/internal/record"
	"github.com/evanxg852000/simpledb/internal/tx/recovery"
)

type IndexInfo struct {
	IndexName string
	FieldName string
	tx        *recovery.Transaction
	schema    *record.Schema
	layout    *record.Layout
	statsInfo StatInfo
}

// Create an IndexInfo object for the specified index.
func NewIndexInfo(indexName string, fieldName string,
	schema *record.Schema, tx *recovery.Transaction, si StatInfo) *IndexInfo {
	fldType := schema.FieldType(fieldName)
	fldLength := schema.FieldLength(fieldName)

	layout := createIndexLayout(fldType, fldLength)
	return &IndexInfo{
		indexName,
		fieldName,
		tx,
		schema,
		layout,
		si,
	}
}

// Open the index described by this object.
func (ii *IndexInfo) Open() index.Index {
	// return new BTreeIndex(ii.tx, ii.indexName, ii.layout)
	return hash.NewHashIndex(ii.tx, ii.IndexName, ii.layout)
}

// Estimate the number of block accesses required to
// find all index records having a particular search key.
// The method uses the table's metadata to estimate the
// size of the index file and the number of index records
// per block.
// It then passes this information to the traversalCost
// method of the appropriate index type,
// which provides the estimate.
func (ii *IndexInfo) BlockAccessed() int64 {
	recordPerBlock := ii.tx.BlockSize() / ii.layout.SlotSize()
	numBlocks := ii.statsInfo.RecordsOutput() / recordPerBlock
	// return btree.SearchCost(numBlocks, recordPerBlock)
	return hash.SearchCost(numBlocks, recordPerBlock)
}

// Return the estimated number of records having a
// search key.  This value is the same as doing a select
// query; that is, it is the number of records in the table
// divided by the number of distinct values of the indexed field.
func (ii *IndexInfo) RecordsOutput() int64 {
	return ii.statsInfo.RecordsOutput() / ii.statsInfo.DistinctValues(ii.FieldName)
}

// Return the distinct values for a specified field
// in the underlying table, or 1 for the indexed field.
func (ii IndexInfo) DistinctValues(fieldName string) int64 {
	if ii.FieldName == fieldName {
		return 1
	}
	return ii.statsInfo.DistinctValues(ii.FieldName)
}

// Return the layout of the index records.
// The schema consists of the dataRID (which is
// represented as two integers, the block number and the
// record ID) and the dataval (which is the indexed field).
// Schema information about the indexed field is obtained
// via the table's schema.
func createIndexLayout(fldType, fldLength int64) *record.Layout {
	schema := record.NewSchema()
	schema.AddIntField("block")
	schema.AddIntField("id")
	if int(fldType) == record.INTEGER_TYPE {
		schema.AddIntField("data_val")
	} else {
		schema.AddStringField("data_val", fldLength)
	}
	return record.NewLayout(schema)
}
