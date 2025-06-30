package ast

import(
	"stag/primitives"
	"testing"
)

func TestString(t *testing.T){
	program:= &Program{
		Statements: []Statement{
			&LetStatement{
				Token: primitives.Token{Type: primitives.Keyword, Literal: "let"},
				Name: &Identifier{
					Token: primitives.Token{Type: primitives.Ident, Literal: "myVar"},
					Value: "myVar",
				},
			},
		},
	}
}