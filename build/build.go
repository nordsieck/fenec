package build

import (
	"io"
	"os"
	"strings"
)

const (
	wExt  = ".w"
	goExt = ".go"

	buffSize = 4 * 1024
)

// ConvertFile reads in a Wendigo file and creates a corresponding Go file
func ConvertFile(filePath string) error {
	wFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer wFile.Close()

	goFile, err := os.Create(filePath + goExt)
	if err != nil {
		return err
	}
	defer goFile.Close()

	return Convert(wFile, goFile)
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
