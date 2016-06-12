package parse

import (
	"fmt"
	"go/scanner"
	"go/token"
)

type Scanner struct {
	s scanner.Scanner

	pos token.Pos
	tok token.Token
	lit string
}

func (s *Scanner) Init(source []byte) {
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(source))
	s.s.Init(file, source, errorHandler, scanner.ScanComments)
	s.Scan()
}

func (s *Scanner) Scan() (pos token.Pos, tok token.Token, lit string) {
	pos, tok, lit = s.pos, s.tok, s.lit

	if s.tok != token.EOF {
		s.pos, s.tok, s.lit = s.s.Scan()
	}
	return pos, tok, lit
}

func (s *Scanner) Peek() (pos token.Pos, tok token.Token, lit string) { return s.pos, s.tok, s.lit }

func errorHandler(pos token.Position, msg string) { fmt.Println(pos, msg) }
