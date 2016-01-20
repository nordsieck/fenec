package build

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/nordsieck/defect"
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
		defect.Equal(t, IsWendigo(k), v)
	}
}

func TestConvertFile(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	defect.Equal(t, err, nil, "Unable to create temporary directory")

	defer func() {
		err := os.RemoveAll(dir)
		defect.Equal(t, err, nil, "Unable to clean up temporary directory")
	}()

	wPath := path.Join(dir, "foo"+wExt)
	wFile, err := os.Create(wPath)
	defect.Equal(t, err, nil, "Unable to create Wendigo file.")

	defer wFile.Close()
	_, err = wFile.WriteString(basicFile)
	defect.Equal(t, err, nil, "Unable to populate Wendigo file.")

	err = ConvertFile(wPath)
	defect.Equal(t, err, nil, "Error in ConvertFile: ", err)

	gFile, err := os.Open(wPath + goExt)
	defect.Equal(t, err, nil, "Unable to open Go file.")

	buffer := make([]byte, 50)
	n, err := gFile.Read(buffer)
	defect.Equal(t, err, nil, "Unable to read from Go file.")
	defect.Equal(t, string(buffer[:n]), basicFile, "Found bad content")
}

func TestConvert(t *testing.T) {
	fn := func(t *testing.T, buffSize int, s string) {
		out := strings.NewReader(s)
		in := &bytes.Buffer{}

		err := convert(out, in, buffSize)
		defect.Equal(t, err, nil)
		defect.Equal(t, string(in.Bytes()), s)
	}

	// 1 pass
	fn(t, 4, "foo")

	// exact
	fn(t, 3, "baz")

	// 2 passes
	fn(t, 2, "bar")
}
