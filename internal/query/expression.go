package query

import (
	"fmt"

	"github.com/evanxg852000/simpledb/internal/parser"
)

type Expression struct {
	inner parser.Expr
}

func NewExpression(inner parser.Expr) *Expression {
	return &Expression{inner}
}

// Evaluate the expression with respect to the
// current record of the specified scan.
func (expr *Expression) Evaluate(scan Scan) Constant {
	if !expr.inner.IsFieldName() {
		return NewConstant(expr.inner.AsLiteralExpr().Value)
	}

	return scan.GetValue(expr.inner.AsFieldExpr())
}

// Return true if the expression is a field reference.
func (expr *Expression) IsFieldName() bool {
	return expr.inner.IsFieldName()
}

// Return the constant corresponding to a constant expression,
// or null if the expression does not
// denote a constant.
func (expr *Expression) AsConstant() Constant {
	return NewConstant(expr.inner.AsLiteralExpr().Value)
}

// Return the field name corresponding to a constant expression,
// or null if the expression does not
// denote a field.
func (expr *Expression) AsFieldName() string {
	return expr.inner.AsFieldExpr()
}

func (expr *Expression) toString() string {
	if !expr.inner.IsFieldName() {
		return fmt.Sprintf("%v", expr.inner.AsLiteralExpr().Value)
	}
	return expr.inner.AsFieldExpr()
}
