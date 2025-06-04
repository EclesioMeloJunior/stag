package lexer

import (
	"fmt"
	"stag/primitives"
)

type Lexer struct {
	input       string
	pos         int
	nextPos     int
	currentChar byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() *primitives.Token {
	var tok *primitives.Token

	switch l.currentChar {
	case '=':
		tok = &primitives.Token{
			Kind:    primitives.Equals,
			Literal: string(l.currentChar),
		}
	case '+':
		tok = &primitives.Token{
			Kind:    primitives.Plus,
			Literal: string(l.currentChar),
		}
	case 0:
		tok = &primitives.Token{
			Kind: primitives.EOF,
		}
	default:
		panic(fmt.Sprintf("char not supported: %v", l.currentChar))
	}

	l.readChar()

	return tok
}

func (l *Lexer) readChar() {
	if l.nextPos >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.nextPos]
	}

	l.pos = l.nextPos
	l.nextPos++
}
