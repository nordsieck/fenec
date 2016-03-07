package parse

import (
	"bytes"
	"testing"

	"github.com/nordsieck/defect"
	"github.com/nordsieck/fenec/testutil"
)

func TestYyParse(t *testing.T) {
	for _, prog := range []string{
		// package
		"package foo",

		// const
		"const a = 1",
		"const a, b = 1, 2",
		`const a = "foo"`,
		`const a, b = 1, "foo"`,
		"const ()",
		`const ( a = 1 )`,
		`const ( a = 1; )`,
		`const ( a, b = 1, 2; c = "foo" )`,
		`const ( a, b int = 1, 2 )`,

		// var
		"var a = 1",
		"var a, b = 1, 2",
		`var a = "foo"`,
		`var a, b = 1, "foo"`,
		"var ()",
		`var ( a = 1 )`,
		`var ( a = 1; )`,
		`var ( a, b = 1, 2; c = "foo" )`,
		`var ( a, b int = 1, 2 )`,
		`var a, b int`,

		// type
		"type a b",

		// comment
		"// foo",
		"/* bar */",
		"var a int // foo",
		"// /* a */",
		"/* // a */",

		// ++/--
		"a++",
		"b--",
	} {
		testFn(t, prog)
	}
}

// Test the different ways types can be used in go
func TestYyParse_Type(t *testing.T) {
	for _, typ := range []string{
		"int",
		"b",
		"***int",
		"[]string",
		"[3]bool",
		"['b']int",
		`func()`,
		`func(func(), int,)`,
		`func() ()`,
		`func() (int, int,)`,
		`func(func() func() func()) func() func() func()`,
		`func(b, c int, d, e string,) (int, int, int)`,
		`func(...int)`,
		`func(a ...int)`,
		`func(a, b string, c ...int,)`,
		`*func()`,
		`*func(*func() *func(*func()))`,
		`struct{}`,
		`struct{a int;}`,
		`struct{b, c int; s string}`,
		`interface{}`,
		`interface{B()}`,
		`interface{B(*func()) (a, b string)}`,
	} {
		testFn(t, "var a "+typ)
		testFn(t, "type a "+typ)
	}
}

func testFn(t *testing.T, prog string) {
	ff := &testutil.FakeFile{Buffer: *bytes.NewBufferString(prog)}
	l := &Lexer{}
	l.Init("", ff)

	defect.Equal(t, yyParse(l), 0)
}
