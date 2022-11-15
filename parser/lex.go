package parser

import (
	"bufio"
	"io"
)

type Position struct {
	line int
	col  int
}

type Lexer struct {
	position Position
	reader   *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		position: Position{
			line: 1,
			col:  0,
		},
		reader: bufio.NewReader(reader),
	}
}
