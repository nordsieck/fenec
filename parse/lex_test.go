package parse

import (
	"go/token"
	"testing"

	"github.com/nordsieck/defect"
)

func TestScanner(t *testing.T) {
	s := Scanner{}
	s.Init([]byte(`package main
`))
	expected := []token.Token{token.PACKAGE, token.IDENT, token.SEMICOLON}

	i := 0
	_, tok, _ := s.Scan()
	for tok != token.EOF {
		defect.Equal(t, tok, expected[i])
		_, tok, _ = s.Scan()
		i++
	}
}
