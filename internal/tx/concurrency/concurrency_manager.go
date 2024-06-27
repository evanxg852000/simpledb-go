package concurrency

import "github.com/evanxg852000/simpledb/internal/file"

// The global lock table. all transactions
// share the same lock table.
var lockTableInstance *LockTable = NewLockTable()

// The concurrency manager for the transaction.
// Each transaction has its own concurrency manager.
// The concurrency manager keeps track of which locks the
// transaction currently has, and interacts with the
// global lock table as needed.
type ConcurrencyManager struct {
	lockTable *LockTable
	locks     map[file.BlockId]string
}

func NewConcurrencyManager() *ConcurrencyManager {
	return &ConcurrencyManager{
		lockTable: lockTableInstance,
		locks:     map[file.BlockId]string{},
	}
}

// Obtain an SLock on the block, if necessary.
// The method will ask the lock table for an SLock
// if the transaction currently has no locks on that block.
func (cm *ConcurrencyManager) SLock(blockId file.BlockId) {
	if _, exists := cm.locks[blockId]; !exists {
		cm.lockTable.SLock(blockId)
		cm.locks[blockId] = "S"
	}
}

// Obtain an XLock on the block, if necessary.
// If the transaction does not have an XLock on that block,
// then the method first gets an SLock on that block
// (if necessary), and then upgrades it to an XLock.
func (cm *ConcurrencyManager) XLock(blockId file.BlockId) {
	if !cm.hasXLock(blockId) {
		cm.SLock(blockId)
		cm.lockTable.XLock(blockId)
		cm.locks[blockId] = "X"
	}
}

// Release all locks by asking the lock table to
// unlock each one.
func (cm *ConcurrencyManager) Release() {
	for blockId := range cm.locks {
		cm.lockTable.Unlock(blockId)
	}
	clear(cm.locks)
}

func (cm *ConcurrencyManager) hasXLock(blockId file.BlockId) bool {
	lockType, exists := cm.locks[blockId]
	if !exists {
		return false
	}
	return lockType == "X"
}
