package tx_test

import (
	"os"
	"path"
	"sync"
	"testing"

	"github.com/evanxg852000/simpledb/internal/server"
	"github.com/evanxg852000/simpledb/internal/tx/recovery"
	"github.com/stretchr/testify/assert"
)

func TestConcurrency(t *testing.T) {
	assert := assert.New(t)
	workspaceDir, err := os.MkdirTemp("", "test_concurrency")
	assert.Nil(err)
	dbDir := path.Join(workspaceDir, "db")
	defer os.RemoveAll(workspaceDir)

	db := server.NewSimpleDB(dbDir, 400, 8)
	fm := db.FileManager()
	bm := db.BufferManager()
	lm := db.LogManager()

	txa := recovery.NewTransaction(fm, lm, bm)

	wg := sync.WaitGroup{}
	wg.Add(3)

	go A(assert, txa, &wg)
	go B(assert, txa, &wg)
	go C(assert, txa, &wg)

	wg.Wait()
}

func A(assert *assert.Assertions, tx *recovery.Transaction, wg *sync.WaitGroup) {
	defer wg.Done()
	//TODO:
}

func B(assert *assert.Assertions, tx *recovery.Transaction, wg *sync.WaitGroup) {
	defer wg.Done()
	//TODO:
}

func C(assert *assert.Assertions, tx *recovery.Transaction, wg *sync.WaitGroup) {
	defer wg.Done()
	//TODO:
}
