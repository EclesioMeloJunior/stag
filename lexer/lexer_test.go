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
	t.Run("test_eof_token", func(t *testing.T) {
		input := ""
		l := lexer.New(input)

		tok := l.NextToken()
		expected := &primitives.Token{
			Kind: primitives.EOF,
		}
		require.Equal(t, expected, tok)

		// if I call next token twice or more
		// in a already ended text input
		// it should return me EOF always
		tok = l.NextToken()
		require.Equal(t, expected, tok)
	})

	t.Run("test_more_tokens", func(t *testing.T) {
		//input := "=+-*^/√{}(),;"
		input := "=+ abc   let x = 5 + 5;"

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
			{
				expected: &primitives.Token{
					Kind:    primitives.Ident,
					Literal: "abc",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.Keyword,
					Literal: "let",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.Ident,
					Literal: "x",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.Equals,
					Literal: "=",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.Number,
					Literal: "5",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.Plus,
					Literal: "+",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.Number,
					Literal: "5",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.Semicolon,
					Literal: ";",
				},
			},
		}

		l := lexer.New(input)

		for _, tt := range tests {
			tok := l.NextToken()
			require.Equal(t, tt.expected, tok, "%s != %s", tt.expected.String(), tok.String())
		}
	})

	t.Run("test_if_token", func(t *testing.T) {
		input := "if 5 == 10;"

		tests := []struct {
			expected *lexer.Keyword
			expected *primitives.Token
		}{
			{
				expected: &lexer.Keyword{
					Kind:    lexer.If,
					Literal: "if",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.Number,
					Literal: "5",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.Equals,
					Literal: "==",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.Number,
					Literal: "10",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.Semicolon,
					Literal: ";",
				},
			},
		}

		l := lexer.New(input)

		for _, tt := range tests {
			tok := l.NextToken()
			require.Equal(t, tt.expected, tok, "%s != %s", tt.expected.String(), tok.String())
		}
	})
}
