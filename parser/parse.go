package parser

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/nordsieck/defect"
)

const (
	ErrMultiplePackages = defect.Error("Multiple packages")
	ErrNoPackages       = defect.Error("No packages")
)

type File struct {
	*ast.File
	FileName string
}

func ParseFile(fset *token.FileSet, filename, source string) (*File, error) {
	file, err := parser.ParseFile(fset, filename, []byte(source),
		parser.ParseComments|parser.AllErrors)
	if err != nil {
		return nil, err
	}
	return &File{File: file, FileName: filename}, nil
}

// TODO: set name, scope and imports
func ParsePackage(files []*File) (*ast.Package, error) {
	pkgs := map[string]*ast.Package{}
	for _, file := range files {
		pkgs[file.Name.Name].Files[file.FileName] = file.File
	}
	if len(pkgs) > 1 {
		return nil, ErrMultiplePackages
	}
	for _, pkg := range pkgs {
		return pkg, nil
	}
	return nil, ErrNoPackages
}
