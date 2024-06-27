package query

// The interface will be implemented by each query scan.
// There is a Scan class for each relational
// algebra operator.
type Scan interface {
	// Position the scan before its first record. A
	// subsequent call to next() will return the first record.
	BeforeFirst()

	// Move the scan to the next record.
	Next() bool

	// Return the value of the specified integer field
	// in the current record.
	GetInt(fieldName string) int64

	// Return the value of the specified string field
	// in the current record.
	GetString(fieldName string) string

	// Return the value of the specified field in the current record.
	// The value is expressed as a Constant.
	GetValue(fieldName string) Constant

	// Return true if the scan has the specified field.
	HasField(fieldName string) bool

	// Close the scan and its sub_scans, if any
	Close()
}
