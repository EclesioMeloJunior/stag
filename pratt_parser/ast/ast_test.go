package ast

import (
	"stag/primitives"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: &primitives.Token{Kind: primitives.Keyword, Literal: "let"},
				Name: &Identifier{
					Token: &primitives.Token{Kind: primitives.Ident, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: &primitives.Token{Kind: primitives.Ident, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
