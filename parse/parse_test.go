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
