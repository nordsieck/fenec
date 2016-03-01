package lex

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/nordsieck/defect"
	"github.com/nordsieck/fenec/testutil"
)

func TestLexGo(t *testing.T) {
	ff := testutil.FakeFile{Buffer: *bytes.NewBufferString(`package foo

const a = 1
const b = "foo"

type C int

// d
type D struct {
	e string
}
func (d *D) String() string { return d.e }

type Stringer interface {
	String() string
}

/* ensure the interface is implemented */
var _ Stringer = D{}

func t(b bool) bool { return true }
`)}
	_, err := Lex("foo", &ff)
	defect.Equal(t, err, nil)
}

func TestLexError(t *testing.T) {
	ff := testutil.FakeFile{Buffer: *bytes.NewBufferString(`"`)}
	_, err := Lex("foo", &ff)
	defect.DeepEqual(t, err, fmt.Errorf(ErrFmt, "foo", 1, 1, "string literal not terminated"))
}

func TestLexWendigo(t *testing.T) {
	ff := testutil.FakeFile{Buffer: *bytes.NewBufferString(`package foo

type B int

func a<t>(j t) t { return j }
`)}
	_, err := Lex("foo", &ff)
	defect.Equal(t, err, nil)
}
