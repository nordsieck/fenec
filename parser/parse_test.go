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
	simple = `package main

func main() {}
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

func TestParseFile_Simple(t *testing.T) {
	fs := token.NewFileSet()
	file, err := ParseFile(fs, "a.go", simple)
	defect.Equal(t, err, nil)

	block := &ast.BlockStmt{Lbrace: 27, Rbrace: 28}
	obj := &ast.Object{Kind: ast.Fun, Name: "main"}
	fnDecl := &ast.FuncDecl{
		Name: &ast.Ident{Name: "main", NamePos: 20, Obj: obj},
		Type: &ast.FuncType{
			Func:   15,
			Params: &ast.FieldList{Opening: 24, Closing: 25},
		},
		Body: block,
	}
	obj.Decl = fnDecl

	defect.DeepEqual(t, file.File, &ast.File{
		Package: 1,
		Name:    &ast.Ident{NamePos: 9, Name: "main"},
		Decls:   []ast.Decl{fnDecl},
		Scope:   &ast.Scope{Objects: map[string]*ast.Object{"main": obj}},
	})
}

func TestParseFile_Hello(t *testing.T) {
	fs := token.NewFileSet()
	file, err := ParseFile(fs, "a.go", hello)
	defect.Equal(t, err, nil)

	expectedDecl := &ast.BlockStmt{
		Lbrace: 41,
		List: []ast.Stmt{
			&ast.ExprStmt{
				X: &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   &ast.Ident{NamePos: 44, Name: "fmt"},
						Sel: &ast.Ident{NamePos: 48, Name: "Println"},
					},
					Lparen: 55,
					Args: []ast.Expr{&ast.BasicLit{
						ValuePos: 56,
						Kind:     token.STRING,
						Value:    `"hello world"`,
					}},
					Rparen: 69,
				},
			},
		},
		Rbrace: 71,
	}

	obj := &ast.Object{
		Kind: ast.Fun,
		Name: "main",
		Decl: expectedDecl,
	}

	fnDecl := &ast.FuncDecl{
		Name: &ast.Ident{Name: "main", NamePos: 34, Obj: obj},
		Type: &ast.FuncType{
			Func:   29,
			Params: &ast.FieldList{Opening: 38, Closing: 39},
		},
		Body: expectedDecl,
	}

	obj.Decl = fnDecl

	importSpec := &ast.ImportSpec{Path: &ast.BasicLit{
		ValuePos: 22,
		Kind:     token.STRING,
		Value:    `"fmt"`,
	}}

	genDecl := &ast.GenDecl{
		TokPos: 15,
		Tok:    token.IMPORT,
		Specs:  []ast.Spec{importSpec},
	}

	defect.DeepEqual(t, file.File, &ast.File{
		Package:    1,
		Name:       &ast.Ident{NamePos: 9, Name: "main"},
		Decls:      []ast.Decl{genDecl, fnDecl},
		Scope:      &ast.Scope{Objects: map[string]*ast.Object{"main": obj}},
		Imports:    []*ast.ImportSpec{importSpec},
		Unresolved: []*ast.Ident{{NamePos: 44, Name: "fmt"}},
	})
}
