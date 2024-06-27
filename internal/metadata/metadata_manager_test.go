package metadata_test

import (
	"os"
	"path"
	"testing"

	"github.com/evanxg852000/simpledb/internal/record"
	"github.com/evanxg852000/simpledb/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestMetadataManager(t *testing.T) {
	assert := assert.New(t)
	workspaceDir, err := os.MkdirTemp("", "test_table_manager")
	assert.Nil(err)
	dbDir := path.Join(workspaceDir, "db")
	defer os.RemoveAll(workspaceDir)

	db := server.NewSimpleDB(dbDir, 400, 8)
	tblManager := db.MetadataManager().GetTableManager()

	tx := db.NewTx()

	schema := record.NewSchema()
	schema.AddIntField("A")
	schema.AddStringField("B", 9)
	tblManager.CreateTable("my_table", schema, tx)

	layout, err := tblManager.GetLayout("my_table", tx)
	assert.Nil(err)

	assert.Equal(int64(33), layout.SlotSize())
	assert.Equal(2, len(layout.Schema.Fields()))
	rows := []struct {
		n string
		t int64
	}{}
	for _, fldName := range layout.Schema.Fields() {
		rows = append(rows, struct {
			n string
			t int64
		}{fldName, layout.Schema.FieldType(fldName)})
	}

	tx.Commit()

	assert.Equal([]struct {
		n string
		t int64
	}{{"A", record.INTEGER_TYPE}, {"B", record.STRING_TYPE}}, rows)
}
