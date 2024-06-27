package plan

import (
	"github.com/evanxg852000/simpledb/internal/parser"
	"github.com/evanxg852000/simpledb/internal/tx/recovery"
)

// The interface implemented by the planners
// for SQL insert, delete, and modify statements.
type UpdatePlanner interface {

	// Executes the specified insert statement, and
	// returns the number of affected records.
	ExecuteInsert(stmt parser.InsertStmt, tx *recovery.Transaction) int64

	// Executes the specified delete statement, and
	// returns the number of affected records.
	ExecuteDelete(stmt parser.DeleteStmt, tx *recovery.Transaction) int64

	// Executes the specified modify statement, and
	// returns the number of affected records.
	ExecuteModify(stmt parser.UpdateStmt, tx *recovery.Transaction) int64

	// Executes the specified create table statement, and
	// returns the number of affected records.
	ExecuteCreateTable(stmt parser.CreateTableStmt, tx *recovery.Transaction) int64

	// Executes the specified create view statement, and
	// returns the number of affected records.
	ExecuteCreateView(stmt parser.CreateViewStmt, tx *recovery.Transaction) int64

	// Executes the specified create index statement, and
	// returns the number of affected records.
	ExecuteCreateIndex(stmt parser.CreateIndexStmt, tx *recovery.Transaction) int64
}
