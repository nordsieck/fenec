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
		`package foo`,

		// import
		`import "foo"`,
		`import . "foo"`,
		`import _ "foo"`,
		`import f "foo"`,
		`import ()`,
		`import ("foo")`,
		`import ("foo";)`,
		`import (. "foo";)`,

		// type
		`type a b`,
		`type (a b;)`,
		`type (a b.c;)`,

		// const
		`const a = 1`,
		`const a, b = 1, 2`,
		`const ()`,
		`const (a, b = 1, 2;)`,
		`const a int`,
		`const a, b int`,
		`const (a, b int = 1, 2;)`,

		// var
		`var a = 1`,
		`var a, b = 1, 2`,
		`var ()`,
		`var (a, b = 1, 2;)`,
		`var a int`,
		`var a, b int`,
		`var (a, b int = 1, 2;)`,

		// function
		`func a(){}`,
		`func a()(){}`,
		`func a(int) int {}`,
		`func a(i, j int, s string) (rune, err) {}`,

		// method
		`func (a A) b(){}`,
		`func (a *A) b(){}`,
	} {
		testFn(t, prog)
	}
}

func TestYyParse_Type(t *testing.T) {
	for _, typ := range []string{
		`b`,
		`b.c`,
		`[3]int`,
		`[]int`,
		`*int`,
		`map[int]int`,
		`chan int`,
		`<-chan int`,
		`chan<- int`,
		`chan<- chan int`,
		`chan<- <-chan int`,
		`<-chan <-chan int`,
		`struct{}`,
		`struct{b}`,
		`struct{b;}`,
		`struct{*b;}`,
		`struct{*b "foo"}`,
		`struct{b int}`,
		`struct{b, c int}`,
		`struct{b, c int "foo";}`,
		`interface{}`,
		`interface{b; c.d;}`,
		`interface{b()}`,
		`interface{b()();}`,
		`interface{b(int,)}`,
		`interface{b(i, j int)}`,
		`interface{b(i ...int)}`,
		`func()`,
		`func()()`,
		`func(i, j, k int) (int, int, int)`,
	} {
		testFn(t, "type a "+typ)
		testFn(t, "var a "+typ)
	}
}

func testFn(t *testing.T, prog string) {
	ff := &testutil.FakeFile{Buffer: *bytes.NewBufferString(prog)}
	l := &Lexer{}
	l.Init("", ff)

	defect.Equal(t, yyParse(l), EOF)
}
