package query

import "github.com/evanxg852000/simpledb/internal/parser"

// A predicate is a Boolean combination of terms.
type Predicate struct {
	Condition parser.Condition
}

func NewPredicate(condition parser.Condition) *Predicate {
	return &Predicate{condition}
}

func (pred *Predicate) IsSatisfied(scan Scan) bool {
	if pred.Condition == (parser.Condition{}) {
		return true
	}

	left := EvaluateTerm(scan, pred.Condition.Left)
	if pred.Condition.Right == (parser.Term{}) {
		return left
	}

	right := EvaluateTerm(scan, pred.Condition.Right)

	switch pred.Condition.Op {
	case "and":
		return left && right
	case "or":
		return left || right
	}
	return false
}

func EvaluateTerm(scan Scan, term parser.Term) bool {
	left := EvaluateExpr(scan, term.Left)
	right := EvaluateExpr(scan, term.Right)
	switch term.Op {
	case "=":
		return left.Equals(right)
	case "!=":
		return !left.Equals(right)
	}
	return false
}

func EvaluateExpr(scan Scan, expr parser.Expr) Constant {
	if expr.IsFieldName() {
		return scan.GetValue(expr.AsFieldExpr())
	}
	return NewConstant(expr.AsLiteralExpr().Value)
}
