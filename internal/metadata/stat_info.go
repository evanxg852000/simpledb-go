package metadata

// A StatInfo object holds three pieces of
// statistical information about a table:
// the number of blocks, the number of records,
// and the number of distinct values for each field.
type StatInfo struct {
	numBlocks  int64
	numRecords int64
}

// Create a StatInfo object.
// Note that the number of distinct values is not
// passed into the constructor.
// The object fakes this value.
func NewStatInfo(numBlocks, numRecords int64) StatInfo {
	return StatInfo{numBlocks, numRecords}
}

// Return the estimated number of blocks in the table.
func (si *StatInfo) BlockAccessed() int64 {
	return si.numBlocks
}

// Return the estimated number of records in the table.
func (si *StatInfo) RecordsOutput() int64 {
	return si.numRecords
}

// Return the estimated number of distinct values
// for the specified field.
// This estimate is a complete guess, because doing something
// reasonable is beyond the scope of this system.
func (si *StatInfo) DistinctValues(fieldName string) int64 {
	_ = fieldName
	return (si.numRecords / 3) + 1
}
