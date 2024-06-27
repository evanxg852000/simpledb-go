package plan

import (
	"github.com/evanxg852000/simpledb/internal/query"
	"github.com/evanxg852000/simpledb/internal/record"
)

type ProjectPlan struct {
	plan   Plan
	schema *record.Schema
}

// Creates a new project node in the query tree,
// having the specified subquery and field list.
func NewProjectPlan(plan Plan, fields []string) *ProjectPlan {
	schema := record.NewSchema()
	if fields[0] == "*" {
		schema.AddAll(plan.Schema())
		fields = fields[1:]
	}
	for _, field := range fields {
		schema.Add(field, plan.Schema())
	}
	return &ProjectPlan{plan, schema}
}

// Creates a project scan for this query.
func (pp *ProjectPlan) Open() query.Scan {
	scan := pp.plan.Open()
	return query.NewProjectScan(scan, pp.schema.Fields())
}

// Estimates the number of block accesses in the projection,
// which is the same as in the underlying query.
func (pp *ProjectPlan) BlockAccessed() int64 {
	return pp.plan.BlockAccessed()
}

// Estimates the number of output records in the projection,
// which is the same as in the underlying query.
func (pp *ProjectPlan) RecordsOutput() int64 {
	return pp.plan.RecordsOutput()
}

// Estimates the number of distinct field values
// in the projection,
// which is the same as in the underlying query.
func (pp *ProjectPlan) DistinctValues(fieldName string) int64 {
	return pp.plan.DistinctValues(fieldName)
}

// Returns the schema of the projection,
// which is taken from the field list.
func (pp *ProjectPlan) Schema() record.Schema {
	return *pp.schema
}
