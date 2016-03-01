package lex

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io"
	"io/ioutil"
)

const ErrFmt = "%v %v:%v | %v"

func Lex(name string, r io.ReadCloser) ([]Token, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	s := scanner.Scanner{}
	fSet := token.NewFileSet()
	file := fSet.AddFile(name, fSet.Base(), len(buf))
	errChan, errHandler := ErrHandler()
	s.Init(file, buf, errHandler, 0)

	ret := []Token{}
	pos := token.Pos(0)
	tok := Token{}
forloop:
	for {
		select {
		case err = <-errChan:
			return nil, err
		default:
			pos, tok.Token, tok.Literal = s.Scan()
			if tok.Token == token.EOF {
				break forloop
			}
			tok.Pos = fSet.Position(pos)
			ret = append(ret, tok)
		}
	}
	return ret, nil
}

type Token struct {
	Pos     token.Position
	Token   token.Token
	Literal string
}

func ErrHandler() (<-chan error, scanner.ErrorHandler) {
	c := make(chan error, 1)
	eh := func(pos token.Position, msg string) {
		// TODO: make this error message as close as possible to the go one
		c <- fmt.Errorf(ErrFmt, pos.Filename, pos.Line, pos.Column, msg)
	}
	return c, eh
}
