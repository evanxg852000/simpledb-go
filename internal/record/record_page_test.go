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

func TestRecordPage(t *testing.T) {
	assert := assert.New(t)
	workspaceDir, err := os.MkdirTemp("", "test_record_page")
	assert.Nil(err)
	dbDir := path.Join(workspaceDir, "db")
	defer os.RemoveAll(workspaceDir)

	db := server.NewSimpleDB(dbDir, 400, 8)
	tx := db.NewTx()

	schema := record.NewSchema()
	schema.AddIntField("A")
	schema.AddStringField("B", 9)

	layout := record.NewLayout(schema)

	blockId, err := tx.Append("test_record_page")
	assert.Nil(err)
	tx.Pin(blockId)

	recordPage := record.NewRecordPage(tx, blockId, layout)
	err = recordPage.Format()
	assert.Nil(err)

	// Filling the page with random records.
	type item struct {
		a int64
		b string
	}
	records := make([]int64, 0)
	slot := recordPage.InsertAfter(-1)
	for slot >= 0 {
		n := int64(rand.Intn(50))
		recordPage.SetInt(slot, "A", n)
		recordPage.SetString(slot, "B", fmt.Sprintf("rec_%d", n))
		slot = recordPage.InsertAfter(slot)
		records = append(records, n)
	}

	// Deleting records, whose A-values are less than 25.
	deletedRecords := make([]item, 0)
	remainingRecords := make([]item, 0)
	slot = recordPage.NextAfter(-1)
	for slot >= 0 {
		aVal, err := recordPage.GetInt(slot, "A")
		assert.Nil(err)
		bVal, err := recordPage.GetString(slot, "B")
		assert.Nil(err)
		rec := item{a: aVal, b: bVal}
		if rec.a < 25 {
			recordPage.Delete(slot)
			deletedRecords = append(deletedRecords, rec)
		} else {
			remainingRecords = append(remainingRecords, rec)
		}
		slot = recordPage.NextAfter(slot)
	}
	tx.Unpin(blockId)
	tx.Commit()

	// Checking remaining records.
	assert.Equal(len(records), len(deletedRecords)+len(remainingRecords))
	for _, rec := range deletedRecords {
		assert.Equal(true, rec.a < 25)
	}
	for _, rec := range remainingRecords {
		assert.Equal(true, rec.a >= 25)
	}
}
