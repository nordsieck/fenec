package build

import (
	"bytes"
	"strings"
	"testing"
)

func TestIsWendigo(t *testing.T) {
	tests := map[string]bool{
		"foo":                false,
		"foo" + goExt:        false,
		"foo" + wExt:         true,
		"foo" + wExt + goExt: false,
	}

	for k, v := range tests {
		if IsWendigo(k) != v {
			t.Fail()
		}
	}
}

func TestConvert(t *testing.T) {
	fn := func(t *testing.T, buffSize int, s string) {
		out := strings.NewReader(s)
		in := &bytes.Buffer{}

		err := convert(out, in, buffSize)
		if err != nil {
			t.Fail()
		}
		if string(in.Bytes()) != s {
			t.Fail()
		}
	}

	// 1 pass
	fn(t, 4, "foo")

	// exact
	fn(t, 3, "baz")

	// 2 passes
	fn(t, 2, "bar")
}
