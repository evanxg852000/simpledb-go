package record

import "github.com/evanxg852000/simpledb/internal/file"

// Description of the structure of a record.
// It contains the name, type, length and offset of
// each field of the table.
type Layout struct {
	Schema   *Schema
	offsets  map[string]int64
	slotSize int64
}

// This constructor creates a Layout object from a schema.
// This constructor is used when a table
// is created. It determines the physical offset of
// each field within the record.
func NewLayout(schema *Schema) *Layout {
	pos := int64(8) // // leave space for the empty/in-use flag
	offsets := map[string]int64{}
	for _, fldName := range schema.Fields() {
		offsets[fldName] = pos
		pos += lengthInBytes(schema, fldName)
	}
	return &Layout{
		Schema:   schema,
		offsets:  offsets,
		slotSize: pos,
	}
}

// Create a Layout object from the specified metadata.
// This constructor is used when the metadata
// is retrieved from the catalog.
func NewLayoutFromMetadata(schema *Schema, offsets map[string]int64, slotSize int64) *Layout {
	return &Layout{
		Schema:   schema,
		offsets:  offsets,
		slotSize: slotSize,
	}
}

// Return the offset of a specified field within a record
func (layout *Layout) Offset(fldName string) int64 {
	return layout.offsets[fldName]
}

// Return the size of a slot, in bytes.
func (layout *Layout) SlotSize() int64 {
	return layout.slotSize
}

func lengthInBytes(schema *Schema, fldName string) int64 {
	fldType := schema.FieldType(fldName)
	if fldType == INTEGER_TYPE {
		return 8
	}
	// STRING_TYPE
	return file.GetEncodingLength(schema.FieldLength(fldName))
}
