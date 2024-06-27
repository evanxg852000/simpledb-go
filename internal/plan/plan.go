package plan

import (
	"github.com/evanxg852000/simpledb/internal/query"
	"github.com/evanxg852000/simpledb/internal/record"
)

// The interface implemented by each query plan.
// There is a Plan class for each relational algebra operator.
type Plan interface {
	// Opens a scan corresponding to this plan.
	// The scan will be positioned before its first record.
	Open() query.Scan

	// Returns an estimate of the number of block accesses
	// that will occur when the scan is read to completion.
	BlockAccessed() int64

	// Returns an estimate of the number of records
	// in the query's output table.
	RecordsOutput() int64

	// Returns an estimate of the number of distinct values
	// for the specified field in the query's output table.
	DistinctValues(string) int64

	// Returns the schema of the query.
	Schema() record.Schema
}
