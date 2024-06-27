package concurrency

import (
	"errors"
	"sync"
	"time"

	"github.com/evanxg852000/simpledb/internal/file"
	"github.com/evanxg852000/simpledb/internal/utils"
)

const (
	MAX_WAIT_TIME = 10 * time.Second
)

// The lock table, which provides methods to lock and unlock blocks.
// If a transaction requests a lock that causes a conflict with an
// existing lock, then that transaction is placed on a wait list.
// There is only one wait list for all blocks.
// When the last lock on a block is unlocked, then all transactions
// are removed from the wait list and rescheduled.
// If one of those transactions discovers that the lock it is waiting for
// is still locked, it will place itself back on the wait list.
type LockTable struct {
	locks map[file.BlockId]int
	mu    *sync.Mutex
	cond  *sync.Cond
}

func NewLockTable() *LockTable {
	mu := new(sync.Mutex)
	cond := sync.NewCond(mu)

	return &LockTable{
		locks: map[file.BlockId]int{},
		mu:    mu,
		cond:  cond,
	}
}

// Grant an SLock on the specified block.
// If an XLock exists when the method is called,
// then the calling thread will be placed on a wait list
// until the lock is released.
// If the thread remains on the wait list for a certain
// amount of time (currently 10 seconds),
// then an exception is thrown.
func (lt *LockTable) SLock(blockId file.BlockId) error {
	lt.mu.Lock()
	defer lt.mu.Unlock()

	startTimestamp := time.Now().UnixMilli()
	for lt.hasXLock(blockId) && !lt.waitToLong(startTimestamp) {
		utils.WaitCondWithTimeout(lt.cond, MAX_WAIT_TIME)
	}

	if lt.hasXLock(blockId) {
		return errors.New("LockAbortException")
	}

	value := lt.getLockValue(blockId) // will not be negative
	lt.locks[blockId] = value + 1
	return nil
}

// Grant an XLock on the specified block.
// If a lock of any type exists when the method is called,
// then the calling thread will be placed on a wait list
// until the locks are released.
// If the thread remains on the wait list for a certain
// amount of time (currently 10 seconds),
// then an exception is thrown.
func (lt *LockTable) XLock(blockId file.BlockId) error {
	lt.mu.Lock()
	defer lt.mu.Unlock()

	startTimestamp := time.Now().UnixMilli()
	for lt.hasOtherSLocks(blockId) && !lt.waitToLong(startTimestamp) {
		utils.WaitCondWithTimeout(lt.cond, MAX_WAIT_TIME)
	}
	if lt.hasOtherSLocks(blockId) {
		return errors.New("LockAbortException")
	}

	lt.locks[blockId] = -1
	return nil
}

func (lt *LockTable) Unlock(blockId file.BlockId) {
	value := lt.getLockValue(blockId)
	if value > 1 {
		lt.locks[blockId] = value - 1
	} else {
		delete(lt.locks, blockId)
		lt.cond.Broadcast()
	}
}

func (lt *LockTable) hasXLock(blockId file.BlockId) bool {
	return lt.getLockValue(blockId) < 0
}

func (lt *LockTable) hasOtherSLocks(blockId file.BlockId) bool {
	return lt.getLockValue(blockId) > 1
}

func (lt *LockTable) waitToLong(startTimestamp int64) bool {
	return time.Now().UnixMilli()-startTimestamp > MAX_WAIT_TIME.Milliseconds()
}

func (lt *LockTable) getLockValue(blockId file.BlockId) int {
	value, exists := lt.locks[blockId]
	if !exists {
		return 0
	}
	return value
}
