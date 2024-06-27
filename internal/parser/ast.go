package parser

// "github.com/evanxg852000/simpledb/internal/record"

const (
	_ = iota
	INTEGER_TYPE
	STRING_TYPE
)

type Literal struct {
	Value any
}

func (c *Literal) Type() int64 {
	if _, ok := c.Value.(string); ok {
		return STRING_TYPE
	}
	return INTEGER_TYPE
}

func (c *Literal) AsInt() int {
	return c.Value.(int)
}

func (c *Literal) AsString() string {
	return c.Value.(string)
}

// type FieldIdent struct {
// 	name string
// }

type Condition struct {
	Left  Term
	Op    string
	Right Term
}

type Term struct {
	Left  Expr
	Op    string
	Right Expr
}

type Expr struct {
	Value any // FieldName or Literal
}

func (e *Expr) IsFieldName() bool {
	if _, ok := e.Value.(string); ok {
		return true
	}
	return false
}

func (e *Expr) AsFieldExpr() string {
	return e.Value.(string)
}

func (e *Expr) AsLiteralExpr() Literal {
	return e.Value.(Literal)
}

type InsertStmt struct {
	Table  string
	Fields []string
	Values []Literal
}

type SelectStmt struct {
	Fields    []string
	Tables    []string
	Condition Condition
}

type UpdateExpr struct {
	Field string
	Value Expr
}

type UpdateStmt struct {
	Table     string
	Exprs     []UpdateExpr
	Condition Condition
}

type DeleteStmt struct {
	Table     string
	Condition Condition
}

type TypeSpec struct {
	DataType int64
	Length   int64
}

type FieldSpec struct {
	Name string
	Spec TypeSpec
}

type CreateTableStmt struct {
	Table  string
	Fields []FieldSpec
}

type CreateViewStmt struct {
	Name     string
	Query    SelectStmt
	QueryStr string
}

type CreateIndexStmt struct {
	Name  string
	Table string
	Field string
}
