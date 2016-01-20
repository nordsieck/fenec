package build

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

const basicFile = `package foo
`

func TestIsWendigo(t *testing.T) {
	tests := map[string]bool{
		"foo":                false,
		"foo" + goExt:        false,
		"foo" + wExt:         true,
		"foo" + wExt + goExt: false,
	}

	for k, v := range tests {
		if IsWendigo(k) != v {
			t.FailNow()
		}
	}
}

func TestConvertFile(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal("Unable to create temporary directory")
	}

	defer func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Fatal("Unable to clean up temporary directory")
		}
	}()

	wPath := path.Join(dir, "foo"+wExt)
	wFile, err := os.Create(wPath)
	if err != nil {
		t.Fatal("Unable to create Wendigo file.")
	}
	defer wFile.Close()
	if _, err = wFile.WriteString(basicFile); err != nil {
		t.Fatal("Unable to populate Wendigo file.")
	}

	if err = ConvertFile(wPath); err != nil {
		t.Fatal("Error in ConvertFile: ", err)
	}

	gFile, err := os.Open(wPath + goExt)
	if err != nil {
		t.Fatal("Unable to open Go file.")
	}
	buffer := make([]byte, 50)
	n, err := gFile.Read(buffer)
	if err != nil {
		t.Fatal("Unable to read from Go file.")
	}
	if string(buffer[:n]) != basicFile {
		t.Fatalf(`Found bad content.
Expected: %v
Got: %v
`, basicFile, string(buffer[:n]))
	}
}

func TestConvert(t *testing.T) {
	fn := func(t *testing.T, buffSize int, s string) {
		out := strings.NewReader(s)
		in := &bytes.Buffer{}

		err := convert(out, in, buffSize)
		if err != nil {
			t.FailNow()
		}
		if string(in.Bytes()) != s {
			t.FailNow()
		}
	}

	// 1 pass
	fn(t, 4, "foo")

	// exact
	fn(t, 3, "baz")

	// 2 passes
	fn(t, 2, "bar")
}
