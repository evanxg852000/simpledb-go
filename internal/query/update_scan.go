package query

// The interface implemented by all updateable scans.
type UpdateScan interface {
	Scan // update scan inherit a scan

	// Modify the field value of the current record.
	SetValue(fldName string, value Constant)

	// Modify the field value of the current record.
	SetInt(fldName string, value int64)

	// Modify the field value of the current record.
	SetString(fldName string, value string)

	// Insert a new record somewhere in the scan.
	Insert()

	// Delete the current record from the scan.
	Delete()

	// Return the id of the current record.
	GetRID() RID

	// Position the scan so that the current record has
	// the specified id.
	MoveToRID(rId RID)
}
