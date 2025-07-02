package shunting_yard

import (
	"stag/lexer"
	"stag/primitives"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimple(t *testing.T) {
	input := "3 + 4"
	l := lexer.New(input)

	var tokens []*primitives.Token
	for {
		tok := l.NextToken()
		if tok.Kind == primitives.EOF {
			break
		}
		tokens = append(tokens, tok)
	}

	rpn := ShuntingYard(tokens)

	expected := []Statement{
		&BinaryOperation{Op: Add, Lhs: &Number{Value: 3}, Rhs: &Number{Value: 4}},
	}
	require.Equal(t, expected, rpn)
}

func TestPrecendence(t *testing.T) {
	input := "3 + 4 * 2"
	l := lexer.New(input)

	var tokens []*primitives.Token
	for {
		tok := l.NextToken()
		if tok.Kind == primitives.EOF {
			break
		}
		tokens = append(tokens, tok)
	}

	rpn := ShuntingYard(tokens)

	expected := []Statement{
		&BinaryOperation{Op: Add, Lhs: &Number{Value: 3},
			Rhs: &BinaryOperation{Op: Mul, Lhs: &Number{Value: 4}, Rhs: &Number{Value: 2}}},
	}
	require.Equal(t, expected, rpn)
}

func TestOpenCloseParen(t *testing.T) {
	input := "(3 + 4 * 2) * 2"
	l := lexer.New(input)

	var tokens []*primitives.Token
	for {
		tok := l.NextToken()
		if tok.Kind == primitives.EOF {
			break
		}
		tokens = append(tokens, tok)
	}

	rpn := ShuntingYard(tokens)

	expected := []Statement{
		&BinaryOperation{Op: Mul, Lhs: &BinaryOperation{Op: Add, Lhs: &Number{Value: 3}, Rhs: &BinaryOperation{Op: Mul, Lhs: &Number{Value: 4}, Rhs: &Number{Value: 2}}}, Rhs: &Number{Value: 2}},
	}
	require.Equal(t, expected, rpn)
}
