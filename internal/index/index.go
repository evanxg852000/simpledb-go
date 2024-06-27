package index

import "github.com/evanxg852000/simpledb/internal/query"

type Index interface {
	// Positions the index before the first record
	// having the specified search key.
	BeforeFirst(searchKey query.Constant)

	// Moves the index to the next record having the
	// search key specified in the beforeFirst method.
	// Returns false if there are no more such index records.
	Next() bool

	// Returns the dataRID value stored in the current index record.
	GetDataRID() query.RID

	// Inserts an index record having the specified
	// dataval and dataRID values.
	Insert(dataVal query.Constant, id query.RID)

	// Deletes the index record having the specified
	// dataval and dataRID values.
	Delete(dataVal query.Constant, id query.RID)

	// Closes the index.
	Close()
}
