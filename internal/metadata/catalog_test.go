package metadata_test

import (
	"os"
	"path"
	"testing"

	"github.com/evanxg852000/simpledb/internal/record"
	"github.com/evanxg852000/simpledb/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestCatalog(t *testing.T) {
	assert := assert.New(t)
	workspaceDir, err := os.MkdirTemp("", "test_catalog")
	assert.Nil(err)
	dbDir := path.Join(workspaceDir, "db")
	defer os.RemoveAll(workspaceDir)

	db := server.NewSimpleDB(dbDir, 400, 8)
	tx := db.NewTx()

	tblManager := db.MetadataManager().GetTableManager()
	// Get all tables and their slot size
	layout, err := tblManager.GetLayout("table_catalog", tx)
	assert.Nil(err)
	tblScan, err := record.NewTableScan(tx, "table_catalog", layout)
	assert.Nil(err)
	rows := []struct {
		tblName  string
		slotSize int64
	}{}
	for tblScan.Next() {
		tblName := tblScan.GetString("table_name")
		slotSize := tblScan.GetInt("slot_size")
		rows = append(rows, struct {
			tblName  string
			slotSize int64
		}{tblName, slotSize})
	}
	assert.Equal([]struct {
		tblName  string
		slotSize int64
	}{
		{"table_catalog", 56},
		{"field_catalog", 112},
		{"view_catalog", 156},
		{"index_catalog", 128},
	}, rows)

	tblScan.Close()

	// Get all fields for each table and their offsets
	layout, err = tblManager.GetLayout("field_catalog", tx)
	assert.Nil(err)
	tblScan, err = record.NewTableScan(tx, "field_catalog", layout)
	assert.Nil(err)
	rows2 := []struct {
		tblName string
		fldName string
		offset  int64
	}{}
	for tblScan.Next() {
		tblName := tblScan.GetString("table_name")
		fldName := tblScan.GetString("field_name")
		offset := tblScan.GetInt("offset")
		rows2 = append(rows2, struct {
			tblName string
			fldName string
			offset  int64
		}{tblName, fldName, offset})
	}
	assert.Equal(12, len(rows2))
	tblScan.Close()
}
