package plan

import (
	"fmt"

	"github.com/evanxg852000/simpledb/internal/metadata"
	"github.com/evanxg852000/simpledb/internal/parser"
	"github.com/evanxg852000/simpledb/internal/query"
	"github.com/evanxg852000/simpledb/internal/tx/recovery"
)

// The simplest, most naive query planner possible.
type BasicQueryPlanner struct {
	mdtManager *metadata.MetadataManager
}

func NewBasicQueryPlanner(mdtManager *metadata.MetadataManager) *BasicQueryPlanner {
	return &BasicQueryPlanner{mdtManager}
}

// Creates a query plan as follows.  It first takes
// the product of all tables and views; it then selects on the predicate;
// and finally it projects on the field list.
func (bqp *BasicQueryPlanner) CreatePlan(selectStmt parser.SelectStmt, tx *recovery.Transaction) Plan {
	// Step 1: Create a plan for each mentioned table or view.
	plans := make([]Plan, 0)
	for _, tableName := range selectStmt.Tables {
		viewDef, err := bqp.mdtManager.GetViewDef(tableName, tx)
		if err != nil {
			panic(fmt.Sprint("error fetching view definition", err))
		}
		if viewDef != "" {
			// Recursively plan the view.
			ast := parser.ParseQuery(viewDef)
			stmts := ast.([]any)
			viewStmt := stmts[0].(parser.SelectStmt)
			plans = append(plans, bqp.CreatePlan(viewStmt, tx))
		} else {
			plan := NewTablePlan(tx, tableName, bqp.mdtManager)
			plans = append(plans, plan)
		}
	}

	// Step 2: Create the product of all table plans
	plan := plans[0]
	plans = plans[1:]
	for _, nextPlan := range plans {
		plan = NewProductPlan(plan, nextPlan)
	}

	// Step 3: Add a selection plan for the predicate
	predicate := query.NewPredicate(selectStmt.Condition)
	plan = NewSelectPlan(plan, predicate)

	// Step 4: Project on the field names
	plan = NewProjectPlan(plan, selectStmt.Fields)

	return plan
}
