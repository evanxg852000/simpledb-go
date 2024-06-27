package plan

import (
	"github.com/evanxg852000/simpledb/internal/metadata"
	"github.com/evanxg852000/simpledb/internal/parser"
	"github.com/evanxg852000/simpledb/internal/query"
	"github.com/evanxg852000/simpledb/internal/record"
	"github.com/evanxg852000/simpledb/internal/tx/recovery"
)

// The basic planner for SQL update statements.
type BasicUpdatePlanner struct {
	mdtManager *metadata.MetadataManager
}

func NewBasicUpdatePlanner(mdtManager *metadata.MetadataManager) *BasicUpdatePlanner {
	return &BasicUpdatePlanner{mdtManager}
}

func (bup *BasicUpdatePlanner) ExecuteInsert(stmt parser.InsertStmt, tx *recovery.Transaction) int64 {
	plan := NewTablePlan(tx, stmt.Table, bup.mdtManager)
	updateScan := plan.Open().(query.UpdateScan)

	updateScan.Insert()
	for i, fieldName := range stmt.Fields {
		value := query.NewConstant(stmt.Values[i].Value)
		updateScan.SetValue(fieldName, value)
	}
	updateScan.Close()
	return 1
}

func (bup *BasicUpdatePlanner) ExecuteDelete(stmt parser.DeleteStmt, tx *recovery.Transaction) int64 {
	var plan Plan
	plan = NewTablePlan(tx, stmt.Table, bup.mdtManager)
	plan = NewSelectPlan(plan, query.NewPredicate(stmt.Condition))
	updateScan := plan.Open().(query.UpdateScan)
	count := 0
	for updateScan.Next() {
		updateScan.Delete()
		count += 1
	}
	updateScan.Close()
	return int64(count)
}

func (bup *BasicUpdatePlanner) ExecuteModify(stmt parser.UpdateStmt, tx *recovery.Transaction) int64 {
	var plan Plan
	plan = NewTablePlan(tx, stmt.Table, bup.mdtManager)
	plan = NewSelectPlan(plan, query.NewPredicate(stmt.Condition))
	updateScan := plan.Open().(query.UpdateScan)
	count := 0
	for updateScan.Next() {
		for _, updateExpr := range stmt.Exprs {
			expr := query.NewExpression(updateExpr.Value)
			value := expr.Evaluate(updateScan)
			updateScan.SetValue(updateExpr.Field, value)
		}
		count += 1
	}
	updateScan.Close()
	return int64(count)
}

func (bup *BasicUpdatePlanner) ExecuteCreateTable(stmt parser.CreateTableStmt, tx *recovery.Transaction) int64 {
	schema := record.NewSchema()
	for _, fieldDesc := range stmt.Fields {
		fieldSpec := fieldDesc.Spec
		schema.AddField(fieldDesc.Name, fieldSpec.DataType, fieldSpec.Length)
	}
	err := bup.mdtManager.CreateTable(stmt.Table, schema, tx)
	if err != nil {
		panic(err)
	}
	return 0
}

func (bup *BasicUpdatePlanner) ExecuteCreateView(stmt parser.CreateViewStmt, tx *recovery.Transaction) int64 {
	err := bup.mdtManager.CreateView(stmt.Name, stmt.QueryStr, tx)
	if err != nil {
		panic(err)
	}
	return 0
}

func (bup *BasicUpdatePlanner) ExecuteCreateIndex(stmt parser.CreateIndexStmt, tx *recovery.Transaction) int64 {
	err := bup.mdtManager.CreateIndex(stmt.Name, stmt.Table, stmt.Field, tx)
	if err != nil {
		panic(err)
	}
	return 0
}
