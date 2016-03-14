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
		`type a b.c`,
		`type (a b;)`,
		`type (a b.c;)`,
		`type a [3]int`,
		`type a []int`,
		`type a *int`,
		`type a map[int]int`,
		`type a chan int`,
		`type a <-chan int`,
		`type a chan<- int`,
		`type a chan<- chan int`,
		`type a chan<- <-chan int`,
		`type a <-chan <-chan int`,
		`type a struct{}`,
		`type a struct{b}`,
		`type a struct{b;}`,
		`type a struct{*b;}`,
		`type a struct{*b "foo"}`,
		`type a struct{b int}`,
		`type a struct{b, c int}`,
		`type a struct{b, c int "foo";}`,
		`type a interface{}`,
		`type a interface{b; c.d;}`,
		`type a interface{b()}`,
		`type a interface{b()();}`,
		`type a interface{b(int,)}`,
		`type a interface{b(i, j int)}`,
		`type a interface{b(i ...int)}`,
		`type a func()`,
		`type a func()()`,
		`type a func(i, j, k int) (int, int, int)`,

		// const
		`const a = 1`,
		`const a, b = 1, 2`,
		`const ()`,
		`const (a, b = 1, 2;)`,
		`const a int`,
		`const a, b int`,
		`const (a, b int = 1, 2;)`,
	} {
		testFn(t, prog)
	}
}

func testFn(t *testing.T, prog string) {
	ff := &testutil.FakeFile{Buffer: *bytes.NewBufferString(prog)}
	l := &Lexer{}
	l.Init("", ff)

	defect.Equal(t, yyParse(l), EOF)
}
