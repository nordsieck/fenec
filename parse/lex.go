//go:generate go tool yacc -o parse.go parse.y

package parse

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io"
	"io/ioutil"
)

const (
	ErrFmt = "%v %v:%v | %v\n"

	EOF = 0 // requiretd by YACC
)

// examples taken from go/token
var t = map[token.Token]int{
	token.EOF:     EOF,
	token.COMMENT: COMMENT,

	token.IDENT:  IDENT,  // main
	token.INT:    INT,    // 12345
	token.FLOAT:  FLOAT,  // 123.45
	token.IMAG:   IMAG,   // 123,45i
	token.CHAR:   CHAR,   // 'a'
	token.STRING: STRING, // "abc"

	token.ADD: '+',
	token.SUB: '-',
	token.MUL: '*',
	token.QUO: '/',
	token.REM: '%',

	token.AND:     '&',
	token.OR:      '|',
	token.XOR:     '^',
	token.SHL:     SHL,     // <<
	token.SHR:     SHR,     // >>
	token.AND_NOT: AND_NOT, // &^

	token.ADD_ASSIGN: ADD_ASSIGN, // +=
	token.SUB_ASSIGN: SUB_ASSIGN, // -=
	token.MUL_ASSIGN: MUL_ASSIGN, // *=
	token.QUO_ASSIGN: QUO_ASSIGN, // /=
	token.REM_ASSIGN: REM_ASSIGN, // %=

	token.AND_ASSIGN:     AND_ASSIGN,     // &=
	token.OR_ASSIGN:      OR_ASSIGN,      // |=
	token.XOR_ASSIGN:     XOR_ASSIGN,     // ^=
	token.SHL_ASSIGN:     SHL_ASSIGN,     // <<=
	token.SHR_ASSIGN:     SHR_ASSIGN,     // >>=
	token.AND_NOT_ASSIGN: AND_NOT_ASSIGN, // &^=

	token.LAND:  LAND,  // &&
	token.LOR:   LOR,   // ||
	token.ARROW: ARROW, // <-
	token.INC:   INC,   // ++
	token.DEC:   DEC,   // --

	token.EQL:    EQL, // ==
	token.LSS:    '<',
	token.GTR:    '>',
	token.ASSIGN: '=',
	token.NOT:    '!',

	token.NEQ:      NEQ,      // !=
	token.LEQ:      LEQ,      // <=
	token.GEQ:      GEQ,      // >=
	token.DEFINE:   DEFINE,   // :=
	token.ELLIPSIS: ELLIPSIS, // ...

	token.LPAREN: '(',
	token.LBRACK: '[',
	token.LBRACE: '{',
	token.COMMA:  ',',
	token.PERIOD: '.',

	token.RPAREN:    ')',
	token.RBRACK:    ']',
	token.RBRACE:    '}',
	token.SEMICOLON: ';',
	token.COLON:     ':',

	token.BREAK:    BREAK,
	token.CASE:     CASE,
	token.CHAN:     CHAN,
	token.CONST:    CONST,
	token.CONTINUE: CONTINUE,

	token.DEFAULT:     DEFAULT,
	token.DEFER:       DEFER,
	token.ELSE:        ELSE,
	token.FALLTHROUGH: FALLTHROUGH,
	token.FOR:         FOR,

	token.FUNC:   FUNC,
	token.GO:     GO,
	token.GOTO:   GOTO,
	token.IF:     IF,
	token.IMPORT: IMPORT,

	token.INTERFACE: INTERFACE,
	token.MAP:       MAP,
	token.PACKAGE:   PACKAGE,
	token.RANGE:     RANGE,
	token.RETURN:    RETURN,

	token.SELECT: SELECT,
	token.STRUCT: STRUCT,
	token.SWITCH: SWITCH,
	token.TYPE:   TYPE,
	token.VAR:    VAR,
}

type Lexer struct {
	s    scanner.Scanner
	fSet *token.FileSet
}

func (l *Lexer) Init(name string, rc io.ReadCloser) error {
	buf, err := ioutil.ReadAll(rc)
	if err != nil {
		return err
	}

	l.fSet = token.NewFileSet()
	file := l.fSet.AddFile(name, l.fSet.Base(), len(buf))
	l.s.Init(file, buf, ErrHandler, 0)
	return nil
}

func (l *Lexer) Lex(lval *yySymType) int {
	tok := Token{}
	var pos token.Pos
	pos, tok.Token, tok.Literal = l.s.Scan()
	tok.Pos = l.fSet.Position(pos)
	return t[tok.Token]
}

func (l *Lexer) Error(e string) { fmt.Println(e) }

type Token struct {
	Pos     token.Position
	Token   token.Token
	Literal string
}

func ErrHandler(pos token.Position, msg string) {
	fmt.Printf(ErrFmt, pos.Filename, pos.Line, pos.Column, msg)
}
