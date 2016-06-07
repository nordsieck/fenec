package parse

import (
	"fmt"
	"go/scanner"
	"go/token"
)

type Scanner struct{ s scanner.Scanner }

func (s *Scanner) Init(source []byte) {
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(source))
	s.s.Init(file, source, errorHandler, scanner.ScanComments)
}

func (s *Scanner) Scan() (pos token.Pos, tok token.Token, lit string) { return s.s.Scan() }

func errorHandler(pos token.Position, msg string) { fmt.Println(pos, msg) }
