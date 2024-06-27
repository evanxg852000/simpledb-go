package plan

import (
	"github.com/evanxg852000/simpledb/internal/parser"
	"github.com/evanxg852000/simpledb/internal/tx/recovery"
)

// The interface implemented by planners for
// the SQL select statement.
type QueryPlanner interface {
	// Creates a plan for the parsed query.
	CreatePlan(selectStmt parser.SelectStmt, tx *recovery.Transaction) Plan
}
