package parse

import (
	"stag/primitives"
)

var precedencia = map[primitives.TokenKind]int{
	primitives.Equal:          1,
	primitives.NotEqual:       1,
	primitives.Less:           2,
	primitives.LessOrEqual:    2,
	primitives.Greater:        2,
	primitives.GreaterOrEqual: 2,
	primitives.Plus:           3,
	primitives.Minus:          3,
	primitives.Star:           4,
	primitives.Slash:          4,
	primitives.Carrot:         5, 
}

func isOperator(kind primitives.TokenKind) bool {
	_, ok := precedencia[kind]
	return ok
}

func ShuntingYard(tokens []*primitives.Token) []*primitives.Token {
	var outputQueue []*primitives.Token
	var operatorStack []*primitives.Token

	for _, token := range tokens {
		switch token.Kind {
		case primitives.Number, primitives.Ident:
			outputQueue = append(outputQueue, token)

		case primitives.OpenParen:
			operatorStack = append(operatorStack, token)

		case primitives.CloseParen:
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1].Kind != primitives.OpenParen {
				outputQueue = append(outputQueue, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			if len(operatorStack) == 0 {
				panic("Parêntese desbalanceado")
			}
			operatorStack = operatorStack[:len(operatorStack)-1] 

		default:
			if isOperator(token.Kind) {
				for len(operatorStack) > 0 {
					top := operatorStack[len(operatorStack)-1]
					if isOperator(top.Kind) &&
						precedencia[top.Kind] >= precedencia[token.Kind] {
						outputQueue = append(outputQueue, top)
						operatorStack = operatorStack[:len(operatorStack)-1]
					} else {
						break
					}
				}
				operatorStack = append(operatorStack, token)
			}
		}
	}

	for len(operatorStack) > 0 {
		op := operatorStack[len(operatorStack)-1]
		if op.Kind == primitives.OpenParen {
			panic("Parêntese desbalanceado")
		}
		outputQueue = append(outputQueue, op)
		operatorStack = operatorStack[:len(operatorStack)-1]
	}

	return outputQueue
}
