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
%token <i> IMPORT PACKAGE TYPE

%token <i> FUNC INTERFACE ELLIPSIS STRUCT
%token <i> MAP

%token <i> SHL SHR AND_NOT
%token <i> ADD_ASSIGN SUB_ASSIGN MUL_ASSIGN QUO_ASSIGN REM_ASSIGN
%token <i> AND_ASSIGN OR_ASSIGN XOR_ASSIGN SHL_ASSIGN SHR_ASSIGN AND_NOT_ASSIGN
%token <i> LAND LOR INC DEC
%token <i> DEFINE
%token <i> BREAK CASE CHAN CONST CONTINUE
%token <i> DEFAULT DEFER ELSE FALLTHROUGH FOR
%token <i> GO GOTO IF
%token <i> RANGE RETURN
%token <i> SELECT SWITCH VAR


%%

file: file root {} | root {} ;

root: PACKAGE IDENT ';' {}
| importDecl ';' {}
| fnDecl ';' {}
| declaration ';' {}
;

declaration: constDecl {} | typeDecl {} | varDecl {} ;

fnDecl: FUNC IDENT fn {} | FUNC receiver IDENT fn {} ;

fn: signature block {} ;

receiver: parameters {} ;

importDecl: IMPORT importSpec {}
| IMPORT '(' ')' {} | IMPORT '(' importSpecList optSemi ')' {} ;

importSpecList: importSpecList ';' importSpec {} | importSpec {} ;

importSpec: STRING {} | '.' STRING {} | IDENT STRING {} ;

varDecl: VAR varSpec {}
| VAR '(' varSpecList optSemi ')' {} | VAR '(' ')' {} ;

varSpecList: varSpecList ';' varSpec {} | varSpec {} ;

varSpec: identList '=' exprList {}
| identList type {}
| identList type '=' exprList {} ;

constDecl: CONST constSpec {}
| CONST '(' constSpecList optSemi ')' {} | CONST '(' ')' {} ;

constSpecList: constSpecList ';' constSpec {} | constSpec {} ;

constSpec: identList '=' exprList {}
| identList type {}
| identList type '=' exprList {} ;

typeDecl: TYPE typeSpec {} | TYPE '(' typeSpecList optSemi ')' {} ;

typeSpecList: typeSpecList ';' typeSpec {} | typeSpec {} ;

typeSpec: IDENT type {} ;

type: typeName {} | typeLit {} | ARROW type {} ;

typeLit: arrayType {}
| sliceType {}
| mapType {}
| CHAN type {}
| structType {}
| interfaceType {}
| FUNC signature {}
| '*' type {}
;

methodSpecList: methodSpecList ';' methodSpec {} | methodSpec {} ;

methodSpec: typeName {} | IDENT signature {} ;

signature: parameters result {} | parameters {} ;

result: parameters {} | type {} ;

parameters: '(' paramList optComma  ')' {} | '(' ')' {} ;

paramList: paramList ',' paramDecl {} | paramDecl {} ;

paramDecl: IDENT type {} | type {} | IDENT ELLIPSIS type {} ;

fieldDeclList: fieldDeclList ';' fieldDecl {} | fieldDecl {} ;

fieldDecl: identList type {} | identList type STRING {}
| anonField {} | anonField STRING {} ;

anonField: '*' typeName {} | typeName {} ;

typeName: IDENT {} | qualifiedIdent {} ;

qualifiedIdent: IDENT '.' IDENT {} ;

identList: identList ',' IDENT {} | IDENT {} ;

// The optional semicolon is incorporated in stmtList
block: '{' '}' {} | '{' stmtList '}' {} ;

stmtList: stmtList ';' stmt {} | stmt {}
| stmtList ';' {} | ';' {} ;

stmt: declaration {} ;

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
| IDENT {}
| literal {} ;
;

literal: basicLit {} | fnLit {} | compositeLiteral {} ;

fnLit: FUNC fn {} ;

compositeLiteral: literalType literalValue {} ;

literalType: structType {}
| arrayType {}
| '[' ELLIPSIS ']' type {}
| sliceType {}
| mapType {}
| typeName {} ;

literalValue: '{' '}' {} | '{' elementList optComma '}' {} ;

elementList: elementList ',' keyedElement {} | keyedElement {} ;

keyedElement: key ':' element {} | element {} ;

key: expr {} | literalValue {} ;

element: expr {} | literalValue {} ;

structType: STRUCT '{' fieldDeclList optSemi '}' {} | STRUCT '{' '}' {} ;

arrayType: '[' expr ']' type {} ;

sliceType: '[' ']' type {} ;

mapType: MAP '[' type ']' type {} ;

interfaceType: INTERFACE '{' methodSpecList optSemi '}' {} | INTERFACE '{' '}' {} ;

basicLit: INT {} | FLOAT {} | IMAG {} | CHAR {} | STRING {} ;

optSemi: {} | ';' {} ;
optComma: {} | ',' {} ;

%%
