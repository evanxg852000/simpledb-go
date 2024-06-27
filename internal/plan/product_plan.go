package plan

import (
	"github.com/evanxg852000/simpledb/internal/query"
	"github.com/evanxg852000/simpledb/internal/record"
)

type ProductPlan struct {
	left   Plan
	right  Plan
	schema *record.Schema
}

// Creates a new product node in the query tree,
// having the two specified sub-queries.
func NewProductPlan(left, right Plan) *ProductPlan {
	schema := record.NewSchema()
	schema.AddAll(left.Schema())
	schema.AddAll(right.Schema())
	return &ProductPlan{left, right, schema}
}

// Creates a product scan for this query.
func (pp *ProductPlan) Open() query.Scan {
	leftScan := pp.left.Open()
	rightScan := pp.right.Open()
	return query.NewProductScan(leftScan, rightScan)
}

// Estimates the number of block accesses in the product.
// The formula is:
// B(product(p1,p2)) = B(p1) + R(p1)*B(p2)
func (pp *ProductPlan) BlockAccessed() int64 {
	return pp.left.BlockAccessed() + (pp.left.BlockAccessed() * pp.right.BlockAccessed())
}

// Estimates the number of output records in the product.
// The formula is:
// R(product(p1,p2)) = R(p1)*R(p2)
func (pp *ProductPlan) RecordsOutput() int64 {
	return pp.left.RecordsOutput() * pp.right.RecordsOutput()
}

// Estimates the distinct number of field values in the product.
// Since the product does not increase or decrease field values,
// the estimate is the same as in the appropriate underlying query.
func (pp *ProductPlan) DistinctValues(fieldName string) int64 {
	leftSchema := pp.left.Schema()
	if leftSchema.HasField(fieldName) {
		return pp.left.DistinctValues(fieldName)
	}
	return pp.right.DistinctValues(fieldName)
}

// Returns the schema of the product,
// which is the union of the schemas of the underlying queries.
func (pp *ProductPlan) Schema() record.Schema {
	return *pp.schema
}
