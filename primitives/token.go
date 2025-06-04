package primitives

type TokenKind uint8

const (
	Let TokenKind = iota
	Ident

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

	EOF
)

type Token struct {
	Kind         TokenKind
	Literal      string
	SourceColumn int
	SourceLine   int
}
