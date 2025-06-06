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
		input := "=+ abc   let x = 5 + 5; [] {} () - * / ^ , < > ! != == <= >="

		tests := []struct {
			expected *primitives.Token
		}{
			{
				expected: &primitives.Token{
					Kind:    primitives.Assign,
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
					Kind:    primitives.Assign,
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

			{
				expected: &primitives.Token{
					Kind:    primitives.OpenBrackets,
					Literal: "[",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.CloseBrackets,
					Literal: "]",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.OpenCurlyBrace,
					Literal: "{",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.CloseCurlyBrace,
					Literal: "}",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.OpenParen,
					Literal: "(",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.CloseParen,
					Literal: ")",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.Minus,
					Literal: "-",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.Star,
					Literal: "*",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.Slash,
					Literal: "/",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.Carrot,
					Literal: "^",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.Comma,
					Literal: ",",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.Less,
					Literal: "<",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.Greater,
					Literal: ">",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.Bang,
					Literal: "!",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.NotEqual,
					Literal: "!=",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.Equal,
					Literal: "==",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.LessOrEqual,
					Literal: "<=",
				},
			},
			{
				expected: &primitives.Token{
					Kind:    primitives.GreaterOrEqual,
					Literal: ">=",
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
