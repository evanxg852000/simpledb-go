// Code generated from SimpleSql.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // SimpleSql
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SimpleSqlParser.
type SimpleSqlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SimpleSqlParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#create_table_stmt.
	VisitCreate_table_stmt(ctx *Create_table_stmtContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#field_specs.
	VisitField_specs(ctx *Field_specsContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#field_spec.
	VisitField_spec(ctx *Field_specContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#type_spec.
	VisitType_spec(ctx *Type_specContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#varchar_spec.
	VisitVarchar_spec(ctx *Varchar_specContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#insert_stmt.
	VisitInsert_stmt(ctx *Insert_stmtContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#constant_list.
	VisitConstant_list(ctx *Constant_listContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#select_stmt.
	VisitSelect_stmt(ctx *Select_stmtContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#ident_list.
	VisitIdent_list(ctx *Ident_listContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#update_stmt.
	VisitUpdate_stmt(ctx *Update_stmtContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#update_expr_list.
	VisitUpdate_expr_list(ctx *Update_expr_listContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#update_expr.
	VisitUpdate_expr(ctx *Update_exprContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#delete_stmt.
	VisitDelete_stmt(ctx *Delete_stmtContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#create_view_stmt.
	VisitCreate_view_stmt(ctx *Create_view_stmtContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#create_index_stmt.
	VisitCreate_index_stmt(ctx *Create_index_stmtContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
