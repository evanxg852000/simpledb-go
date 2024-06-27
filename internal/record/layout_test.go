package record_test

import (
	"testing"

	"github.com/evanxg852000/simpledb/internal/record"
	"github.com/stretchr/testify/assert"
)

func TestLayout(t *testing.T) {
	assert := assert.New(t)

	schema := record.NewSchema()
	schema.AddIntField("A")
	schema.AddStringField("B", 9)
	schema.AddIntField("C")

	layout := record.NewLayout(schema)
	expectedOffsets := []int64{8, 16, 33}
	for idx, fldName := range layout.Schema.Fields() {
		offset := layout.Offset(fldName)
		assert.Equal(expectedOffsets[idx], offset)
	}
}
