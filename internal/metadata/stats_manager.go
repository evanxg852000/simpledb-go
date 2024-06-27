package metadata

import (
	"sync"

	"github.com/evanxg852000/simpledb/internal/record"
	"github.com/evanxg852000/simpledb/internal/tx/recovery"
)

// The statistics manager is responsible for
// keeping statistical information about each table.
// The manager does not store this information in the database.
// Instead, it calculates this information on system startup,
// and periodically refreshes it.
type StatsManager struct {
	tableManager TableManager
	tableStats   map[string]StatInfo
	numCalls     int
	mu           *sync.Mutex
}

// Create the statistics manager.
// The initial statistics are calculated by
// traversing the entire database.
func NewStatsManager(tableManager *TableManager, tx *recovery.Transaction) *StatsManager {
	statsManager := &StatsManager{
		tableManager: *tableManager,
		tableStats:   make(map[string]StatInfo),
		numCalls:     0,
		mu:           new(sync.Mutex),
	}
	statsManager.mu.Lock()
	statsManager.refreshStatistics(tx)
	statsManager.mu.Unlock()
	return statsManager
}

// Return the statistical information about the specified table.
func (statsManager *StatsManager) GetStatInfo(tblName string, layout *record.Layout, tx *recovery.Transaction) StatInfo {
	statsManager.mu.Lock()
	defer statsManager.mu.Unlock()

	statsManager.numCalls += 1
	if statsManager.numCalls > 100 {
		statsManager.refreshStatistics(tx)
	}

	si, exist := statsManager.tableStats[tblName]
	if !exist {
		si, _ = statsManager.calcTableStats(tblName, layout, tx)
		statsManager.tableStats[tblName] = si
	}
	return si
}

func (statsManager *StatsManager) refreshStatistics(tx *recovery.Transaction) error {
	statsManager.tableStats = map[string]StatInfo{}
	statsManager.numCalls = 0
	layout, err := statsManager.tableManager.GetLayout(TABLE_CATALOG, tx)
	if err != nil {
		return err
	}

	tableScan, err := record.NewTableScan(tx, TABLE_CATALOG, layout)
	if err != nil {
		return err
	}

	for tableScan.Next() {
		tblName := tableScan.GetString("table_name")
		layout, _ := statsManager.tableManager.GetLayout(tblName, tx)
		si, _ := statsManager.calcTableStats(tblName, layout, tx)
		statsManager.tableStats[tblName] = si
	}
	tableScan.Close()
	return nil
}

func (statsManager *StatsManager) calcTableStats(tblName string, layout *record.Layout, tx *recovery.Transaction) (StatInfo, error) {

	numBlocks := int64(0)
	numRecords := int64(0)
	tableScan, err := record.NewTableScan(tx, tblName, layout)
	if err != nil {
		return StatInfo{}, err
	}

	for tableScan.Next() {
		numRecords += 1
		numBlocks = tableScan.GetRID().BlockNum + 1
	}
	tableScan.Close()
	return NewStatInfo(numBlocks, numRecords), nil
}
