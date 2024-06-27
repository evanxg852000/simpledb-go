package hash

import (
	"github.com/evanxg852000/simpledb/internal/query"
	"github.com/evanxg852000/simpledb/internal/record"
	"github.com/evanxg852000/simpledb/internal/tx/recovery"
)

const (
	NUM_BUCKETS = 100
)

// A static hash implementation of the Index interface.
// A fixed number of buckets is allocated (currently, 100),
// and each bucket is implemented as a file of index records.
type HashIndex struct {
}

func NewHashIndex(tx *recovery.Transaction, indexName string, layout *record.Layout) *HashIndex {
	//TODO:
	return nil
}

func (hi *HashIndex) BeforeFirst(searchKey query.Constant) {

}

func (hi *HashIndex) Next() bool {
	//TODO:
	return false
}

func (hi *HashIndex) GetDataRID() query.RID {
	//TODO:
	return query.NewRID(0, 1)
}

func (hi *HashIndex) Insert(dataVal query.Constant, id query.RID) {
	//TODO:
}

func (hi *HashIndex) Delete(dataVal query.Constant, id query.RID) {
	//TODO:
}

func (hi *HashIndex) Close() {
	//TODO:
}

func SearchCost(numBlocks, recordPerBlock int64) int64 {
	return numBlocks / NUM_BUCKETS
}
