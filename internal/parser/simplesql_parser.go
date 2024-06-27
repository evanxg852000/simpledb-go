// Code generated from SimpleSql.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // SimpleSql
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 32, 197,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 7, 2, 48, 10, 2, 12, 2, 14, 2, 51, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 3, 7, 3, 58, 10, 3, 12, 3, 14, 3, 61, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 5, 4, 70, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 6, 7, 6, 82, 10, 6, 12, 6, 14, 6, 85, 11, 6, 3, 7,
	3, 7, 3, 7, 3, 8, 3, 8, 5, 8, 92, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 106, 10, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 7, 11, 116, 10, 11,
	12, 11, 14, 11, 119, 11, 11, 3, 12, 3, 12, 3, 12, 5, 12, 124, 10, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 5, 12, 130, 10, 12, 3, 13, 3, 13, 3, 13, 7, 13,
	135, 10, 13, 12, 13, 14, 13, 138, 11, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 5, 14, 146, 10, 14, 3, 15, 3, 15, 3, 15, 7, 15, 151, 10, 15,
	12, 15, 14, 15, 154, 11, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 5, 17, 165, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 20, 3, 20, 3, 20, 5, 20, 185, 10, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	22, 3, 22, 5, 22, 193, 10, 22, 3, 23, 3, 23, 3, 23, 2, 2, 24, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 2, 5, 3, 2, 22, 23, 3, 2, 25, 26, 3, 2, 30, 31, 2, 194, 2, 49, 3, 2,
	2, 2, 4, 54, 3, 2, 2, 2, 6, 69, 3, 2, 2, 2, 8, 71, 3, 2, 2, 2, 10, 78,
	3, 2, 2, 2, 12, 86, 3, 2, 2, 2, 14, 91, 3, 2, 2, 2, 16, 93, 3, 2, 2, 2,
	18, 98, 3, 2, 2, 2, 20, 112, 3, 2, 2, 2, 22, 120, 3, 2, 2, 2, 24, 131,
	3, 2, 2, 2, 26, 139, 3, 2, 2, 2, 28, 147, 3, 2, 2, 2, 30, 155, 3, 2, 2,
	2, 32, 159, 3, 2, 2, 2, 34, 166, 3, 2, 2, 2, 36, 172, 3, 2, 2, 2, 38, 181,
	3, 2, 2, 2, 40, 186, 3, 2, 2, 2, 42, 192, 3, 2, 2, 2, 44, 194, 3, 2, 2,
	2, 46, 48, 5, 4, 3, 2, 47, 46, 3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47,
	3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 52, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2,
	52, 53, 7, 2, 2, 3, 53, 3, 3, 2, 2, 2, 54, 59, 5, 6, 4, 2, 55, 56, 7, 28,
	2, 2, 56, 58, 5, 6, 4, 2, 57, 55, 3, 2, 2, 2, 58, 61, 3, 2, 2, 2, 59, 57,
	3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 5, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2,
	62, 70, 5, 8, 5, 2, 63, 70, 5, 18, 10, 2, 64, 70, 5, 22, 12, 2, 65, 70,
	5, 26, 14, 2, 66, 70, 5, 32, 17, 2, 67, 70, 5, 34, 18, 2, 68, 70, 5, 36,
	19, 2, 69, 62, 3, 2, 2, 2, 69, 63, 3, 2, 2, 2, 69, 64, 3, 2, 2, 2, 69,
	65, 3, 2, 2, 2, 69, 66, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 68, 3, 2, 2,
	2, 70, 7, 3, 2, 2, 2, 71, 72, 7, 5, 2, 2, 72, 73, 7, 15, 2, 2, 73, 74,
	7, 29, 2, 2, 74, 75, 7, 3, 2, 2, 75, 76, 5, 10, 6, 2, 76, 77, 7, 4, 2,
	2, 77, 9, 3, 2, 2, 2, 78, 83, 5, 12, 7, 2, 79, 80, 7, 27, 2, 2, 80, 82,
	5, 12, 7, 2, 81, 79, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2,
	83, 84, 3, 2, 2, 2, 84, 11, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 87, 7,
	29, 2, 2, 87, 88, 5, 14, 8, 2, 88, 13, 3, 2, 2, 2, 89, 92, 7, 20, 2, 2,
	90, 92, 5, 16, 9, 2, 91, 89, 3, 2, 2, 2, 91, 90, 3, 2, 2, 2, 92, 15, 3,
	2, 2, 2, 93, 94, 7, 21, 2, 2, 94, 95, 7, 3, 2, 2, 95, 96, 7, 30, 2, 2,
	96, 97, 7, 4, 2, 2, 97, 17, 3, 2, 2, 2, 98, 99, 7, 6, 2, 2, 99, 100, 7,
	13, 2, 2, 100, 105, 7, 29, 2, 2, 101, 102, 7, 3, 2, 2, 102, 103, 5, 24,
	13, 2, 103, 104, 7, 4, 2, 2, 104, 106, 3, 2, 2, 2, 105, 101, 3, 2, 2, 2,
	105, 106, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 108, 7, 14, 2, 2, 108,
	109, 7, 3, 2, 2, 109, 110, 5, 20, 11, 2, 110, 111, 7, 4, 2, 2, 111, 19,
	3, 2, 2, 2, 112, 117, 5, 44, 23, 2, 113, 114, 7, 27, 2, 2, 114, 116, 5,
	44, 23, 2, 115, 113, 3, 2, 2, 2, 116, 119, 3, 2, 2, 2, 117, 115, 3, 2,
	2, 2, 117, 118, 3, 2, 2, 2, 118, 21, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2,
	120, 123, 7, 7, 2, 2, 121, 124, 7, 24, 2, 2, 122, 124, 5, 24, 13, 2, 123,
	121, 3, 2, 2, 2, 123, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 126,
	7, 10, 2, 2, 126, 129, 5, 24, 13, 2, 127, 128, 7, 12, 2, 2, 128, 130, 5,
	38, 20, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 23, 3, 2, 2,
	2, 131, 136, 7, 29, 2, 2, 132, 133, 7, 27, 2, 2, 133, 135, 7, 29, 2, 2,
	134, 132, 3, 2, 2, 2, 135, 138, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136,
	137, 3, 2, 2, 2, 137, 25, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 139, 140, 7,
	8, 2, 2, 140, 141, 7, 29, 2, 2, 141, 142, 7, 11, 2, 2, 142, 145, 5, 28,
	15, 2, 143, 144, 7, 12, 2, 2, 144, 146, 5, 38, 20, 2, 145, 143, 3, 2, 2,
	2, 145, 146, 3, 2, 2, 2, 146, 27, 3, 2, 2, 2, 147, 152, 5, 30, 16, 2, 148,
	149, 7, 27, 2, 2, 149, 151, 5, 30, 16, 2, 150, 148, 3, 2, 2, 2, 151, 154,
	3, 2, 2, 2, 152, 150, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 29, 3, 2,
	2, 2, 154, 152, 3, 2, 2, 2, 155, 156, 7, 29, 2, 2, 156, 157, 7, 25, 2,
	2, 157, 158, 5, 42, 22, 2, 158, 31, 3, 2, 2, 2, 159, 160, 7, 9, 2, 2, 160,
	161, 7, 10, 2, 2, 161, 164, 7, 29, 2, 2, 162, 163, 7, 12, 2, 2, 163, 165,
	5, 38, 20, 2, 164, 162, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 33, 3, 2,
	2, 2, 166, 167, 7, 5, 2, 2, 167, 168, 7, 17, 2, 2, 168, 169, 7, 29, 2,
	2, 169, 170, 7, 18, 2, 2, 170, 171, 5, 22, 12, 2, 171, 35, 3, 2, 2, 2,
	172, 173, 7, 5, 2, 2, 173, 174, 7, 16, 2, 2, 174, 175, 7, 29, 2, 2, 175,
	176, 7, 19, 2, 2, 176, 177, 7, 29, 2, 2, 177, 178, 7, 3, 2, 2, 178, 179,
	7, 29, 2, 2, 179, 180, 7, 4, 2, 2, 180, 37, 3, 2, 2, 2, 181, 184, 5, 40,
	21, 2, 182, 183, 9, 2, 2, 2, 183, 185, 5, 40, 21, 2, 184, 182, 3, 2, 2,
	2, 184, 185, 3, 2, 2, 2, 185, 39, 3, 2, 2, 2, 186, 187, 5, 42, 22, 2, 187,
	188, 9, 3, 2, 2, 188, 189, 5, 42, 22, 2, 189, 41, 3, 2, 2, 2, 190, 193,
	7, 29, 2, 2, 191, 193, 5, 44, 23, 2, 192, 190, 3, 2, 2, 2, 192, 191, 3,
	2, 2, 2, 193, 43, 3, 2, 2, 2, 194, 195, 9, 4, 2, 2, 195, 45, 3, 2, 2, 2,
	17, 49, 59, 69, 83, 91, 105, 117, 123, 129, 136, 145, 152, 164, 184, 192,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'create'", "'insert'", "'select'", "'update'", "'delete'",
	"'from'", "'set'", "'where'", "'into'", "'values'", "'table'", "'index'",
	"'view'", "'as'", "'on'", "'int'", "'varchar'", "'and'", "'or'", "'*'",
	"'='", "'!='", "','", "';'",
}
var symbolicNames = []string{
	"", "", "", "CREATE_", "INSERT_", "SELECT_", "UPDATE_", "DELETE_", "FROM_",
	"SET_", "WHERE_", "INTO_", "VALUES_", "TABLE_", "INDEX_", "VIEW_", "AS_",
	"ON_", "INT_", "VAR_CHAR_", "AND_", "OR_", "STAR", "EQUAL", "NOT_EQUAL",
	"COMMA", "SEMI_COLON", "IDENT", "INT_LITERAL", "STR_LITERAL", "SPACES",
}

var ruleNames = []string{
	"parse", "statementList", "statement", "create_table_stmt", "field_specs",
	"field_spec", "type_spec", "varchar_spec", "insert_stmt", "constant_list",
	"select_stmt", "ident_list", "update_stmt", "update_expr_list", "update_expr",
	"delete_stmt", "create_view_stmt", "create_index_stmt", "condition", "term",
	"expression", "literal",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SimpleSqlParser struct {
	*antlr.BaseParser
}

func NewSimpleSqlParser(input antlr.TokenStream) *SimpleSqlParser {
	this := new(SimpleSqlParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "SimpleSql.g4"

	return this
}

// SimpleSqlParser tokens.
const (
	SimpleSqlParserEOF         = antlr.TokenEOF
	SimpleSqlParserT__0        = 1
	SimpleSqlParserT__1        = 2
	SimpleSqlParserCREATE_     = 3
	SimpleSqlParserINSERT_     = 4
	SimpleSqlParserSELECT_     = 5
	SimpleSqlParserUPDATE_     = 6
	SimpleSqlParserDELETE_     = 7
	SimpleSqlParserFROM_       = 8
	SimpleSqlParserSET_        = 9
	SimpleSqlParserWHERE_      = 10
	SimpleSqlParserINTO_       = 11
	SimpleSqlParserVALUES_     = 12
	SimpleSqlParserTABLE_      = 13
	SimpleSqlParserINDEX_      = 14
	SimpleSqlParserVIEW_       = 15
	SimpleSqlParserAS_         = 16
	SimpleSqlParserON_         = 17
	SimpleSqlParserINT_        = 18
	SimpleSqlParserVAR_CHAR_   = 19
	SimpleSqlParserAND_        = 20
	SimpleSqlParserOR_         = 21
	SimpleSqlParserSTAR        = 22
	SimpleSqlParserEQUAL       = 23
	SimpleSqlParserNOT_EQUAL   = 24
	SimpleSqlParserCOMMA       = 25
	SimpleSqlParserSEMI_COLON  = 26
	SimpleSqlParserIDENT       = 27
	SimpleSqlParserINT_LITERAL = 28
	SimpleSqlParserSTR_LITERAL = 29
	SimpleSqlParserSPACES      = 30
)

// SimpleSqlParser rules.
const (
	SimpleSqlParserRULE_parse             = 0
	SimpleSqlParserRULE_statementList     = 1
	SimpleSqlParserRULE_statement         = 2
	SimpleSqlParserRULE_create_table_stmt = 3
	SimpleSqlParserRULE_field_specs       = 4
	SimpleSqlParserRULE_field_spec        = 5
	SimpleSqlParserRULE_type_spec         = 6
	SimpleSqlParserRULE_varchar_spec      = 7
	SimpleSqlParserRULE_insert_stmt       = 8
	SimpleSqlParserRULE_constant_list     = 9
	SimpleSqlParserRULE_select_stmt       = 10
	SimpleSqlParserRULE_ident_list        = 11
	SimpleSqlParserRULE_update_stmt       = 12
	SimpleSqlParserRULE_update_expr_list  = 13
	SimpleSqlParserRULE_update_expr       = 14
	SimpleSqlParserRULE_delete_stmt       = 15
	SimpleSqlParserRULE_create_view_stmt  = 16
	SimpleSqlParserRULE_create_index_stmt = 17
	SimpleSqlParserRULE_condition         = 18
	SimpleSqlParserRULE_term              = 19
	SimpleSqlParserRULE_expression        = 20
	SimpleSqlParserRULE_literal           = 21
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserEOF, 0)
}

func (s *ParseContext) AllStatementList() []IStatementListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementListContext)(nil)).Elem())
	var tst = make([]IStatementListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementListContext)
		}
	}

	return tst
}

func (s *ParseContext) StatementList(i int) IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SimpleSqlParserRULE_parse)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimpleSqlParserCREATE_)|(1<<SimpleSqlParserINSERT_)|(1<<SimpleSqlParserSELECT_)|(1<<SimpleSqlParserUPDATE_)|(1<<SimpleSqlParserDELETE_))) != 0 {
		{
			p.SetState(44)
			p.StatementList()
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(50)
		p.Match(SimpleSqlParserEOF)
	}

	return localctx
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_statementList
	return p
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementListContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListContext) AllSEMI_COLON() []antlr.TerminalNode {
	return s.GetTokens(SimpleSqlParserSEMI_COLON)
}

func (s *StatementListContext) SEMI_COLON(i int) antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserSEMI_COLON, i)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitStatementList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SimpleSqlParserRULE_statementList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Statement()
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SimpleSqlParserSEMI_COLON {
		{
			p.SetState(53)
			p.Match(SimpleSqlParserSEMI_COLON)
		}
		{
			p.SetState(54)
			p.Statement()
		}

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Create_table_stmt() ICreate_table_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreate_table_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreate_table_stmtContext)
}

func (s *StatementContext) Insert_stmt() IInsert_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInsert_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInsert_stmtContext)
}

func (s *StatementContext) Select_stmt() ISelect_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelect_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelect_stmtContext)
}

func (s *StatementContext) Update_stmt() IUpdate_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdate_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdate_stmtContext)
}

func (s *StatementContext) Delete_stmt() IDelete_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelete_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDelete_stmtContext)
}

func (s *StatementContext) Create_view_stmt() ICreate_view_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreate_view_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreate_view_stmtContext)
}

func (s *StatementContext) Create_index_stmt() ICreate_index_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreate_index_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreate_index_stmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SimpleSqlParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			p.Create_table_stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.Insert_stmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(62)
			p.Select_stmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(63)
			p.Update_stmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(64)
			p.Delete_stmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(65)
			p.Create_view_stmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(66)
			p.Create_index_stmt()
		}

	}

	return localctx
}

// ICreate_table_stmtContext is an interface to support dynamic dispatch.
type ICreate_table_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreate_table_stmtContext differentiates from other interfaces.
	IsCreate_table_stmtContext()
}

type Create_table_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreate_table_stmtContext() *Create_table_stmtContext {
	var p = new(Create_table_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_create_table_stmt
	return p
}

func (*Create_table_stmtContext) IsCreate_table_stmtContext() {}

func NewCreate_table_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Create_table_stmtContext {
	var p = new(Create_table_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_create_table_stmt

	return p
}

func (s *Create_table_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Create_table_stmtContext) CREATE_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserCREATE_, 0)
}

func (s *Create_table_stmtContext) TABLE_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserTABLE_, 0)
}

func (s *Create_table_stmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserIDENT, 0)
}

func (s *Create_table_stmtContext) Field_specs() IField_specsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_specsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_specsContext)
}

func (s *Create_table_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Create_table_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Create_table_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitCreate_table_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Create_table_stmt() (localctx ICreate_table_stmtContext) {
	localctx = NewCreate_table_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SimpleSqlParserRULE_create_table_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(SimpleSqlParserCREATE_)
	}
	{
		p.SetState(70)
		p.Match(SimpleSqlParserTABLE_)
	}
	{
		p.SetState(71)
		p.Match(SimpleSqlParserIDENT)
	}
	{
		p.SetState(72)
		p.Match(SimpleSqlParserT__0)
	}
	{
		p.SetState(73)
		p.Field_specs()
	}
	{
		p.SetState(74)
		p.Match(SimpleSqlParserT__1)
	}

	return localctx
}

// IField_specsContext is an interface to support dynamic dispatch.
type IField_specsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_specsContext differentiates from other interfaces.
	IsField_specsContext()
}

type Field_specsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_specsContext() *Field_specsContext {
	var p = new(Field_specsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_field_specs
	return p
}

func (*Field_specsContext) IsField_specsContext() {}

func NewField_specsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_specsContext {
	var p = new(Field_specsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_field_specs

	return p
}

func (s *Field_specsContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_specsContext) AllField_spec() []IField_specContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IField_specContext)(nil)).Elem())
	var tst = make([]IField_specContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IField_specContext)
		}
	}

	return tst
}

func (s *Field_specsContext) Field_spec(i int) IField_specContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_specContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IField_specContext)
}

func (s *Field_specsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SimpleSqlParserCOMMA)
}

func (s *Field_specsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserCOMMA, i)
}

func (s *Field_specsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_specsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_specsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitField_specs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Field_specs() (localctx IField_specsContext) {
	localctx = NewField_specsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SimpleSqlParserRULE_field_specs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Field_spec()
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SimpleSqlParserCOMMA {
		{
			p.SetState(77)
			p.Match(SimpleSqlParserCOMMA)
		}
		{
			p.SetState(78)
			p.Field_spec()
		}

		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IField_specContext is an interface to support dynamic dispatch.
type IField_specContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_specContext differentiates from other interfaces.
	IsField_specContext()
}

type Field_specContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_specContext() *Field_specContext {
	var p = new(Field_specContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_field_spec
	return p
}

func (*Field_specContext) IsField_specContext() {}

func NewField_specContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_specContext {
	var p = new(Field_specContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_field_spec

	return p
}

func (s *Field_specContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_specContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserIDENT, 0)
}

func (s *Field_specContext) Type_spec() IType_specContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_specContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_specContext)
}

func (s *Field_specContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_specContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_specContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitField_spec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Field_spec() (localctx IField_specContext) {
	localctx = NewField_specContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SimpleSqlParserRULE_field_spec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(SimpleSqlParserIDENT)
	}
	{
		p.SetState(85)
		p.Type_spec()
	}

	return localctx
}

// IType_specContext is an interface to support dynamic dispatch.
type IType_specContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_specContext differentiates from other interfaces.
	IsType_specContext()
}

type Type_specContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_specContext() *Type_specContext {
	var p = new(Type_specContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_type_spec
	return p
}

func (*Type_specContext) IsType_specContext() {}

func NewType_specContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_specContext {
	var p = new(Type_specContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_type_spec

	return p
}

func (s *Type_specContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_specContext) INT_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserINT_, 0)
}

func (s *Type_specContext) Varchar_spec() IVarchar_specContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarchar_specContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarchar_specContext)
}

func (s *Type_specContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_specContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_specContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitType_spec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Type_spec() (localctx IType_specContext) {
	localctx = NewType_specContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SimpleSqlParserRULE_type_spec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(89)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimpleSqlParserINT_:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Match(SimpleSqlParserINT_)
		}

	case SimpleSqlParserVAR_CHAR_:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(88)
			p.Varchar_spec()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVarchar_specContext is an interface to support dynamic dispatch.
type IVarchar_specContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarchar_specContext differentiates from other interfaces.
	IsVarchar_specContext()
}

type Varchar_specContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarchar_specContext() *Varchar_specContext {
	var p = new(Varchar_specContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_varchar_spec
	return p
}

func (*Varchar_specContext) IsVarchar_specContext() {}

func NewVarchar_specContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Varchar_specContext {
	var p = new(Varchar_specContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_varchar_spec

	return p
}

func (s *Varchar_specContext) GetParser() antlr.Parser { return s.parser }

func (s *Varchar_specContext) VAR_CHAR_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserVAR_CHAR_, 0)
}

func (s *Varchar_specContext) INT_LITERAL() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserINT_LITERAL, 0)
}

func (s *Varchar_specContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Varchar_specContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Varchar_specContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitVarchar_spec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Varchar_spec() (localctx IVarchar_specContext) {
	localctx = NewVarchar_specContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SimpleSqlParserRULE_varchar_spec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(SimpleSqlParserVAR_CHAR_)
	}
	{
		p.SetState(92)
		p.Match(SimpleSqlParserT__0)
	}
	{
		p.SetState(93)
		p.Match(SimpleSqlParserINT_LITERAL)
	}
	{
		p.SetState(94)
		p.Match(SimpleSqlParserT__1)
	}

	return localctx
}

// IInsert_stmtContext is an interface to support dynamic dispatch.
type IInsert_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInsert_stmtContext differentiates from other interfaces.
	IsInsert_stmtContext()
}

type Insert_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsert_stmtContext() *Insert_stmtContext {
	var p = new(Insert_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_insert_stmt
	return p
}

func (*Insert_stmtContext) IsInsert_stmtContext() {}

func NewInsert_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Insert_stmtContext {
	var p = new(Insert_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_insert_stmt

	return p
}

func (s *Insert_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Insert_stmtContext) INSERT_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserINSERT_, 0)
}

func (s *Insert_stmtContext) INTO_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserINTO_, 0)
}

func (s *Insert_stmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserIDENT, 0)
}

func (s *Insert_stmtContext) VALUES_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserVALUES_, 0)
}

func (s *Insert_stmtContext) Constant_list() IConstant_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstant_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstant_listContext)
}

func (s *Insert_stmtContext) Ident_list() IIdent_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdent_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdent_listContext)
}

func (s *Insert_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Insert_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Insert_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitInsert_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Insert_stmt() (localctx IInsert_stmtContext) {
	localctx = NewInsert_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SimpleSqlParserRULE_insert_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(SimpleSqlParserINSERT_)
	}
	{
		p.SetState(97)
		p.Match(SimpleSqlParserINTO_)
	}
	{
		p.SetState(98)
		p.Match(SimpleSqlParserIDENT)
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleSqlParserT__0 {
		{
			p.SetState(99)
			p.Match(SimpleSqlParserT__0)
		}
		{
			p.SetState(100)
			p.Ident_list()
		}
		{
			p.SetState(101)
			p.Match(SimpleSqlParserT__1)
		}

	}
	{
		p.SetState(105)
		p.Match(SimpleSqlParserVALUES_)
	}
	{
		p.SetState(106)
		p.Match(SimpleSqlParserT__0)
	}
	{
		p.SetState(107)
		p.Constant_list()
	}
	{
		p.SetState(108)
		p.Match(SimpleSqlParserT__1)
	}

	return localctx
}

// IConstant_listContext is an interface to support dynamic dispatch.
type IConstant_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstant_listContext differentiates from other interfaces.
	IsConstant_listContext()
}

type Constant_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstant_listContext() *Constant_listContext {
	var p = new(Constant_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_constant_list
	return p
}

func (*Constant_listContext) IsConstant_listContext() {}

func NewConstant_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Constant_listContext {
	var p = new(Constant_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_constant_list

	return p
}

func (s *Constant_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Constant_listContext) AllLiteral() []ILiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILiteralContext)(nil)).Elem())
	var tst = make([]ILiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILiteralContext)
		}
	}

	return tst
}

func (s *Constant_listContext) Literal(i int) ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *Constant_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SimpleSqlParserCOMMA)
}

func (s *Constant_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserCOMMA, i)
}

func (s *Constant_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Constant_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Constant_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitConstant_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Constant_list() (localctx IConstant_listContext) {
	localctx = NewConstant_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SimpleSqlParserRULE_constant_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Literal()
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SimpleSqlParserCOMMA {
		{
			p.SetState(111)
			p.Match(SimpleSqlParserCOMMA)
		}
		{
			p.SetState(112)
			p.Literal()
		}

		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISelect_stmtContext is an interface to support dynamic dispatch.
type ISelect_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelect_stmtContext differentiates from other interfaces.
	IsSelect_stmtContext()
}

type Select_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelect_stmtContext() *Select_stmtContext {
	var p = new(Select_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_select_stmt
	return p
}

func (*Select_stmtContext) IsSelect_stmtContext() {}

func NewSelect_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select_stmtContext {
	var p = new(Select_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_select_stmt

	return p
}

func (s *Select_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Select_stmtContext) SELECT_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserSELECT_, 0)
}

func (s *Select_stmtContext) FROM_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserFROM_, 0)
}

func (s *Select_stmtContext) AllIdent_list() []IIdent_listContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdent_listContext)(nil)).Elem())
	var tst = make([]IIdent_listContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdent_listContext)
		}
	}

	return tst
}

func (s *Select_stmtContext) Ident_list(i int) IIdent_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdent_listContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdent_listContext)
}

func (s *Select_stmtContext) STAR() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserSTAR, 0)
}

func (s *Select_stmtContext) WHERE_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserWHERE_, 0)
}

func (s *Select_stmtContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *Select_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Select_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Select_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitSelect_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Select_stmt() (localctx ISelect_stmtContext) {
	localctx = NewSelect_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SimpleSqlParserRULE_select_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(SimpleSqlParserSELECT_)
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimpleSqlParserSTAR:
		{
			p.SetState(119)
			p.Match(SimpleSqlParserSTAR)
		}

	case SimpleSqlParserIDENT:
		{
			p.SetState(120)
			p.Ident_list()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(123)
		p.Match(SimpleSqlParserFROM_)
	}
	{
		p.SetState(124)
		p.Ident_list()
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleSqlParserWHERE_ {
		{
			p.SetState(125)
			p.Match(SimpleSqlParserWHERE_)
		}
		{
			p.SetState(126)
			p.Condition()
		}

	}

	return localctx
}

// IIdent_listContext is an interface to support dynamic dispatch.
type IIdent_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdent_listContext differentiates from other interfaces.
	IsIdent_listContext()
}

type Ident_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdent_listContext() *Ident_listContext {
	var p = new(Ident_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_ident_list
	return p
}

func (*Ident_listContext) IsIdent_listContext() {}

func NewIdent_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ident_listContext {
	var p = new(Ident_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_ident_list

	return p
}

func (s *Ident_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Ident_listContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(SimpleSqlParserIDENT)
}

func (s *Ident_listContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserIDENT, i)
}

func (s *Ident_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SimpleSqlParserCOMMA)
}

func (s *Ident_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserCOMMA, i)
}

func (s *Ident_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ident_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ident_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitIdent_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Ident_list() (localctx IIdent_listContext) {
	localctx = NewIdent_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SimpleSqlParserRULE_ident_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(SimpleSqlParserIDENT)
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SimpleSqlParserCOMMA {
		{
			p.SetState(130)
			p.Match(SimpleSqlParserCOMMA)
		}
		{
			p.SetState(131)
			p.Match(SimpleSqlParserIDENT)
		}

		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUpdate_stmtContext is an interface to support dynamic dispatch.
type IUpdate_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdate_stmtContext differentiates from other interfaces.
	IsUpdate_stmtContext()
}

type Update_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_stmtContext() *Update_stmtContext {
	var p = new(Update_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_update_stmt
	return p
}

func (*Update_stmtContext) IsUpdate_stmtContext() {}

func NewUpdate_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_stmtContext {
	var p = new(Update_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_update_stmt

	return p
}

func (s *Update_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Update_stmtContext) UPDATE_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserUPDATE_, 0)
}

func (s *Update_stmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserIDENT, 0)
}

func (s *Update_stmtContext) SET_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserSET_, 0)
}

func (s *Update_stmtContext) Update_expr_list() IUpdate_expr_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdate_expr_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdate_expr_listContext)
}

func (s *Update_stmtContext) WHERE_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserWHERE_, 0)
}

func (s *Update_stmtContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *Update_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Update_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitUpdate_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Update_stmt() (localctx IUpdate_stmtContext) {
	localctx = NewUpdate_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SimpleSqlParserRULE_update_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.Match(SimpleSqlParserUPDATE_)
	}
	{
		p.SetState(138)
		p.Match(SimpleSqlParserIDENT)
	}
	{
		p.SetState(139)
		p.Match(SimpleSqlParserSET_)
	}
	{
		p.SetState(140)
		p.Update_expr_list()
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleSqlParserWHERE_ {
		{
			p.SetState(141)
			p.Match(SimpleSqlParserWHERE_)
		}
		{
			p.SetState(142)
			p.Condition()
		}

	}

	return localctx
}

// IUpdate_expr_listContext is an interface to support dynamic dispatch.
type IUpdate_expr_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdate_expr_listContext differentiates from other interfaces.
	IsUpdate_expr_listContext()
}

type Update_expr_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_expr_listContext() *Update_expr_listContext {
	var p = new(Update_expr_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_update_expr_list
	return p
}

func (*Update_expr_listContext) IsUpdate_expr_listContext() {}

func NewUpdate_expr_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_expr_listContext {
	var p = new(Update_expr_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_update_expr_list

	return p
}

func (s *Update_expr_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Update_expr_listContext) AllUpdate_expr() []IUpdate_exprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUpdate_exprContext)(nil)).Elem())
	var tst = make([]IUpdate_exprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUpdate_exprContext)
		}
	}

	return tst
}

func (s *Update_expr_listContext) Update_expr(i int) IUpdate_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdate_exprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUpdate_exprContext)
}

func (s *Update_expr_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SimpleSqlParserCOMMA)
}

func (s *Update_expr_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserCOMMA, i)
}

func (s *Update_expr_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_expr_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Update_expr_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitUpdate_expr_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Update_expr_list() (localctx IUpdate_expr_listContext) {
	localctx = NewUpdate_expr_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SimpleSqlParserRULE_update_expr_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Update_expr()
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SimpleSqlParserCOMMA {
		{
			p.SetState(146)
			p.Match(SimpleSqlParserCOMMA)
		}
		{
			p.SetState(147)
			p.Update_expr()
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUpdate_exprContext is an interface to support dynamic dispatch.
type IUpdate_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdate_exprContext differentiates from other interfaces.
	IsUpdate_exprContext()
}

type Update_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_exprContext() *Update_exprContext {
	var p = new(Update_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_update_expr
	return p
}

func (*Update_exprContext) IsUpdate_exprContext() {}

func NewUpdate_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_exprContext {
	var p = new(Update_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_update_expr

	return p
}

func (s *Update_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Update_exprContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserIDENT, 0)
}

func (s *Update_exprContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserEQUAL, 0)
}

func (s *Update_exprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Update_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Update_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitUpdate_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Update_expr() (localctx IUpdate_exprContext) {
	localctx = NewUpdate_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SimpleSqlParserRULE_update_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(SimpleSqlParserIDENT)
	}
	{
		p.SetState(154)
		p.Match(SimpleSqlParserEQUAL)
	}
	{
		p.SetState(155)
		p.Expression()
	}

	return localctx
}

// IDelete_stmtContext is an interface to support dynamic dispatch.
type IDelete_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelete_stmtContext differentiates from other interfaces.
	IsDelete_stmtContext()
}

type Delete_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelete_stmtContext() *Delete_stmtContext {
	var p = new(Delete_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_delete_stmt
	return p
}

func (*Delete_stmtContext) IsDelete_stmtContext() {}

func NewDelete_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Delete_stmtContext {
	var p = new(Delete_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_delete_stmt

	return p
}

func (s *Delete_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Delete_stmtContext) DELETE_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserDELETE_, 0)
}

func (s *Delete_stmtContext) FROM_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserFROM_, 0)
}

func (s *Delete_stmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserIDENT, 0)
}

func (s *Delete_stmtContext) WHERE_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserWHERE_, 0)
}

func (s *Delete_stmtContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *Delete_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Delete_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Delete_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitDelete_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Delete_stmt() (localctx IDelete_stmtContext) {
	localctx = NewDelete_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SimpleSqlParserRULE_delete_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(SimpleSqlParserDELETE_)
	}
	{
		p.SetState(158)
		p.Match(SimpleSqlParserFROM_)
	}
	{
		p.SetState(159)
		p.Match(SimpleSqlParserIDENT)
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleSqlParserWHERE_ {
		{
			p.SetState(160)
			p.Match(SimpleSqlParserWHERE_)
		}
		{
			p.SetState(161)
			p.Condition()
		}

	}

	return localctx
}

// ICreate_view_stmtContext is an interface to support dynamic dispatch.
type ICreate_view_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreate_view_stmtContext differentiates from other interfaces.
	IsCreate_view_stmtContext()
}

type Create_view_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreate_view_stmtContext() *Create_view_stmtContext {
	var p = new(Create_view_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_create_view_stmt
	return p
}

func (*Create_view_stmtContext) IsCreate_view_stmtContext() {}

func NewCreate_view_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Create_view_stmtContext {
	var p = new(Create_view_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_create_view_stmt

	return p
}

func (s *Create_view_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Create_view_stmtContext) CREATE_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserCREATE_, 0)
}

func (s *Create_view_stmtContext) VIEW_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserVIEW_, 0)
}

func (s *Create_view_stmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserIDENT, 0)
}

func (s *Create_view_stmtContext) AS_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserAS_, 0)
}

func (s *Create_view_stmtContext) Select_stmt() ISelect_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelect_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelect_stmtContext)
}

func (s *Create_view_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Create_view_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Create_view_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitCreate_view_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Create_view_stmt() (localctx ICreate_view_stmtContext) {
	localctx = NewCreate_view_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SimpleSqlParserRULE_create_view_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Match(SimpleSqlParserCREATE_)
	}
	{
		p.SetState(165)
		p.Match(SimpleSqlParserVIEW_)
	}
	{
		p.SetState(166)
		p.Match(SimpleSqlParserIDENT)
	}
	{
		p.SetState(167)
		p.Match(SimpleSqlParserAS_)
	}
	{
		p.SetState(168)
		p.Select_stmt()
	}

	return localctx
}

// ICreate_index_stmtContext is an interface to support dynamic dispatch.
type ICreate_index_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreate_index_stmtContext differentiates from other interfaces.
	IsCreate_index_stmtContext()
}

type Create_index_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreate_index_stmtContext() *Create_index_stmtContext {
	var p = new(Create_index_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_create_index_stmt
	return p
}

func (*Create_index_stmtContext) IsCreate_index_stmtContext() {}

func NewCreate_index_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Create_index_stmtContext {
	var p = new(Create_index_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_create_index_stmt

	return p
}

func (s *Create_index_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Create_index_stmtContext) CREATE_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserCREATE_, 0)
}

func (s *Create_index_stmtContext) INDEX_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserINDEX_, 0)
}

func (s *Create_index_stmtContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(SimpleSqlParserIDENT)
}

func (s *Create_index_stmtContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserIDENT, i)
}

func (s *Create_index_stmtContext) ON_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserON_, 0)
}

func (s *Create_index_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Create_index_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Create_index_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitCreate_index_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Create_index_stmt() (localctx ICreate_index_stmtContext) {
	localctx = NewCreate_index_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SimpleSqlParserRULE_create_index_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(SimpleSqlParserCREATE_)
	}
	{
		p.SetState(171)
		p.Match(SimpleSqlParserINDEX_)
	}
	{
		p.SetState(172)
		p.Match(SimpleSqlParserIDENT)
	}
	{
		p.SetState(173)
		p.Match(SimpleSqlParserON_)
	}
	{
		p.SetState(174)
		p.Match(SimpleSqlParserIDENT)
	}
	{
		p.SetState(175)
		p.Match(SimpleSqlParserT__0)
	}
	{
		p.SetState(176)
		p.Match(SimpleSqlParserIDENT)
	}
	{
		p.SetState(177)
		p.Match(SimpleSqlParserT__1)
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) GetOp() antlr.Token { return s.op }

func (s *ConditionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ConditionContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *ConditionContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ConditionContext) AND_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserAND_, 0)
}

func (s *ConditionContext) OR_() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserOR_, 0)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SimpleSqlParserRULE_condition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Term()
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleSqlParserAND_ || _la == SimpleSqlParserOR_ {
		{
			p.SetState(180)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ConditionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SimpleSqlParserAND_ || _la == SimpleSqlParserOR_) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ConditionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(181)
			p.Term()
		}

	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOperator returns the operator token.
	GetOperator() antlr.Token

	// SetOperator sets the operator token.
	SetOperator(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExpressionContext

	// GetRight returns the right rule contexts.
	GetRight() IExpressionContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpressionContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpressionContext)

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	left     IExpressionContext
	operator antlr.Token
	right    IExpressionContext
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) GetOperator() antlr.Token { return s.operator }

func (s *TermContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *TermContext) GetLeft() IExpressionContext { return s.left }

func (s *TermContext) GetRight() IExpressionContext { return s.right }

func (s *TermContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *TermContext) SetRight(v IExpressionContext) { s.right = v }

func (s *TermContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *TermContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TermContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserEQUAL, 0)
}

func (s *TermContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserNOT_EQUAL, 0)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SimpleSqlParserRULE_term)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)

		var _x = p.Expression()

		localctx.(*TermContext).left = _x
	}
	{
		p.SetState(185)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*TermContext).operator = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == SimpleSqlParserEQUAL || _la == SimpleSqlParserNOT_EQUAL) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*TermContext).operator = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(186)

		var _x = p.Expression()

		localctx.(*TermContext).right = _x
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserIDENT, 0)
}

func (s *ExpressionContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SimpleSqlParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(190)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimpleSqlParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.Match(SimpleSqlParserIDENT)
		}

	case SimpleSqlParserINT_LITERAL, SimpleSqlParserSTR_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(189)
			p.Literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) INT_LITERAL() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserINT_LITERAL, 0)
}

func (s *LiteralContext) STR_LITERAL() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserSTR_LITERAL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SimpleSqlParserRULE_literal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SimpleSqlParserINT_LITERAL || _la == SimpleSqlParserSTR_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
