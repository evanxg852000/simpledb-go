package metadata

import (
	"github.com/evanxg852000/simpledb/internal/record"
	"github.com/evanxg852000/simpledb/internal/tx/recovery"
)

const (
	MAX_VIEW_DEF = 100

	VIEW_CATALOG = "view_catalog"
)

type ViewManager struct {
	tableManager *TableManager
}

func NewViewManager(isNew bool, tableManager *TableManager, tx *recovery.Transaction) *ViewManager {
	viewManager := &ViewManager{tableManager}
	if isNew {
		schema := record.NewSchema()
		schema.AddStringField("view_name", MAX_NAME_LENGTH)
		schema.AddStringField("view_def", MAX_VIEW_DEF)
		tableManager.CreateTable("view_catalog", schema, tx)
	}
	return viewManager
}

func (vm *ViewManager) CreateView(vName string, viewDef string, tx *recovery.Transaction) error {
	layout, err := vm.tableManager.GetLayout("view_catalog", tx)
	if err != nil {
		return err
	}
	tableScan, err := record.NewTableScan(tx, VIEW_CATALOG, layout)
	if err != nil {
		return err
	}
	tableScan.Insert()
	tableScan.SetString("view_name", vName)
	tableScan.SetString("view_def", viewDef)
	tableScan.Close()
	return nil
}

func (vm *ViewManager) GetViewDef(vName string, tx *recovery.Transaction) (string, error) {
	viewDef := ""
	layout, err := vm.tableManager.GetLayout("view_catalog", tx)
	if err != nil {
		return viewDef, err
	}
	tableScan, err := record.NewTableScan(tx, VIEW_CATALOG, layout)
	if err != nil {
		return viewDef, err
	}

	for tableScan.Next() {
		storedViewName := tableScan.GetString("view_name")
		if storedViewName == vName {
			viewDef = tableScan.GetString("view_def")
			break
		}
	}
	tableScan.Close()
	return viewDef, nil
}
