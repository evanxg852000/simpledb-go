// Code generated from SimpleSql.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // SimpleSql
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSimpleSqlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSimpleSqlVisitor) VisitParse(ctx *ParseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitCreate_table_stmt(ctx *Create_table_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitField_specs(ctx *Field_specsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitField_spec(ctx *Field_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitType_spec(ctx *Type_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitVarchar_spec(ctx *Varchar_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitInsert_stmt(ctx *Insert_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitConstant_list(ctx *Constant_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitSelect_stmt(ctx *Select_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitIdent_list(ctx *Ident_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitUpdate_stmt(ctx *Update_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitUpdate_expr_list(ctx *Update_expr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitUpdate_expr(ctx *Update_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitDelete_stmt(ctx *Delete_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitCreate_view_stmt(ctx *Create_view_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitCreate_index_stmt(ctx *Create_index_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
