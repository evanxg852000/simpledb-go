package parser

import (
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func ParseQuery(input string) interface{} {
	is := antlr.NewInputStream(input)
	lexer := NewSimpleSqlLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewSimpleSqlParser(tokens)
	visitor := NewSimpleSqlAstBuilder()
	return parser.Parse().Accept(visitor)
}

type SimpleSqlAstBuilder struct {
	*antlr.BaseParseTreeVisitor
}

func NewSimpleSqlAstBuilder() *SimpleSqlAstBuilder {
	return &SimpleSqlAstBuilder{}
}

func (v *SimpleSqlAstBuilder) VisitParse(ctx *ParseContext) interface{} {
	stmtListCtx := ctx.StatementList(0)
	return v.VisitStatementList(stmtListCtx.(*StatementListContext))
}

func (v *SimpleSqlAstBuilder) VisitStatementList(ctx *StatementListContext) interface{} {
	statements := make([]any, 0)
	nodes := ctx.AllStatement()
	for _, node := range nodes {
		statements = append(statements, v.VisitStatement(node.(*StatementContext)))
	}
	return statements
}

func (v *SimpleSqlAstBuilder) VisitStatement(ctx *StatementContext) interface{} {
	if stmt := ctx.Create_table_stmt(); stmt != nil {
		return v.VisitCreate_table_stmt(stmt.(*Create_table_stmtContext))
	}

	if stmt := ctx.Insert_stmt(); stmt != nil {
		return v.VisitInsert_stmt(stmt.(*Insert_stmtContext))
	}

	if stmt := ctx.Select_stmt(); stmt != nil {
		return v.VisitSelect_stmt(stmt.(*Select_stmtContext))
	}

	if stmt := ctx.Update_stmt(); stmt != nil {
		return v.VisitUpdate_stmt(stmt.(*Update_stmtContext))
	}

	if stmt := ctx.Delete_stmt(); stmt != nil {
		return v.VisitDelete_stmt(stmt.(*Delete_stmtContext))
	}

	if stmt := ctx.Create_view_stmt(); stmt != nil {
		return v.VisitCreate_view_stmt(stmt.(*Create_view_stmtContext))
	}

	if stmt := ctx.Create_index_stmt(); stmt != nil {
		return v.VisitCreate_index_stmt(stmt.(*Create_index_stmtContext))
	}

	return v.VisitChildren(ctx)
}

func (v *SimpleSqlAstBuilder) VisitCreate_table_stmt(ctx *Create_table_stmtContext) interface{} {
	if ctx.CREATE_() == nil {
		return nil
	}

	if ctx.TABLE_() == nil {
		return nil
	}

	tableName := ctx.IDENT().GetText()
	fieldSpecs := v.VisitField_specs(ctx.Field_specs().(*Field_specsContext))
	return CreateTableStmt{tableName, fieldSpecs.([]FieldSpec)}
}

func (v *SimpleSqlAstBuilder) VisitField_specs(ctx *Field_specsContext) interface{} {
	listCtxs := ctx.AllField_spec()
	fieldSpec := v.VisitField_spec(listCtxs[0].(*Field_specContext)).(FieldSpec)
	fieldSpecs := []FieldSpec{fieldSpec}
	for i, fieldSpecCtx := range listCtxs[1:] {
		if ctx.COMMA(i) == nil {
			return nil
		}
		fieldSpec := v.VisitField_spec(fieldSpecCtx.(*Field_specContext)).(FieldSpec)
		fieldSpecs = append(fieldSpecs, fieldSpec)
	}
	return fieldSpecs
}

func (v *SimpleSqlAstBuilder) VisitField_spec(ctx *Field_specContext) interface{} {
	field := ctx.IDENT().GetText()
	typeSpec := v.VisitType_spec(ctx.Type_spec().(*Type_specContext))
	return FieldSpec{field, typeSpec.(TypeSpec)}
}

func (v *SimpleSqlAstBuilder) VisitType_spec(ctx *Type_specContext) interface{} {
	if ctx.INT_() != nil {
		return TypeSpec{INTEGER_TYPE, 0}
	}

	typeSpec := v.VisitVarchar_spec(ctx.Varchar_spec().(*Varchar_specContext))
	return typeSpec
}

func (v *SimpleSqlAstBuilder) VisitVarchar_spec(ctx *Varchar_specContext) interface{} {
	if ctx.VAR_CHAR_() == nil {
		return nil
	}
	length, err := strconv.ParseInt(ctx.INT_LITERAL().GetText(), 10, 64)
	if err != nil {
		return nil
	}
	return TypeSpec{STRING_TYPE, length}
}

func (v *SimpleSqlAstBuilder) VisitInsert_stmt(ctx *Insert_stmtContext) interface{} {
	if ctx.INSERT_() == nil {
		return nil
	}

	if ctx.INTO_() == nil {
		return nil
	}

	tableName := ctx.IDENT().GetText()

	fields := v.VisitIdent_list(ctx.Ident_list().(*Ident_listContext)).([]string)
	fieldList := make([]string, 0, len(fields))
	fieldList = append(fieldList, fields...)

	if ctx.VALUES_() == nil {
		return nil
	}

	valueList := v.VisitConstant_list(ctx.Constant_list().(*Constant_listContext)).([]Literal)
	return InsertStmt{tableName, fieldList, valueList}
}

func (v *SimpleSqlAstBuilder) VisitConstant_list(ctx *Constant_listContext) interface{} {
	listCtxs := ctx.AllLiteral()
	literal := v.VisitLiteral(listCtxs[0].(*LiteralContext)).(Literal)
	literalList := []Literal{literal}
	for i, litCtx := range listCtxs[1:] {
		if ctx.COMMA(i) == nil {
			return nil
		}
		literal := v.VisitLiteral(litCtx.(*LiteralContext)).(Literal)
		literalList = append(literalList, literal)
	}
	return literalList
}

func (v *SimpleSqlAstBuilder) VisitSelect_stmt(ctx *Select_stmtContext) interface{} {
	if ctx.SELECT_() == nil {
		return nil
	}

	fieldList := make([]string, 0)
	identIdx := 0
	if startSelect := ctx.STAR(); startSelect != nil {
		fieldList = append(fieldList, "*")
	} else {
		identListCtx := ctx.Ident_list(identIdx)
		fields := v.VisitIdent_list(identListCtx.(*Ident_listContext)).([]string)
		fieldList = append(fieldList, fields...)
		identIdx += 1
	}

	if ctx.FROM_() == nil {
		return nil
	}

	identListCtx := ctx.Ident_list(identIdx)
	tableList := v.VisitIdent_list(identListCtx.(*Ident_listContext)).([]string)

	condition := Condition{}
	if ctx.WHERE_() != nil {
		condition = v.VisitCondition(ctx.Condition().(*ConditionContext)).(Condition)
	}

	return SelectStmt{fieldList, tableList, condition}
}

func (v *SimpleSqlAstBuilder) VisitIdent_list(ctx *Ident_listContext) interface{} {
	list := ctx.AllIDENT()
	identList := []string{list[0].GetText()}
	for i, item := range list[1:] {
		if ctx.COMMA(i) == nil {
			return nil
		}
		identList = append(identList, item.GetText())
	}
	return identList
}

func (v *SimpleSqlAstBuilder) VisitUpdate_stmt(ctx *Update_stmtContext) interface{} {
	if ctx.UPDATE_() == nil {
		return nil
	}

	tableName := ctx.IDENT().GetText()

	if ctx.SET_() == nil {
		return nil
	}

	updateExprCtx := ctx.Update_expr_list()
	updateExprs := v.VisitUpdate_expr_list(updateExprCtx.(*Update_expr_listContext)).([]UpdateExpr)

	condition := Condition{}
	if ctx.WHERE_() != nil {
		condition = v.VisitCondition(ctx.Condition().(*ConditionContext)).(Condition)
	}
	return UpdateStmt{tableName, updateExprs, condition}
}

func (v *SimpleSqlAstBuilder) VisitUpdate_expr_list(ctx *Update_expr_listContext) interface{} {
	updateExprCtxs := ctx.AllUpdate_expr()

	expr := v.VisitUpdate_expr(updateExprCtxs[0].(*Update_exprContext)).(UpdateExpr)
	updateExprList := []UpdateExpr{expr}
	for i, exprCtx := range updateExprCtxs[1:] {
		if ctx.COMMA(i) == nil {
			return nil
		}
		expr := v.VisitUpdate_expr(exprCtx.(*Update_exprContext)).(UpdateExpr)
		updateExprList = append(updateExprList, expr)
	}

	return updateExprList
}

func (v *SimpleSqlAstBuilder) VisitUpdate_expr(ctx *Update_exprContext) interface{} {
	field := ctx.IDENT().GetText()
	if ctx.EQUAL() == nil {
		return nil
	}

	expr := v.VisitExpression(ctx.Expression().(*ExpressionContext))
	return UpdateExpr{field, expr.(Expr)}
}

func (v *SimpleSqlAstBuilder) VisitDelete_stmt(ctx *Delete_stmtContext) interface{} {
	tableName := ctx.IDENT().GetText()

	condition := Condition{}
	if ctx.WHERE_() != nil {
		condition = v.VisitCondition(ctx.Condition().(*ConditionContext)).(Condition)
	}
	return DeleteStmt{tableName, condition}
}

func (v *SimpleSqlAstBuilder) VisitCreate_view_stmt(ctx *Create_view_stmtContext) interface{} {
	if ctx.CREATE_() == nil {
		return nil
	}

	if ctx.VIEW_() == nil {
		return nil
	}

	tableName := ctx.IDENT().GetText()

	if ctx.AS_() == nil {
		return nil
	}

	selectStmt := v.VisitSelect_stmt(ctx.Select_stmt().(*Select_stmtContext))
	selectStmtText := ctx.Select_stmt().GetText()
	return CreateViewStmt{tableName, selectStmt.(SelectStmt), selectStmtText}
}

func (v *SimpleSqlAstBuilder) VisitCreate_index_stmt(ctx *Create_index_stmtContext) interface{} {
	if ctx.CREATE_() == nil {
		return nil
	}

	if ctx.INDEX_() == nil {
		return nil
	}

	indexName := ctx.IDENT(0).GetText()

	if ctx.ON_() == nil {
		return nil
	}

	tableName := ctx.IDENT(1).GetText()
	field := ctx.IDENT(2).GetText()
	return CreateIndexStmt{indexName, tableName, field}
}

func (v *SimpleSqlAstBuilder) VisitCondition(ctx *ConditionContext) interface{} {
	termCtx := ctx.Term(0)
	left := v.VisitTerm(termCtx.(*TermContext))

	if orTerm := ctx.OR_(); orTerm != nil {
		termCtx = ctx.Term(1)
		right := v.VisitTerm(termCtx.(*TermContext))
		return Condition{left.(Term), "or", right.(Term)}
	}

	if andTerm := ctx.AND_(); andTerm != nil {
		termCtx = ctx.Term(1)
		right := v.VisitTerm(termCtx.(*TermContext))
		return Condition{left.(Term), "and", right.(Term)}
	}
	return Condition{left.(Term), "", Term{}}
}

func (v *SimpleSqlAstBuilder) VisitTerm(ctx *TermContext) interface{} {
	lhs := ctx.left.Accept(v)
	op := ctx.operator.GetText()
	rhs := ctx.right.Accept(v)
	return Term{lhs.(Expr), op, rhs.(Expr)}
}

func (v *SimpleSqlAstBuilder) VisitExpression(ctx *ExpressionContext) interface{} {
	if expr := ctx.IDENT(); expr != nil {
		return Expr{expr.GetText()}
	}
	literal := ctx.Literal().Accept(v)
	return Expr{literal}
}

func (v *SimpleSqlAstBuilder) VisitLiteral(ctx *LiteralContext) interface{} {
	if intLit := ctx.INT_LITERAL(); intLit != nil {
		intValue, _ := strconv.ParseInt(intLit.GetText(), 10, 64)
		return Literal{intValue}
	}
	strLit := ctx.STR_LITERAL().GetText()
	return Literal{strings.Trim(strLit, "'")}
}
