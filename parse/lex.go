package parse

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io"
	"io/ioutil"
)

const ErrFmt = "%v %v:%v | %v\n"

type Lexer struct {
	s    scanner.Scanner
	fSet *token.FileSet
}

func (l *Lexer) Init(name string, rc io.ReadCloser) error {
	buf, err := ioutil.ReadAll(rc)
	if err != nil {
		return err
	}

	l.fSet = token.NewFileSet()
	file := l.fSet.AddFile(name, l.fSet.Base(), len(buf))
	l.s.Init(file, buf, ErrHandler, 0)
	return nil
}

func (l *Lexer) Lex(lval *yySymType) int {
	tok := Token{}
	var pos token.Pos
	pos, tok.Token, tok.Literal = l.s.Scan()
	tok.Pos = l.fSet.Position(pos)
	return int(tok.Token)
}

func (l *Lexer) Error(e string) { fmt.Println(e) }

type Token struct {
	Pos     token.Position
	Token   token.Token
	Literal string
}

func ErrHandler(pos token.Position, msg string) {
	fmt.Printf(ErrFmt, pos.Filename, pos.Line, pos.Column, msg)
}

type yySymType struct{}
