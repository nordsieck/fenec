package build

import (
	"io"
	"strings"
)

const (
	wExt  = ".w"
	goExt = ".go"

	buffSize = 4 * 1024
)

func ConvertFile(filePath string) error {
	return nil
}

// Convert converts Wendigo source code to Go
func Convert(r io.Reader, w io.Writer) error { return convert(r, w, buffSize) }
func convert(r io.Reader, w io.Writer, buffSize int) error {
	buffer := make([]byte, buffSize)

	for {
		n, err := r.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		n, err = w.Write(buffer[:n])
		if err != nil {
			return err
		}
	}
	return nil
}

func IsWendigo(name string) bool { return strings.HasSuffix(name, wExt) }
