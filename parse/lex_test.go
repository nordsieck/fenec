package parse

import (
	"go/token"
	"testing"

	"github.com/nordsieck/defect"
)

func TestScanner(t *testing.T) {
	s := Scanner{}

	expected := map[string][]token.Token{
		`package main
`: {token.PACKAGE, token.IDENT, token.SEMICOLON},
		`package main

import "fmt"

func main(){
	fmt.Println("hello world")
}
`: {token.PACKAGE, token.IDENT, token.SEMICOLON,
			token.IMPORT, token.STRING, token.SEMICOLON,
			token.FUNC, token.IDENT, token.LPAREN, token.RPAREN, token.LBRACE,
			token.IDENT, token.PERIOD, token.IDENT, token.LPAREN, token.STRING, token.RPAREN, token.SEMICOLON,
			token.RBRACE, token.SEMICOLON},
	}

	for text, symbols := range expected {
		s.Init([]byte(text))
		i := 0
		_, tok, _ := s.Scan()
		for tok != token.EOF {
			defect.Equal(t, tok, symbols[i])
			_, tok, _ = s.Scan()
			i++
		}
	}
}
