package tx_test

import (
	"os"
	"path"
	"testing"

	"github.com/evanxg852000/simpledb/internal/file"
	"github.com/evanxg852000/simpledb/internal/server"
	"github.com/evanxg852000/simpledb/internal/tx/recovery"
	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {
	assert := assert.New(t)

	workspaceDir, err := os.MkdirTemp("", "test_transaction")
	assert.Nil(err)
	dbDir := path.Join(workspaceDir, "db")
	defer os.RemoveAll(workspaceDir)

	db := server.NewSimpleDB(dbDir, 400, 8)
	fm := db.FileManager()
	bm := db.BufferManager()
	lm := db.LogManager()

	tx1 := recovery.NewTransaction(fm, lm, bm)
	blockId := file.NewBlockId("testfile", 1)
	tx1.Pin(blockId)
	// The block initially contains unknown bytes,
	// so don't log those values here.
	tx1.SetInt(blockId, 80, 1, false)
	tx1.SetString(blockId, 40, "one", false)
	tx1.Commit()

	tx2 := recovery.NewTransaction(fm, lm, bm)
	tx2.Pin(blockId)
	iVal, err := tx2.GetInt(blockId, 80)
	assert.Nil(err)
	assert.Equal(int64(1), iVal)

	sVal, err := tx2.GetString(blockId, 40)
	assert.Nil(err)
	assert.Equal("one", sVal)

	iVal = iVal + 1
	sVal = sVal + "!"
	tx2.SetInt(blockId, 80, iVal, true)
	tx2.SetString(blockId, 40, sVal, true)
	tx2.Commit()

	tx3 := recovery.NewTransaction(fm, lm, bm)
	tx3.Pin(blockId)
	iVal, err = tx3.GetInt(blockId, 80)
	assert.Nil(err)
	assert.Equal(int64(2), iVal)

	sVal, err = tx3.GetString(blockId, 40)
	assert.Nil(err)
	assert.Equal("one!", sVal)

	tx3.SetInt(blockId, 80, 9999, true)
	iVal, _ = tx3.GetInt(blockId, 80)
	assert.Equal(int64(9999), iVal)
	tx3.Rollback()

	tx4 := recovery.NewTransaction(fm, lm, bm)
	tx4.Pin(blockId)
	iVal, _ = tx4.GetInt(blockId, 80)
	assert.Equal(int64(2), iVal)
	tx4.Commit()
}
