package testutil

import "bytes"

type FakeFile struct {
	Path   string
	Closed bool
	bytes.Buffer
}

func (f *FakeFile) Close() error {
	f.Closed = true
	return nil
}
