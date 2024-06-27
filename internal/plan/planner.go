package plan

import (
	"fmt"

	"github.com/evanxg852000/simpledb/internal/parser"
	"github.com/evanxg852000/simpledb/internal/tx/recovery"
)

type Planner struct {
	queryPlanner  QueryPlanner
	updatePlanner UpdatePlanner
}

// The object that executes SQL statements.
func NewPlanner(queryPlanner QueryPlanner, updatePlanner UpdatePlanner) *Planner {
	return &Planner{queryPlanner, updatePlanner}
}

// Parses an SQL statement,
// - returns a plan for data selection queries
// - execute and returns affected rows for modification queries
func (planner *Planner) ExecuteQuery(queryStr string, tx *recovery.Transaction) (any, error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("ERR:", err)
		}
	}()

	ast := parser.ParseQuery(queryStr)
	sqlStmt := ast.([]any)[0]
	err := planner.VerifyStatement(sqlStmt)
	if err != nil {
		return nil, err
	}

	switch stmt := sqlStmt.(type) {
	case parser.SelectStmt:
		return planner.queryPlanner.CreatePlan(stmt, tx), nil
	case parser.InsertStmt:
		return planner.updatePlanner.ExecuteInsert(stmt, tx), nil
	case parser.UpdateStmt:
		return planner.updatePlanner.ExecuteModify(stmt, tx), nil
	case parser.DeleteStmt:
		return planner.updatePlanner.ExecuteDelete(stmt, tx), nil
	case parser.CreateTableStmt:
		return planner.updatePlanner.ExecuteCreateTable(stmt, tx), nil
	case parser.CreateViewStmt:
		return planner.updatePlanner.ExecuteCreateView(stmt, tx), nil
	case parser.CreateIndexStmt:
		return planner.updatePlanner.ExecuteCreateIndex(stmt, tx), nil
	}
	return nil, fmt.Errorf("unknown SQL statement")
}

func (planner *Planner) VerifyStatement(sqlStmt any) error {
	//TODO: verify statement and return error if any
	return nil
}
