package parse

import (
	"bytes"
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

	for _, expected := range []int{
		PACKAGE, IDENT, ';',
		TYPE, IDENT, IDENT, ';',
		FUNC, IDENT, '(', IDENT, IDENT, ')', IDENT, '{', RETURN, IDENT, '}', ';',
		EOF,
	} {
		v := l.Lex(&yySymType{})
		defect.Equal(t, v, expected)
	}
}
