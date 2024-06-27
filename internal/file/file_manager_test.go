package file_test

import (
	"os"
	"testing"

	"github.com/evanxg852000/simpledb/internal/file"
	"github.com/stretchr/testify/assert"
)

func TestFileManager(t *testing.T) {
	assert := assert.New(t)

	dbDirectory, err := os.MkdirTemp("", "test_file_manager_")
	assert.Nil(err)
	defer os.RemoveAll(dbDirectory)

	fm, err := file.NewFileManager(dbDirectory, 800)
	assert.Nil(err)

	blockId := file.NewBlockId("testfile", 2)
	pos1 := int64(88)

	p1 := file.NewPage(fm.BlockSize())
	n, err := p1.WriteString(pos1, "abcdefghijklm")
	assert.Nil(err)
	pos2 := pos1 + n
	p1.WriteInt(pos2, -345)
	pos3 := pos2 + 8
	p1.WriteFloat(pos3, -3.14)
	fm.Write(blockId, &p1)

	p2 := file.NewPage(fm.BlockSize())
	fm.Read(blockId, &p2)

	s, err := p2.ReadString(pos1)
	assert.Nil(err)
	assert.Equal("abcdefghijklm", s)

	v, err := p2.ReadInt(pos2)
	assert.Nil(err)
	assert.Equal(int64(-345), v)

	f, err := p2.ReadFloat(pos3)
	assert.Nil(err)
	assert.Equal(float64(-3.14), f)
}
