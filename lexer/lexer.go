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
		if l.input[l.nextPos] == '=' {
			l.readChar()
			tok = &primitives.Token{
				Kind:    primitives.Equal,
				Literal: "==",
			}
		} else {
			tok = &primitives.Token{
				Kind:    primitives.Assign,
				Literal: string(l.currentChar),
			}
		}
	case '<':
		if l.input[l.nextPos] == '=' {
			l.readChar()
			tok = &primitives.Token{
				Kind:    primitives.LessOrEqual,
				Literal: "<=",
			}
		} else {
			tok = &primitives.Token{
				Kind:    primitives.Less,
				Literal: string(l.currentChar),
			}
		}
	case '>':
		if l.input[l.nextPos] == '=' {
			l.readChar()
			tok = &primitives.Token{
				Kind:    primitives.GreaterOrEqual,
				Literal: ">=",
			}
		} else {
			tok = &primitives.Token{
				Kind:    primitives.Greater,
				Literal: string(l.currentChar),
			}
		}

	case '!':
		if l.input[l.nextPos] == '=' {
			l.readChar()
			tok = &primitives.Token{
				Kind:    primitives.NotEqual,
				Literal: "!=",
			}
		} else {
			tok = &primitives.Token{
				Kind:    primitives.Bang,
				Literal: string(l.currentChar),
			}
		}

	case '+':
		tok = &primitives.Token{
			Kind:    primitives.Plus,
			Literal: string(l.currentChar),
		}
	case '-':
		tok = &primitives.Token{
			Kind:    primitives.Minus,
			Literal: string(l.currentChar),
		}
	case '*':
		tok = &primitives.Token{
			Kind:    primitives.Star,
			Literal: string(l.currentChar),
		}
	case '/':
		tok = &primitives.Token{
			Kind:    primitives.Slash,
			Literal: string(l.currentChar),
		}
	case '^':
		tok = &primitives.Token{
			Kind:    primitives.Carrot,
			Literal: string(l.currentChar),
		}
	case '{':
		tok = &primitives.Token{
			Kind:    primitives.OpenCurlyBrace,
			Literal: string(l.currentChar),
		}
	case '}':
		tok = &primitives.Token{
			Kind:    primitives.CloseCurlyBrace,
			Literal: string(l.currentChar),
		}
	case '(':
		tok = &primitives.Token{
			Kind:    primitives.OpenParen,
			Literal: string(l.currentChar),
		}
	case ')':
		tok = &primitives.Token{
			Kind:    primitives.CloseParen,
			Literal: string(l.currentChar),
		}
	case '[':
		tok = &primitives.Token{
			Kind:    primitives.OpenBrackets,
			Literal: string(l.currentChar),
		}
	case ']':
		tok = &primitives.Token{
			Kind:    primitives.CloseBrackets,
			Literal: string(l.currentChar),
		}
	case ',':
		tok = &primitives.Token{
			Kind:    primitives.Comma,
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
