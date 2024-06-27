package metadata

import (
	"github.com/evanxg852000/simpledb/internal/record"
	"github.com/evanxg852000/simpledb/internal/tx/recovery"
)

type MetadataManager struct {
	tableManager *TableManager
	viewManager  *ViewManager
	statsManager *StatsManager
	indexManager *IndexManager
}

func NewMetadataManager(isNew bool, tx *recovery.Transaction) *MetadataManager {
	tableManager := NewTableManager(isNew, tx)
	viewManager := NewViewManager(isNew, tableManager, tx)
	statsManager := NewStatsManager(tableManager, tx)
	indexManager := NewIndexManager(isNew, tableManager, statsManager, tx)
	return &MetadataManager{
		tableManager,
		viewManager,
		statsManager,
		indexManager,
	}
}

func (mdtManager *MetadataManager) CreateTable(tblName string, schema *record.Schema, tx *recovery.Transaction) error {
	return mdtManager.tableManager.CreateTable(tblName, schema, tx)
}

func (mdtManager *MetadataManager) GetLayout(tblName string, tx *recovery.Transaction) (*record.Layout, error) {
	return mdtManager.tableManager.GetLayout(tblName, tx)
}

func (mdtManager *MetadataManager) CreateView(viewName, viewDef string, tx *recovery.Transaction) error {
	return mdtManager.viewManager.CreateView(viewName, viewDef, tx)
}

func (mdtManager *MetadataManager) GetViewDef(viewName string, tx *recovery.Transaction) (string, error) {
	return mdtManager.viewManager.GetViewDef(viewName, tx)
}

func (mdtManager *MetadataManager) CreateIndex(idxName, tblName, fldName string, tx *recovery.Transaction) error {
	return mdtManager.indexManager.CreateIndex(idxName, tblName, fldName, tx)
}

func (mdtManager *MetadataManager) GetIndexInfo(tblName string, tx *recovery.Transaction) (map[string]IndexInfo, error) {
	return mdtManager.indexManager.GetIndexInfo(tblName, tx)
}

func (mdtManager *MetadataManager) GetStatInfo(tblName string, layout *record.Layout, tx *recovery.Transaction) StatInfo {
	return mdtManager.statsManager.GetStatInfo(tblName, layout, tx)
}

func (mdtManager *MetadataManager) GetTableManager() *TableManager {
	return mdtManager.tableManager
}

func (mdtManager *MetadataManager) GetViewManager() *ViewManager {
	return mdtManager.viewManager
}

func (mdtManager *MetadataManager) GetStatsManager() *StatsManager {
	return mdtManager.statsManager
}

func (mdtManager *MetadataManager) GetIndexManager() *IndexManager {
	return mdtManager.indexManager
}
