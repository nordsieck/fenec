package build

import (
	"io"
	"os"
	"path"
	"strings"
)

const (
	wExt  = ".w"
	goExt = ".go"

	buffSize = 4 * 1024
)

// ConvertDir converts all Wendigo files in the directory to Go files.
func ConvertDir(dirPath string) error {
	return nil
}

func convertDir(dirPath string, readDir ReadDirFunc, open, create FileFunc) error {
	dirs := []string{dirPath}

	for len(dirs) > 0 {
		dir := dirs[0]
		dirs[0] = dirs[len(dirs)-1]
		dirs = dirs[:len(dirs)-1]

		localFiles, localDirs, err := inDir(dir, readDir)
		if err != nil {
			return err
		}
		for _, localFile := range localFiles {
			if err = convertFile(localFile, open, create); err != nil {
				return err
			}
		}

		for _, localDir := range localDirs {
			dirs = append(dirs, localDir)
		}
	}
	return nil
}

func inDir(dirPath string, readDir ReadDirFunc) (files, dirs []string, err error) {
	fileInfos, err := readDir(dirPath)
	if err != nil {
		return nil, nil, err
	}

	files = []string{}
	dirs = []string{}
	for _, file := range fileInfos {
		if file.IsDir() {
			dirs = append(dirs, path.Join(dirPath, file.Name()))
		} else if IsWendigo(file.Name()) {
			files = append(files, path.Join(dirPath, file.Name()))
		}
	}
	return files, dirs, nil
}

func convertFile(filePath string, open, create FileFunc) error {
	wFile, err := open(filePath)
	if err != nil {
		return err
	}
	defer wFile.Close()

	goFile, err := create(filePath + goExt)
	if err != nil {
		return err
	}
	defer goFile.Close()

	return Convert(wFile, goFile)
}

func wrap(fn func(string) (*os.File, error)) FileFunc {
	return func(s string) (io.ReadWriteCloser, error) { return fn(s) }
}

// Convert converts Wendigo source code to Go.
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

type FileFunc func(string) (io.ReadWriteCloser, error)
type ReadDirFunc func(string) ([]os.FileInfo, error)
