package build

import (
	"bytes"
	"os"
	"time"
)

var _ os.FileInfo = &FileInfo{}

type FileInfo struct {
	name  string
	size  int64
	isDir bool
}

func (f *FileInfo) Name() string       { return f.name }
func (f *FileInfo) Size() int64        { return f.size }
func (f *FileInfo) ModTime() time.Time { return time.Time{} }
func (f *FileInfo) IsDir() bool        { return f.isDir }
func (f *FileInfo) Sys() interface{}   { return nil }
func (f *FileInfo) Mode() os.FileMode {
	if f.isDir {
		return os.ModeDir
	}
	return os.FileMode(0)
}

type FakeFile struct {
	Path   string
	Closed bool
	bytes.Buffer
}

func (f *FakeFile) Close() error {
	f.Closed = true
	return nil
}
