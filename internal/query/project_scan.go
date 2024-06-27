package query

import (
	"fmt"
	"slices"
)

// The scan class corresponding to the <i>project</i> relational
// algebra operator.
// All methods except hasField delegate their work to the
// underlying scan.
type ProjectScan struct {
	scan   Scan
	fields []string
}

// Create a project scan having the specified
// underlying scan and field list.
func NewProjectScan(scan Scan, fields []string) *ProjectScan {
	return &ProjectScan{scan, fields}
}

func (ps *ProjectScan) BeforeFirst() {
	ps.scan.BeforeFirst()
}

func (ps *ProjectScan) Next() bool {
	return ps.scan.Next()
}

func (ps *ProjectScan) GetInt(fieldName string) int64 {
	if ps.HasField(fieldName) {
		return ps.scan.GetInt(fieldName)
	}
	panic(fmt.Sprintf("field `%v` not found.", fieldName))
}

func (ps *ProjectScan) GetString(fieldName string) string {
	if ps.HasField(fieldName) {
		return ps.scan.GetString(fieldName)
	}
	panic(fmt.Sprintf("field `%v` not found.", fieldName))
}

func (ps *ProjectScan) GetValue(fieldName string) Constant {
	if ps.HasField(fieldName) {
		return ps.scan.GetValue(fieldName)
	}
	panic(fmt.Sprintf("field `%v` not found.", fieldName))
}

func (ps *ProjectScan) HasField(fieldName string) bool {
	return slices.Contains(ps.fields, fieldName)
}

func (ps *ProjectScan) Close() {
	ps.scan.Close()
}
