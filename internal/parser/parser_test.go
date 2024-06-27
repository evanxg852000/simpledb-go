package parser_test

import (
	"testing"

	"github.com/evanxg852000/simpledb/internal/parser"
	"github.com/evanxg852000/simpledb/internal/record"
	"github.com/stretchr/testify/assert"
)

func TestParseCreateTableStmt(t *testing.T) {
	assert := assert.New(t)
	input := "create table foo(a int, b varchar(4), c int)"
	ast := parser.ParseQuery(input)

	stmts := ast.([]any)
	assert.Equal(len(stmts), 1)

	createStmt := stmts[0].(parser.CreateTableStmt)
	assert.Equal(parser.CreateTableStmt{"foo", []parser.FieldSpec{
		{"a", parser.TypeSpec{record.INTEGER_TYPE, 0}},
		{"b", parser.TypeSpec{record.STRING_TYPE, 4}},
		{"c", parser.TypeSpec{record.INTEGER_TYPE, 0}},
	}}, createStmt)
}

func TestParseInsertStmt(t *testing.T) {
	assert := assert.New(t)
	input := "insert into foo(a, b) values (2, 'evan')"
	ast := parser.ParseQuery(input)

	stmts := ast.([]any)
	assert.Equal(len(stmts), 1)

	insertStmt := stmts[0].(parser.InsertStmt)
	assert.Equal(parser.InsertStmt{
		"foo",
		[]string{"a", "b"},
		[]parser.Literal{{int64(2)}, {"evan"}},
	}, insertStmt)
}

func TestParseSelectStmt(t *testing.T) {
	assert := assert.New(t)
	input := "select a, b from foo where a=1"
	ast := parser.ParseQuery(input)

	stmts := ast.([]any)
	assert.Equal(len(stmts), 1)

	selectStmt := stmts[0].(parser.SelectStmt)
	assert.Equal(parser.SelectStmt{
		[]string{"a", "b"},
		[]string{"foo"},
		parser.Condition{
			Left: parser.Term{
				parser.Expr{"a"},
				"=",
				parser.Expr{parser.Literal{int64(1)}},
			},
			Op:    "",
			Right: parser.Term{},
		},
	}, selectStmt)
}

func TestParseUpdateStmt(t *testing.T) {
	assert := assert.New(t)
	input := "update foo set a=2, b=1 where a=1 or b != 2"
	ast := parser.ParseQuery(input)

	stmts := ast.([]any)
	assert.Equal(len(stmts), 1)

	updateStmt := stmts[0].(parser.UpdateStmt)
	assert.Equal(parser.UpdateStmt{
		"foo",
		[]parser.UpdateExpr{
			{"a", parser.Expr{parser.Literal{int64(2)}}},
			{"b", parser.Expr{parser.Literal{int64(1)}}},
		},
		parser.Condition{
			Left: parser.Term{
				parser.Expr{"a"},
				"=",
				parser.Expr{parser.Literal{int64(1)}},
			},
			Op: "or",
			Right: parser.Term{
				parser.Expr{"b"},
				"!=",
				parser.Expr{parser.Literal{int64(2)}},
			},
		},
	}, updateStmt)
}

func TestParseDeleteStmt(t *testing.T) {
	assert := assert.New(t)
	input := "delete from foo where a=23 and f!=100"
	ast := parser.ParseQuery(input)

	stmts := ast.([]any)
	assert.Equal(len(stmts), 1)

	deleteStmt := stmts[0].(parser.DeleteStmt)
	assert.Equal(parser.DeleteStmt{
		"foo",
		parser.Condition{
			Left: parser.Term{
				parser.Expr{"a"},
				"=",
				parser.Expr{parser.Literal{int64(23)}},
			},
			Op: "and",
			Right: parser.Term{
				parser.Expr{"f"},
				"!=",
				parser.Expr{parser.Literal{int64(100)}},
			},
		},
	}, deleteStmt)
}

func TestParseCreateViewStmt(t *testing.T) {
	assert := assert.New(t)
	input := "create view view1 as select * from foo where a=23"
	ast := parser.ParseQuery(input)

	stmts := ast.([]any)
	assert.Equal(len(stmts), 1)

	createViewStmt := stmts[0].(parser.CreateViewStmt)
	assert.Equal(parser.CreateViewStmt{
		"view1",
		parser.SelectStmt{
			[]string{"*"},
			[]string{"foo"},
			parser.Condition{
				Left: parser.Term{
					parser.Expr{"a"},
					"=",
					parser.Expr{parser.Literal{int64(23)}},
				},
				Op:    "",
				Right: parser.Term{},
			},
		},
		"select*fromfoowherea=23",
	}, createViewStmt)
}

func TestParseCreateIndexStmt(t *testing.T) {
	assert := assert.New(t)
	input := "create index index_b on foo (b)"
	ast := parser.ParseQuery(input)

	stmts := ast.([]any)
	assert.Equal(len(stmts), 1)

	createIndexStmt := stmts[0].(parser.CreateIndexStmt)
	assert.Equal(parser.CreateIndexStmt{"index_b", "foo", "b"}, createIndexStmt)
}
