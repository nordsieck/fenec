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
%nonassoc UMINUS

%token <i> IDENT INT FLOAT IMAG CHAR STRING

%token <i> COMMENT
%token <i> SHL SHR AND_NOT
%token <i> ADD_ASSIGN SUB_ASSIGN MUL_ASSIGN QUO_ASSIGN REM_ASSIGN
%token <i> AND_ASSIGN OR_ASSIGN XOR_ASSIGN SHL_ASSIGN SHR_ASSIGN AND_NOT_ASSIGN
%token <i> LAND LOR ARROW INC DEC
%token <i> DEFINE ELLIPSIS
%token <i> BREAK CASE CHAN CONST CONTINUE
%token <i> DEFAULT DEFER ELSE FALLTHROUGH FOR
%token <i> FUNC GO GOTO IF IMPORT
%token <i> INTERFACE MAP PACKAGE RANGE RETURN
%token <i> SELECT STRUCT SWITCH TYPE VAR


%%

file: file expr {} | expr {} ;

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

%%
