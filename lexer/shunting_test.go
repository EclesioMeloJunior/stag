package lexer

import (
	"stag/primitives"
	"stag/parse"
	"testing"
)

func TestShuntingYard(t *testing.T) {
	input := "3 + 4 * 2"
	l := New(input)

	var tokens []*primitives.Token
	for {
		tok := l.NextToken()
		if tok.Kind == primitives.EOF {
			break
		}
		tokens = append(tokens, tok)
	}

	rpn := parse.ShuntingYard(tokens)

	expected := []string{"3", "4", "2", "*", "+"}

	if len(rpn) != len(expected) {
		t.Fatalf("Tamanho do resultado esperado: %d, obtido: %d", len(expected), len(rpn))
	}

	for i := range expected {
		if rpn[i].Literal != expected[i] {
			t.Errorf("Esperado token %q no índice %d, obtido %q", expected[i], i, rpn[i].Literal)
		}
	} 
}
