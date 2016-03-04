%{
  package parse
%}

%union{
  i int
}

%left EQL '<' '>' GEQ LEQ NEQ
%left LAND LOR
%left '+' '-'
%left '*' '/' '%'
%left SHL SHR AND_NOT '&' '|' '^'
%nonassoc '!'
%nonassoc UMINUS

%token <i> IDENT INT FLOAT IMAG CHAR STRING

%token <i> COMMENT
%token <i> INC DEC
%token <i> CONST VAR

%token <i> ADD_ASSIGN SUB_ASSIGN MUL_ASSIGN QUO_ASSIGN REM_ASSIGN
%token <i> AND_ASSIGN OR_ASSIGN XOR_ASSIGN SHL_ASSIGN SHR_ASSIGN AND_NOT_ASSIGN
%token <i> ARROW
%token <i> DEFINE ELLIPSIS
%token <i> BREAK CASE CHAN CONTINUE
%token <i> DEFAULT DEFER ELSE FALLTHROUGH FOR
%token <i> FUNC GO GOTO IF IMPORT
%token <i> INTERFACE MAP PACKAGE RANGE RETURN
%token <i> SELECT STRUCT SWITCH TYPE

%type <i> root val expr exprList assignment assignmentList
%type <i> stmt

%%

file: file root | root ;

root: PACKAGE IDENT ';'
| COMMENT
| IMPORT STRING ';'
| CONST assignment ';'
| CONST '(' assignmentList ')' ';'
| VAR assignment ';'
| VAR '(' assignmentList ')' ';'
| stmt // TODO: fix this
;

assignmentList: assignment ';' assignmentList {}
| assignment ';' {}
| assignment {}
;

assignment: identList '=' exprList {} ;

identList: identList ',' IDENT | IDENT

exprList: exprList ',' expr | expr

stmt: IDENT INC ';' | IDENT DEC ';'
;

expr: val
| expr '+' expr
| expr '-' expr
| expr '*' expr
| expr '/' expr
| expr '%' expr
| expr SHL expr
| expr SHR expr
| expr AND_NOT expr
| expr LAND expr
| expr LOR expr
| expr '&' expr
| expr '|' expr
| expr '^' expr
;

val: IDENT | INT | FLOAT | IMAG | CHAR | STRING
| '!' val {}
| '-' val %prec UMINUS {}
;




%%
