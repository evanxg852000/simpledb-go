grammar SimpleSql;

/* antlr4 4.7.2 */

parse
    : statementList* EOF    
;

statementList
    : statement (SEMI_COLON statement)* 
;

statement
    : create_table_stmt
    | insert_stmt
    | select_stmt
    | update_stmt
    | delete_stmt
    | create_view_stmt
    | create_index_stmt
;

create_table_stmt: CREATE_ TABLE_ IDENT '(' field_specs ')' ;
field_specs: field_spec (COMMA field_spec)* ;
field_spec: IDENT type_spec ;
type_spec: INT_ | varchar_spec ;
varchar_spec: VAR_CHAR_ '(' INT_LITERAL ')' ;

insert_stmt: INSERT_ INTO_ IDENT ( '(' ident_list ')' )? VALUES_ '(' constant_list ')' ;
constant_list: literal (COMMA literal)* ;

select_stmt: SELECT_ (STAR | ident_list) FROM_ ident_list (WHERE_ condition)? ;
ident_list: IDENT (COMMA IDENT)* ;

update_stmt: UPDATE_ IDENT SET_ update_expr_list (WHERE_ condition)? ;
update_expr_list: update_expr (COMMA update_expr)* ;
update_expr: IDENT '=' expression ;

delete_stmt: DELETE_ FROM_ IDENT (WHERE_ condition)? ;

create_view_stmt: CREATE_ VIEW_ IDENT AS_ select_stmt ;

create_index_stmt: CREATE_ INDEX_ IDENT ON_ IDENT '(' IDENT ')' ;


condition: term ( op=(AND_ | OR_) term)?;
term: left=expression operator=(EQUAL|NOT_EQUAL) right=expression ;
expression: IDENT | literal ;
literal: INT_LITERAL | STR_LITERAL ;

/* keywords */

CREATE_: 'create' ;
INSERT_: 'insert' ;
SELECT_: 'select' ;
UPDATE_: 'update' ;
DELETE_: 'delete' ;
FROM_: 'from' ;
SET_: 'set' ;
WHERE_: 'where' ;
INTO_: 'into' ;
VALUES_: 'values' ;
TABLE_: 'table' ;
INDEX_: 'index' ;
VIEW_: 'view' ;
AS_: 'as' ;
ON_: 'on' ;
INT_: 'int';
VAR_CHAR_: 'varchar' ;
AND_: 'and' ;
OR_: 'or' ;

STAR: '*' ;
EQUAL: '=' ;
NOT_EQUAL: '!=' ;
COMMA: ',';
SEMI_COLON: ';';

IDENT: [a-zA-Z_][a-zA-Z0-9_]* ;
INT_LITERAL: '0'|[-+]?[1-9][0-9]* ;
STR_LITERAL: '\'' ( ~'\'' | '\'\'')* '\'' ;

SPACES: [ \t\r\n] -> skip ;

