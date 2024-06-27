package record_test

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"testing"

	"github.com/evanxg852000/simpledb/internal/record"
	"github.com/evanxg852000/simpledb/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestTableScan(t *testing.T) {
	assert := assert.New(t)
	workspaceDir, err := os.MkdirTemp("", "test_table_scan")
	assert.Nil(err)
	dbDir := path.Join(workspaceDir, "db")
	defer os.RemoveAll(workspaceDir)

	db := server.NewSimpleDB(dbDir, 400, 8)
	tx := db.NewTx()

	schema := record.NewSchema()
	schema.AddIntField("A")
	schema.AddStringField("B", 9)

	layout := record.NewLayout(schema)

	// Filling the table with 50 random records.
	tblScan, err := record.NewTableScan(tx, "T", layout)
	assert.Nil(err)
	for i := 0; i < 50; i++ {
		tblScan.Insert()

		n := int64(rand.Intn(50))
		tblScan.SetInt("A", n)
		tblScan.SetString("B", fmt.Sprintf("rec_%d", n))

		rId := tblScan.GetRID()
		assert.True(rId.BlockNum >= 0)
		assert.True(rId.Slot >= 0)
	}

	// Deleting records, whose A-values are less than 25.
	deletedRecords := make([]int64, 0)
	remainingRecords := make([]int64, 0)
	tblScan.BeforeFirst()
	for tblScan.Next() {
		aVal := tblScan.GetInt("A")
		_ = tblScan.GetString("B")
		assert.Nil(err)
		if aVal < 25 {
			tblScan.Delete()
			deletedRecords = append(deletedRecords, aVal)
		} else {
			remainingRecords = append(remainingRecords, aVal)
		}
	}

	assert.Equal(50, len(deletedRecords)+len(remainingRecords))
	tblScan.Close()
	tx.Commit()
}
