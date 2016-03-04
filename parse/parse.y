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
%nonassoc VALOF
%nonassoc ADDROF
%nonassoc '(' ')'

%token <i> IDENT INT FLOAT IMAG CHAR STRING
%token <i> COMMENT INC DEC CONST MAP PACKAGE

%token <i> VAR FUNC

%token <i> ADD_ASSIGN SUB_ASSIGN MUL_ASSIGN QUO_ASSIGN REM_ASSIGN
%token <i> AND_ASSIGN OR_ASSIGN XOR_ASSIGN SHL_ASSIGN SHR_ASSIGN AND_NOT_ASSIGN
%token <i> ARROW
%token <i> DEFINE ELLIPSIS
%token <i> BREAK CASE CHAN CONTINUE
%token <i> DEFAULT DEFER ELSE FALLTHROUGH FOR
%token <i> GO GOTO IF IMPORT
%token <i> INTERFACE RANGE RETURN
%token <i> SELECT STRUCT SWITCH TYPE

%type <i> root val expr exprList assignment assignmentList
%type <i> stmt type fnReturn fnParams optSemi
%type <i> paramList typeDecl paramDecl typeList fieldList
%type <i> ifaceFuncSig ifaceFuncList

%%

file: file root | root ;

root: PACKAGE IDENT ';' {}
| COMMENT {}
| IMPORT STRING ';' {}
| CONST assignment ';' {}
| CONST assignments ';' {}
| VAR assignment ';' {}
| VAR assignments ';' {}
| TYPE IDENT type ';' {}
| STRUCT IDENT structFields ';' {}
| stmt {} // TODO: fix this
;

optSemi: ';' {} | {} ;
optComma: ',' {} | {} ;

assignments: '(' assignmentList optSemi ')' {} | '(' ')' {} ;

assignmentList: assignmentList ';' assignment {} | assignment {} ;

assignment: identList '=' exprList {}
| identList type '=' exprList {}
| identList type {}
;

ifaceFuncSig: IDENT fnParams fnReturn {} | IDENT fnParams {} ;

ifaceFuncList: ifaceFuncList ';' ifaceFuncSig {} | ifaceFuncSig {} ;

ifaceMethods: '{' ifaceFuncList optSemi '}' {} | '{' '}' {} ;

identList: identList ',' IDENT {} | IDENT {} ;

exprList: exprList ',' expr {} | expr {} ;

stmt: IDENT INC ';' {} | IDENT DEC ';' {} ;

typeDecl: identList type {} ;

typeList: typeList ',' typeDecl {} | typeDecl {} ;

fieldList: fieldList ';' typeList {} | typeList

structFields: '{' fieldList optSemi '}' {} | '{' '}' {} ;

paramDecl: type {} | IDENT type {} | ELLIPSIS type {} | IDENT ELLIPSIS type {} ;

paramList: paramList ',' paramDecl {} | paramDecl {} ;

fnParams: '(' paramList optComma ')' {} | '(' ')' {} ;

fnReturn: type {} | fnParams {} ;

type: IDENT {}
| '[' ']' type {}
| '[' val ']' IDENT {}
| MAP '[' type ']' type {}
| FUNC fnParams fnReturn {}
| FUNC fnParams {}
| STRUCT structFields {}
| INTERFACE ifaceMethods {}
| '*' type {}
;

expr: val
| expr '+' expr {}
| expr '-' expr {}
| expr '*' expr {}
| expr '/' expr {}
| expr '%' expr {}
| expr SHL expr {}
| expr SHR expr {}
| expr AND_NOT expr {}
| expr LAND expr {}
| expr LOR expr {}
| expr '&' expr {}
| expr '|' expr {}
| expr '^' expr {}
| '(' expr ')' {}
;

val: IDENT {} | INT {} | FLOAT {} | IMAG {} | CHAR {} | STRING {}
| '!' val {}
| '-' val %prec UMINUS {}
| '*' val %prec VALOF {}
| '&' val %prec ADDROF {}
;




%%
