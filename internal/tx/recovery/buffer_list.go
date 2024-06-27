package recovery

import (
	"slices"

	"github.com/evanxg852000/simpledb/internal/buffer"
	"github.com/evanxg852000/simpledb/internal/file"
)

type BufferList struct {
	buffers       map[file.BlockId]*buffer.Buffer
	pins          []file.BlockId
	bufferManager *buffer.BufferManager
}

func NewBufferList(bufferManager *buffer.BufferManager) *BufferList {
	return &BufferList{
		buffers:       make(map[file.BlockId]*buffer.Buffer),
		pins:          make([]file.BlockId, 0),
		bufferManager: bufferManager,
	}
}

func (bl *BufferList) GetBuffer(blockId file.BlockId) *buffer.Buffer {
	return bl.buffers[blockId]
}

func (bl *BufferList) Pin(blockId file.BlockId) error {
	buffer, err := bl.bufferManager.Pin(blockId)
	if err != nil {
		return err
	}
	bl.buffers[blockId] = buffer
	bl.pins = append(bl.pins, blockId)
	return nil
}

func (bl *BufferList) Unpin(blockId file.BlockId) {
	buffer, exist := bl.buffers[blockId]
	if !exist {
		return
	}
	bl.bufferManager.Unpin(buffer)
	deleted := false
	bl.pins = slices.DeleteFunc(bl.pins, func(probe file.BlockId) bool {
		if !deleted && probe == blockId {
			deleted = true
			return true
		}
		return false
	})

	if !slices.Contains(bl.pins, blockId) {
		delete(bl.buffers, blockId)
	}
}

func (bl *BufferList) UnpinAll() {
	for _, blockId := range bl.pins {
		buffer := bl.buffers[blockId]
		bl.bufferManager.Unpin(buffer)
	}
	clear(bl.buffers)
	clear(bl.pins)
}
