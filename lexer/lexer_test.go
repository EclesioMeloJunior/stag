package lexer_test

import (
	"stag/lexer"
	"stag/primitives"
	"testing"

	"github.com/stretchr/testify/require"
)

// TDD
// Test Driven Development
// Desenvolviment Orientado a Testes
// 1. Escreve o teste da feature que vc esta implementando
// 2. Executa o test, obvio que ele vai falhar
// 3. Implementar a feature
// 4. Garantir que o teste, que antes falhou, agora passe.

func TestNextToken(t *testing.T) {
	//input := "=+-*^/√{}(),;"
	input := "=+"

	tests := []struct {
		expected *primitives.Token
	}{
		{
			expected: &primitives.Token{
				Kind:    primitives.Equals,
				Literal: "=",
			},
		},
		{
			expected: &primitives.Token{
				Kind:    primitives.Plus,
				Literal: "+",
			},
		},
	}

	l := lexer.New(input)

	for _, tt := range tests {
		tok := l.NextToken()
		require.Equal(t, tt.expected, tok)
	}
}
