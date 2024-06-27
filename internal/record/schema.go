package record

import (
	"slices"

	"github.com/evanxg852000/simpledb/internal/query"
)

const (
	_ = iota
	INTEGER_TYPE
	STRING_TYPE
)

type FieldInfo struct {
	dataType int64
	length   int64
}

func NewFieldInfo(fType int64, length int64) FieldInfo {
	return FieldInfo{fType, length}
}

// The record schema of a table.
// A schema contains the name and type of
// each field of the table, as well as the length
// of each varchar field.
type Schema struct {
	fields []string
	info   map[string]FieldInfo
}

func NewSchema() *Schema {
	return &Schema{
		fields: make([]string, 0),
		info:   make(map[string]FieldInfo),
	}
}

// Add a field to the schema having a specified
// name, type, and length.
// If the field type is "integer", then the length
// value is irrelevant.
func (schema *Schema) AddField(fldName string, fldType int64, fldLength int64) {
	schema.fields = append(schema.fields, fldName)
	schema.info[fldName] = NewFieldInfo(fldType, fldLength)
}

// Add an integer field to the schema.
func (schema *Schema) AddIntField(fldName string) {
	schema.AddField(fldName, INTEGER_TYPE, 0)
}

// Add a string field to the schema.
// The length is the conceptual length of the field.
// For example, if the field is defined as varchar(8),
// then its length is 8.
func (schema *Schema) AddStringField(fldName string, fldLength int64) {
	schema.AddField(fldName, STRING_TYPE, fldLength)
}

// Add a field to the schema having the same
// type and length as the corresponding field
// in another schema.
func (schema *Schema) Add(fldName string, sch Schema) {
	fldType := sch.FieldType(fldName)
	fldLength := sch.FieldLength(fldName)
	schema.AddField(fldName, fldType, fldLength)
}

// Add all of the fields in the specified schema
// to the current schema.
func (schema *Schema) AddAll(sch Schema) {
	for _, fldName := range sch.fields {
		schema.Add(fldName, sch)
	}
}

// Return a collection containing the name of
// each field in the schema.
func (schema *Schema) Fields() []string {
	return schema.fields
}

// Return true if the specified field
// is in the schema
func (schema *Schema) HasField(fldName string) bool {
	return slices.Contains(schema.fields, fldName)
}

// Return the type of the specified field.
func (schema *Schema) FieldType(fldName string) int64 {
	return schema.info[fldName].dataType
}

// Return the conceptual length of the specified field.
// If the field is not a string field, then
// the return value is undefined.
func (schema *Schema) FieldLength(fldName string) int64 {
	return schema.info[fldName].length
}

// Determine if all of the fields mentioned in this expression
// are contained in the specified schema.
func (schema *Schema) AppliesTo(expr *query.Expression) bool {
	if !expr.IsFieldName() {
		return true
	}
	return schema.HasField(expr.AsFieldName())
}
