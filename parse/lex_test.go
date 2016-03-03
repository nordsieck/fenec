package parse

import (
	"bytes"
	"go/token"
	"testing"

	"github.com/nordsieck/defect"
	"github.com/nordsieck/fenec/testutil"
)

func TestLexGo(t *testing.T) {
	ff := &testutil.FakeFile{Buffer: *bytes.NewBufferString(`package foo

type B int

func a(j t) t { return j }
`)}
	l := Lexer{}
	l.Init("foo", ff)

	for _, expected := range []token.Token{
		token.PACKAGE, token.IDENT, token.SEMICOLON,
		token.TYPE, token.IDENT, token.IDENT, token.SEMICOLON,
		token.FUNC, token.IDENT, token.LPAREN, token.IDENT,
		token.IDENT, token.RPAREN, token.IDENT, token.LBRACE,
		token.RETURN, token.IDENT, token.RBRACE, token.SEMICOLON,
		token.EOF,
	} {
		v := l.Lex(&yySymType{})
		defect.Equal(t, token.Token(v), expected)
	}
}
