package buffer_test

import (
	"os"
	"testing"

	"github.com/evanxg852000/simpledb/internal/buffer"
	"github.com/evanxg852000/simpledb/internal/file"
	walog "github.com/evanxg852000/simpledb/internal/log"
	"github.com/stretchr/testify/assert"
)

func TestBufferManager(t *testing.T) {
	assert := assert.New(t)

	dbDirectory, err := os.MkdirTemp("", "test_buffer_manager_")
	assert.Nil(err)
	defer os.RemoveAll(dbDirectory)

	fm, err := file.NewFileManager(dbDirectory, 512)
	assert.Nil(err)

	lm, err := walog.NewLogManager(fm, "log_file")
	assert.Nil(err)

	bm := buffer.NewBufferManager(fm, lm, 3)
	buffers := make([]*buffer.Buffer, 6)

	buffers[0], err = bm.Pin(file.NewBlockId("testfile", 0))
	assert.Nil(err)

	buffers[1], err = bm.Pin(file.NewBlockId("testfile", 1))
	assert.Nil(err)

	buffers[2], err = bm.Pin(file.NewBlockId("testfile", 2))
	assert.Nil(err)

	bm.Unpin(buffers[1])
	buffers[1] = nil

	buffers[3], err = bm.Pin(file.NewBlockId("testfile", 0)) // block 0 pinned twice
	assert.Nil(err)
	assert.Equal(int64(1), bm.Available())

	buffers[4], err = bm.Pin(file.NewBlockId("testfile", 1)) // block 1 repinned
	assert.Nil(err)

	// Attempting to pin block 3 will not work, no buffer left
	buffers[5], err = bm.Pin(file.NewBlockId("testfile", 3))
	assert.NotNil(err)

	bm.Unpin(buffers[2])
	buffers[2] = nil
	buffers[5], err = bm.Pin(file.NewBlockId("testfile", 3)) // now this works
	assert.Nil(err)

	//check final buffers state
	assert.Equal(int64(0), buffers[0].Block().BlockNum)
	assert.Empty(buffers[1])
	assert.Empty(buffers[2])
	assert.Equal(int64(0), buffers[3].Block().BlockNum)
	assert.Equal(int64(1), buffers[4].Block().BlockNum)
	assert.Equal(int64(3), buffers[5].Block().BlockNum)
}
