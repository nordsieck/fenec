%{
  package parse
%}

%union{
  i int
}

%left EQL '<' '>' GEQ LEQ NEQ
%nonassoc '!'
%left '+' '-'
%left '*' '/' '%'
%left '.'
%nonassoc UMINUS ARROW

%token <i> IDENT INT FLOAT IMAG CHAR STRING
%token <i> IMPORT PACKAGE

%token <i> SHL SHR AND_NOT
%token <i> ADD_ASSIGN SUB_ASSIGN MUL_ASSIGN QUO_ASSIGN REM_ASSIGN
%token <i> AND_ASSIGN OR_ASSIGN XOR_ASSIGN SHL_ASSIGN SHR_ASSIGN AND_NOT_ASSIGN
%token <i> LAND LOR INC DEC
%token <i> DEFINE ELLIPSIS
%token <i> BREAK CASE CHAN CONST CONTINUE
%token <i> DEFAULT DEFER ELSE FALLTHROUGH FOR
%token <i> FUNC GO GOTO IF
%token <i> INTERFACE MAP RANGE RETURN
%token <i> SELECT STRUCT SWITCH TYPE VAR


%%

file: file root {} | root {} ;

root: PACKAGE IDENT ';' {}
| importDecl ';' {}
| constDecl ';' {}
| typeDecl ';' {}
;

importDecl: IMPORT importSpec {}
| IMPORT '(' ')' {} | IMPORT '(' importSpecList optSemi ')' {} ;

importSpecList: importSpecList ';' importSpec {} | importSpec {} ;

importSpec: STRING {} | '.' STRING {} | IDENT STRING {} ;

constDecl: CONST assignment {}
| CONST '(' assignmentList optSemi ')' {} | CONST '(' ')' {} ;

assignmentList: assignmentList ';' assignment {} | assignment {} ;

assignment: identList '=' exprList {} ;

typeDecl: TYPE typeSpec {} | TYPE '(' typeSpecList optSemi ')' {} ;

typeSpecList: typeSpecList ';' typeSpec {} | typeSpec {} ;

typeSpec: IDENT type {} ;

type: typeName {} | typeLit {} | ARROW type {} ;

typeLit: '[' expr ']' type {}
| '[' ']' type {}
| MAP '[' type ']' type {}
| CHAN type {}
| STRUCT '{' fieldDeclList optSemi '}' {} | STRUCT '{' '}' {}
| '*' type {}
;

fieldDeclList: fieldDeclList ';' fieldDecl {} | fieldDecl {} ;

fieldDecl: identList type {} | identList type STRING {}
| anonField {} | anonField STRING {} ;

anonField: '*' typeName {} | typeName {} ;

typeName: IDENT {} | qualifiedIdent {} ;

qualifiedIdent: IDENT '.' IDENT {} ;

identList: identList ',' IDENT {} | IDENT {} ;

exprList: exprList ',' expr {} | expr {} ;

expr: expr '+' expr {}
| expr '-' expr {}
| expr '*' expr {}
| expr '/' expr {}
| expr '%' expr {}
| expr EQL expr {}
| expr '<' expr {}
| expr '>' expr {}
| expr GEQ expr {}
| expr LEQ expr {}
| expr NEQ expr {}
| IDENT {} | INT {} | FLOAT {} | IMAG {} | CHAR {} | STRING {}
;

optSemi: {} | ';' {} ;

%%
