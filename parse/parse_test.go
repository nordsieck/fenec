package parse

import (
	"bytes"
	"testing"

	"github.com/nordsieck/defect"
	"github.com/nordsieck/fenec/testutil"
)

func TestYyParse(t *testing.T) {
	fn := func(prog string) {
		ff := &testutil.FakeFile{Buffer: *bytes.NewBufferString(prog)}
		l := &Lexer{}
		l.Init("", ff)

		defect.Equal(t, yyParse(l), 0)
	}

	for _, prog := range []string{
		// package
		"package foo",

		// const
		"const a = 1",
		"const a, b = 1, 2",
		`const a = "foo"`,
		`const a, b = 1, "foo"`,
		`const ( a = 1 )`,
		`const ( a = 1; )`,
		`const ( a, b = 1, 2; c = "foo" )`,

		// var
		"var a = 1",
		"var a, b = 1, 2",
		`var a = "foo"`,
		`var a, b = 1, "foo"`,
		`var ( a = 1 )`,
		`var ( a = 1; )`,
		`var ( a, b = 1, 2; c = "foo" )`,
	} {
		fn(prog)
	}
}
