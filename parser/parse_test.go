package parser

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/nordsieck/defect"
)

const (
	minimal = `package main
`

	hello = `package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
`
)

func TestParseFile_Minimal(t *testing.T) {
	fs := token.NewFileSet()
	file, err := ParseFile(fs, "a.go", minimal)
	defect.Equal(t, err, nil)
	defect.DeepEqual(t, *file.File, ast.File{
		Package: token.Pos(1),
		Name:    &ast.Ident{NamePos: 9, Name: "main"},
		Scope:   &ast.Scope{Objects: map[string]*ast.Object{}},
	})
}

func TestParseFile_Hello(t *testing.T) {
	fs := token.NewFileSet()
	_, err := ParseFile(fs, "a.go", hello)
	defect.Equal(t, err, nil)
}
