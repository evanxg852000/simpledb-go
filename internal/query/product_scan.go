package query

type ProductScan struct {
	left  Scan
	right Scan
}

// Create a product scan having the two underlying scans.
func NewProductScan(left Scan, right Scan) *ProductScan {
	scan := &ProductScan{left, right}
	scan.BeforeFirst()
	return scan
}

// Position the scan before its first record.
// In particular, the LHS scan is positioned at
// its first record, and the RHS scan
// is positioned before its first record.
func (ps *ProductScan) BeforeFirst() {
	ps.left.BeforeFirst()
	ps.left.Next()
	ps.right.BeforeFirst()
}

// Move the scan to the next record.
// The method moves to the next RHS record, if possible.
// Otherwise, it moves to the next LHS record and the
// first RHS record.
// If there are no more LHS records, the method returns false.
func (ps *ProductScan) Next() bool {
	if ps.left.Next() {
		return true
	} else {
		ps.right.BeforeFirst()
		return ps.right.Next() && ps.left.Next()
	}
}

// Return the integer value of the specified field.
// The value is obtained from whichever scan
// contains the field.
func (ps *ProductScan) GetInt(fieldName string) int64 {
	if ps.left.HasField(fieldName) {
		return ps.left.GetInt(fieldName)
	}
	return ps.right.GetInt(fieldName)
}

// Returns the string value of the specified field.
// The value is obtained from whichever scan
// contains the field.
func (ps *ProductScan) GetString(fieldName string) string {
	if ps.left.HasField(fieldName) {
		return ps.left.GetString(fieldName)
	}
	return ps.right.GetString(fieldName)
}

// Return the value of the specified field.
// The value is obtained from whichever scan
// contains the field.
func (ps *ProductScan) GetValue(fieldName string) Constant {
	if ps.left.HasField(fieldName) {
		return ps.left.GetValue(fieldName)
	}
	return ps.right.GetValue(fieldName)
}

// Returns true if the specified field is in
// either of the underlying scans.
func (ps *ProductScan) HasField(fieldName string) bool {
	return ps.left.HasField(fieldName) || ps.right.HasField(fieldName)
}

func (ps *ProductScan) Close() {
	ps.left.Close()
	ps.right.Close()
}
