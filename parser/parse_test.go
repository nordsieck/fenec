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
	comment = `// comment
package main
`
	comments = `// comment
package main

// foo
// bar
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

func TestParseFile_Comment(t *testing.T) {
	fs := token.NewFileSet()
	file, err := ParseFile(fs, "a.go", comment)
	defect.Equal(t, err, nil)
	comment := &ast.Comment{Slash: 1, Text: "// comment"}
	defect.DeepEqual(t, *file.File, ast.File{
		Doc:      &ast.CommentGroup{List: []*ast.Comment{comment}},
		Package:  token.Pos(12),
		Name:     &ast.Ident{NamePos: 20, Name: "main"},
		Scope:    &ast.Scope{Objects: map[string]*ast.Object{}},
		Comments: []*ast.CommentGroup{{List: []*ast.Comment{comment}}},
	})
}

func TestParseFile_Comments(t *testing.T) {
	fs := token.NewFileSet()
	file, err := ParseFile(fs, "a.go", comments)
	defect.Equal(t, err, nil)
	top := &ast.CommentGroup{List: []*ast.Comment{{Slash: 1, Text: "// comment"}}}
	next := &ast.CommentGroup{List: []*ast.Comment{
		{Slash: 26, Text: "// foo"},
		{Slash: 33, Text: "// bar"},
	}}

	defect.DeepEqual(t, *file.File, ast.File{
		Doc:      top,
		Package:  token.Pos(12),
		Name:     &ast.Ident{NamePos: 20, Name: "main"},
		Scope:    &ast.Scope{Objects: map[string]*ast.Object{}},
		Comments: []*ast.CommentGroup{top, next},
	})
}

func TestParseFile_Hello(t *testing.T) {
	fs := token.NewFileSet()
	_, err := ParseFile(fs, "a.go", hello)
	defect.Equal(t, err, nil)
}
