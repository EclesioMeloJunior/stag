package lexer

import (
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

	l.skipWhiteSpace()

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
	case ';':
		tok = &primitives.Token{
			Kind:    primitives.Semicolon,
			Literal: string(l.currentChar),
		}
	case 0:
		tok = &primitives.Token{
			Kind: primitives.EOF,
		}
	default:
		if isLetter(l.currentChar) {
			return l.readIdentifierOrKeyword()
		}

		if isNumber(l.currentChar) {
			return l.readNumber()
		}

		return &primitives.Token{Kind: primitives.Illegal, Literal: string(l.currentChar)}
	}

	l.readChar()

	return tok
}

func (l *Lexer) readIdentifierOrKeyword() *primitives.Token {
	literal := l.readLiteral(isLetter)

	if _, ok := Keywords[literal]; ok {
		return &primitives.Token{
			Kind:    primitives.Keyword,
			Literal: literal,
		}
	}

	return &primitives.Token{
		Kind:    primitives.Ident,
		Literal: literal,
	}
}

func (l *Lexer) readNumber() *primitives.Token {
	literal := l.readLiteral(isNumber)

	return &primitives.Token{
		Kind:    primitives.Number,
		Literal: literal,
	}
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

func (l *Lexer) skipWhiteSpace() {
	for l.currentChar == ' ' || l.currentChar == '\t' || l.currentChar == '\n' || l.currentChar == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readLiteral(cond func(ch byte) bool) string {
	startPos := l.pos
	for cond(l.currentChar) {
		l.readChar()
	}

	return l.input[startPos:l.pos]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
