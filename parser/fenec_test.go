// Test the fenec specific grammer in the parser

package parser

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/nordsieck/defect"
)

func TestEmptyGenerics_literals(t *testing.T) {
	src := []string{
		"struct{}{}",   // empty struct
		"struct[]{}{}", // empty generic struct

		"func(){}",   // empty function
		"func[](){}", // empty generic function
	}

	for _, s := range src {
		_, err := ParseExpr(s)
		defect.Equal(t, err, nil)
	}
}

func TestEmptyGenerics_vars(t *testing.T) {
	preamble := "package main;"

	src := []string{
		"func f(){}",         // empty function
		"func f[](){}",       // empty generic function
		"func (r R) f(){}",   // empty method
		"func (r R) f[](){}", // empty generic method
	}

	for _, s := range src {
		_ = ParseTestFile(t, preamble+s)
	}
}

func ParseTestFile(t *testing.T, s string) *ast.File {
	fset := token.NewFileSet()
	file, err := ParseFile(fset, "", s, AllErrors)
	defect.Equal(t, err, nil)
	return file
}
