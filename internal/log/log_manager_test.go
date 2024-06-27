package log_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/evanxg852000/simpledb/internal/file"
	walog "github.com/evanxg852000/simpledb/internal/log"
	"github.com/stretchr/testify/assert"
)

func TestLogManager(t *testing.T) {
	assert := assert.New(t)

	dbDirectory, err := os.MkdirTemp("", "test_log_manager_")
	assert.Nil(err)
	defer os.RemoveAll(dbDirectory)

	fm, err := file.NewFileManager(dbDirectory, 512)
	assert.Nil(err)

	logManager, err := walog.NewLogManager(fm, "log_file")
	assert.Nil(err)

	// create records
	expectedItems := []string{}
	suffix := strings.Repeat("X", 10)
	for i := 1; i <= 1500; i++ {
		item := fmt.Sprintf("LOG_%d_%s", i, suffix)
		expectedItems = append(expectedItems, item)
		lsn, err := logManager.Append([]byte(item))
		assert.Nil(err)
		if i%150 == 0 {
			err = logManager.Flush(lsn)
			assert.Nil(err)
		}
	}

	// iterate
	logIterator, err := logManager.Iterator()
	assert.Nil(err)
	index := 1500
	for logIterator.HasNext() {
		index = index - 1
		foundItem, err := logIterator.Next()
		assert.Nil(err)
		assert.Equal(expectedItems[index], string(foundItem))
	}

}
