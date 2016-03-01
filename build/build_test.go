package build

import (
	"bytes"
	"io"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/nordsieck/defect"
)

const basicFile = `package foo
`

func TestConvertDir(t *testing.T) {
	bFile := `package bar
`
	dFile := `package c
`

	root := path.Join("foo", "bar")
	files := map[string]*FakeFile{
		path.Join(root, "a"):                 &FakeFile{Buffer: *bytes.NewBufferString("a")},
		path.Join(root, "b"+fExt):            &FakeFile{Buffer: *bytes.NewBufferString(bFile)},
		path.Join(root, "b"+fExt+goExt):      &FakeFile{},
		path.Join(root, "c", "d"+fExt):       &FakeFile{Buffer: *bytes.NewBufferString(dFile)},
		path.Join(root, "c", "d"+fExt+goExt): &FakeFile{},
	}
	infos := map[string][]os.FileInfo{
		root: []os.FileInfo{
			&FileInfo{name: "a", size: 1},
			&FileInfo{name: "b" + fExt, size: int64(len(bFile))},
			&FileInfo{name: "c", isDir: true},
		},
		path.Join(root, "c"): []os.FileInfo{
			&FileInfo{name: "d" + fExt, size: int64(len(dFile))},
		},
	}

	readDir := func(s string) ([]os.FileInfo, error) { return infos[s], nil }
	open := func(s string) (io.ReadWriteCloser, error) { return files[s], nil }
	create := func(s string) (io.ReadWriteCloser, error) { return files[s], nil }

	err := convertDir(root, readDir, open, create)

	defect.Equal(t, err, nil)
	defect.Equal(t, files[path.Join(root, "b"+fExt+goExt)].String(), header+bFile)
	defect.Equal(t, files[path.Join(root, "c", "d"+fExt+goExt)].String(), header+dFile)
}

func TestIsFenec(t *testing.T) {
	tests := map[string]bool{
		"foo":                false,
		"foo" + goExt:        false,
		"foo" + fExt:         true,
		"foo" + fExt + goExt: false,
	}

	for k, v := range tests {
		defect.Equal(t, IsFenec(k), v)
	}
}

func TestInDir(t *testing.T) {
	readDir := func(string) ([]os.FileInfo, error) {
		return []os.FileInfo{
			&FileInfo{name: "foo" + goExt, size: 12, isDir: false},
			&FileInfo{name: "bar" + fExt, size: int64(len(basicFile)), isDir: false},
			&FileInfo{name: "baz", size: 0, isDir: true},
			&FileInfo{name: "quux" + fExt, size: 0, isDir: true},
		}, nil
	}
	p := path.Join("foo", "bar", "baz")
	files, dirs, err := inDir(p, readDir)

	defect.Equal(t, err, nil)
	defect.DeepEqual(t, files, []string{path.Join(p, "bar"+fExt)})
	defect.DeepEqual(t, dirs, []string{
		path.Join(p, "baz"),
		path.Join(p, "quux"+fExt),
	})
}

func TestConvertFile(t *testing.T) {
	name := "/foo/bar/baz" + fExt
	of := FakeFile{}
	o := func(s string) (io.ReadWriteCloser, error) {
		of.Path = s
		of.WriteString(basicFile)
		return &of, nil
	}

	cf := FakeFile{}
	c := func(s string) (io.ReadWriteCloser, error) {
		cf.Path = s
		return &cf, nil
	}
	err := convertFile(name, o, c)

	defect.Equal(t, err, nil)
	defect.Equal(t, cf.String(), header+basicFile)
	defect.Equal(t, of.Closed, true)
	defect.Equal(t, cf.Closed, true)
	defect.Equal(t, of.Path, name)
	defect.Equal(t, cf.Path, name+goExt)
}

func TestConvert(t *testing.T) {
	fn := func(buffSize int, s string) {
		out := strings.NewReader(s)
		in := &bytes.Buffer{}
		err := convert(out, in, buffSize)

		defect.Equal(t, err, nil)
		defect.Equal(t, string(in.Bytes()), header+s)
	}

	fn(4, "foo") // 1 pass
	fn(3, "baz") // exact
	fn(2, "bar") // 2 passes
}
