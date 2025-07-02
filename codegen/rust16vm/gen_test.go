package rust16vm

import (
	"fmt"
	"stag/lexer"
	"stag/primitives"
	"stag/shunting_yard"
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

	ast := shunting_yard.ShuntingYard(tokens)

	exp := "MOV A, #3\nMOV B, #4\nADDR C, A, B\n"

	require.Equal(t, exp, Generate(ast))
}

func TestBinaryOpWithManyNodes(t *testing.T) {
	input := "3 * 4 + 2"
	l := lexer.New(input)

	var tokens []*primitives.Token
	for {
		tok := l.NextToken()
		if tok.Kind == primitives.EOF {
			break
		}
		tokens = append(tokens, tok)
	}

	ast := shunting_yard.ShuntingYard(tokens)

	//exp := "MOV A, #3\nMOV B, #4\nADDR C, A, B\n"

	fmt.Println(Generate(ast))
}
