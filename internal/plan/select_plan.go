package plan

import (
	"github.com/evanxg852000/simpledb/internal/query"
	"github.com/evanxg852000/simpledb/internal/record"
)

// The Plan class corresponding to the select
// relational algebra operator.
type SelectPlan struct {
	plan      Plan
	predicate *query.Predicate
}

// Creates a new select node in the query tree,
// having the specified subquery and predicate.
func NewSelectPlan(plan Plan, predicate *query.Predicate) *SelectPlan {
	return &SelectPlan{plan, predicate}
}

// Creates a select scan for this query.
func (sp *SelectPlan) Open() query.Scan {
	scan := sp.plan.Open()
	return query.NewSelectScan(scan, sp.predicate)
}

// Estimates the number of block accesses in the selection,
// which is the same as in the underlying query.
func (sp *SelectPlan) BlockAccessed() int64 {
	return sp.plan.BlockAccessed()
}

// Estimates the number of records in the table,
// which is obtainable from the statistics manager.
func (sp *SelectPlan) RecordsOutput() int64 {
	return sp.plan.RecordsOutput() / ReductionFactor(sp.predicate, sp.plan)
}

// Estimates the number of distinct field values in the table,
// which is obtainable from the statistics manager.
func (sp *SelectPlan) DistinctValues(fieldName string) int64 {
	//TODO: missing optimal query planner
	return 1
}

// Returns the schema of the selection,
// which is the same as in the underlying query.
func (sp *SelectPlan) Schema() record.Schema {
	return sp.plan.Schema()
}

func ReductionFactor(pred *query.Predicate, p Plan) int64 {
	//TODO: missing optimal query planner
	return 1
}
