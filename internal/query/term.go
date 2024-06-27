package query

// A term is a comparison between two expressions.
// type Term struct {
// 	left  *Expression
// 	right *Expression
// }

// // Create a new term that compares two expressions
// // for equality.
// func NewTerm(left *Expression, right *Expression) *Term {
// 	return &Term{left, right}
// }

// // Return true if both of the term's expressions
// // evaluate to the same constant,
// // with respect to the specified scan.
// func (t *Term) IsSatisfied(scan Scan) bool {
// 	left := t.left.Evaluate(scan)
// 	right := t.right.Evaluate(scan)
// 	_ = right
// 	_ = left
// 	//TODO:
// 	return false
// }

// // func (t *Term) ReductionFactor(p plan.Plan) int {
// // 	//TODO:
// // 	return 2
// // }

// // Determine if this term is of the form "F=c"
// // where F is the specified field and c is some constant.
// // If so, the method returns that constant.
// // If not, the method returns null.
// func (t *Term) EquatesWithConstant(fieldName string) *Constant {
// 	//TODO:
// 	return nil
// }

// func (t *Term) ToString(fieldName string) string {
// 	return fmt.Sprintf("%s = %s ", t.left.toString(), t.right.toString())
// }
