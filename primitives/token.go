package primitives

import "fmt"

type TokenKind uint8

func (tk TokenKind) String() string {
	switch tk {
	case Keyword:
		return "Keyword"
	case Ident:
		return "Ident"
	case Number:
		return "Number"
	case Equals:
		return "Equals"
	case Plus:
		return "Plus"
	case Minus:
		return "Minus"
	case Star:
		return "Star"
	case Slash:
		return "Slash"
	case Carrot:
		return "Carrot"
	case Sqrt:
		return "Sqrt"
	case OpenBrackets:
		return "OpenBrackets"
	case CloseBrackets:
		return "CloseBrackets"
	case OpenParen:
		return "OpenParen"
	case CloseParen:
		return "CloseParen"
	case Comma:
		return "Comma"
	case Semicolon:
		return "Semicolon"
	case Illegal:
		return "Illegal"
	case EOF:
		return "EOF"
	default:
		return fmt.Sprintf("Unknown(%d)", tk)
	}
}

const (
	Keyword TokenKind = iota
	Ident
	Number

	Equals // =
	Plus   // +
	Minus  // -
	Star   // *
	Slash  // /
	Carrot // ^
	Sqrt   // √

	OpenBrackets  // {
	CloseBrackets // }

	OpenParen  // (
	CloseParen // )

	Comma     // ,
	Semicolon // ;

	Illegal

	EOF
)

type Token struct {
	Kind         TokenKind
	Literal      string
	SourceColumn int
	SourceLine   int
}

func (t *Token) String() string {
	return fmt.Sprintf("Token (Kind: %s, Literal: %s)", t.Kind, t.Literal)
}
