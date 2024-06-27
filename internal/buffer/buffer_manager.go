package buffer

import (
	"fmt"
	"sync"
	"time"

	"github.com/evanxg852000/simpledb/internal/file"
	walog "github.com/evanxg852000/simpledb/internal/log"
	"github.com/evanxg852000/simpledb/internal/utils"
)

const (
	MAX_WAIT_TIME = 5 * time.Second
)

// Manages the pinning and unpinning of buffers to blocks.
type BufferManager struct {
	bufferPool   []Buffer
	numAvailable int64
	mu           *sync.Mutex
	cond         *sync.Cond
}

// Creates a buffer manager having the specified number
// of buffer slots.
func NewBufferManager(fileManager *file.FileManager, logManager *walog.LogManager, numSlot int) *BufferManager {
	bufferPool := make([]Buffer, numSlot)
	for i := 0; i < numSlot; i++ {
		bufferPool[i] = NewBuffer(fileManager, logManager)
	}

	mu := new(sync.Mutex)
	cond := sync.NewCond(mu)
	return &BufferManager{
		bufferPool:   bufferPool,
		numAvailable: int64(numSlot),
		mu:           mu,
		cond:         cond,
	}
}

// Returns the number of available (i.e. unpinned) buffers.
func (bm *BufferManager) Available() int64 {
	bm.mu.Lock()
	defer bm.mu.Unlock()
	return bm.numAvailable
}

// Flushes the dirty buffers modified by the specified transaction
func (bm *BufferManager) FlushAll(txNum int64) error {
	for _, buffer := range bm.bufferPool {
		if buffer.ModifyingTx() == txNum {
			err := buffer.flush()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Unpins the specified data buffer. If its pin count
// goes to zero, then notify any waiting threads.
func (bm *BufferManager) Unpin(buff *Buffer) {
	bm.mu.Lock()
	defer bm.mu.Unlock()
	buff.Unpin()
	// is this slot (buffer) free to be used by another blockId?
	if !buff.IsPinned() {
		bm.numAvailable += 1
		bm.cond.Broadcast()
	}
}

// Pins a buffer to the specified block, potentially
// waiting until a slot (buffer) becomes available.
// If no slot becomes available within a fixed
// time period (10 seconds), then an error is returned.
func (bm *BufferManager) Pin(blockId file.BlockId) (*Buffer, error) {
	bm.mu.Lock()
	defer bm.mu.Unlock()

	startTimestamp := time.Now().UnixMilli()
	buff := bm.tryToPin(blockId)
	for buff == nil && !bm.waitToLong(startTimestamp) {
		utils.WaitCondWithTimeout(bm.cond, MAX_WAIT_TIME)
		buff = bm.tryToPin(blockId)
	}

	if buff == nil {
		return nil, fmt.Errorf("buffer request waited for too long")
	}
	return buff, nil
}

func (bm *BufferManager) waitToLong(startTimestamp int64) bool {
	return time.Now().UnixMilli()-startTimestamp > MAX_WAIT_TIME.Milliseconds()
}

// Tries to pin a buffer to the specified block.
// If there is already a buffer assigned to that block
// then that buffer is used;
// otherwise, an unpinned buffer from the pool is chosen.
// Returns a null value if there are no available buffers.
func (bm *BufferManager) tryToPin(blockId file.BlockId) *Buffer {
	buff := bm.findExistingBuffer(blockId)
	if buff == nil {
		buff = bm.chooseUnpinnedBuffer()
		if buff == nil {
			return nil
		}
		buff.AssignToBlock(blockId)
	}

	if !buff.IsPinned() {
		bm.numAvailable -= 1
	}
	buff.Pin()
	return buff
}

func (bm *BufferManager) findExistingBuffer(blockId file.BlockId) *Buffer {
	for i := 0; i < len(bm.bufferPool); i++ {
		buff := &bm.bufferPool[i]
		if blockId.Equals(buff.blockId) {
			return buff
		}
	}
	return nil
}

func (bm *BufferManager) chooseUnpinnedBuffer() *Buffer {
	for i := 0; i < len(bm.bufferPool); i++ {
		buff := &bm.bufferPool[i]
		if !buff.IsPinned() {
			return buff
		}
	}
	return nil
}
