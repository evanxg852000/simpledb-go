package metadata_test

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"testing"

	"github.com/evanxg852000/simpledb/internal/metadata"
	"github.com/evanxg852000/simpledb/internal/record"
	"github.com/evanxg852000/simpledb/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestTableManager(t *testing.T) {
	assert := assert.New(t)
	workspaceDir, err := os.MkdirTemp("", "test_table_manager")
	assert.Nil(err)
	dbDir := path.Join(workspaceDir, "db")
	defer os.RemoveAll(workspaceDir)

	db := server.NewSimpleDB(dbDir, 400, 8)
	mdtManager := db.MetadataManager()
	tx := db.NewTx()

	schema := record.NewSchema()
	schema.AddIntField("A")
	schema.AddStringField("B", 9)
	mdtManager.CreateTable("my_table", schema, tx)

	// Statistics metadata
	layout, err := mdtManager.GetLayout("my_table", tx)
	assert.Nil(err)
	tblScan, err := record.NewTableScan(tx, "my_table", layout)
	assert.Nil(err)
	for i := 0; i < 50; i++ {
		tblScan.Insert()
		n := int64(rand.Intn(50))
		tblScan.SetInt("A", n)
		tblScan.SetString("B", fmt.Sprintf("rec_%d", n))
	}
	si := mdtManager.GetStatInfo("my_table", layout, tx)
	assert.Equal(metadata.NewStatInfo(5, 50), si)

	// View metadata
	viewDef := "select B from MyTable where A = 1"
	mdtManager.CreateView("my_view", viewDef, tx)
	v, _ := mdtManager.GetViewDef("my_view", tx)
	assert.Equal(viewDef, v)

	// Index metadata
	mdtManager.CreateIndex("idx_a", "my_table", "A", tx)
	mdtManager.CreateIndex("idx_b", "my_table", "B", tx)
	idxMap, _ := mdtManager.GetIndexInfo("my_table", tx)

	idxA := idxMap["A"]
	assert.Equal("idx_a", idxA.IndexName)
	assert.Equal("A", idxA.FieldName)

	idxB := idxMap["B"]
	assert.Equal("idx_b", idxB.IndexName)
	assert.Equal("B", idxB.FieldName)
	tx.Commit()
}
