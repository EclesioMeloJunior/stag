package shunting_yard

import (
	"stag/primitives"
	"strconv"
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

func ShuntingYard(tokens []*primitives.Token) []Statement {
	var outputQueue []Statement
	var operatorStack []*primitives.Token

	for _, token := range tokens {
		switch token.Kind {
		// case primitives.Ident:
		// outputQueue = append(outputQueue, token)

		case primitives.Number:
			n, err := strconv.Atoi(token.Literal)
			if err != nil {
				panic(err)
			}
			outputQueue = append(outputQueue, &Number{Value: int64(n)})

		case primitives.OpenParen:
			operatorStack = append(operatorStack, token)

		case primitives.CloseParen:
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1].Kind != primitives.OpenParen {
				top := operatorStack[len(operatorStack)-1]
				operatorStack = operatorStack[:len(operatorStack)-1]

				rhs := outputQueue[len(outputQueue)-1]
				lhs := outputQueue[len(outputQueue)-2]
				outputQueue = outputQueue[:len(outputQueue)-2]

				switch top.Literal {
				case "+":
					outputQueue = append(outputQueue, &BinaryOperation{Rhs: rhs.(Expression), Lhs: lhs.(Expression), Op: Add})
				case "-":
					outputQueue = append(outputQueue, &BinaryOperation{Rhs: rhs.(Expression), Lhs: lhs.(Expression), Op: Sub})
				case "*":
					outputQueue = append(outputQueue, &BinaryOperation{Rhs: rhs.(Expression), Lhs: lhs.(Expression), Op: Mul})
				case "/":
					outputQueue = append(outputQueue, &BinaryOperation{Rhs: rhs.(Expression), Lhs: lhs.(Expression), Op: Div})
				}
			}

			if len(operatorStack) == 0 {
				panic("Parêntese desbalanceado")
			}

			operatorStack = operatorStack[:len(operatorStack)-1]

		default:
			if isOperator(token.Kind) {
				for len(operatorStack) > 0 {
					top := operatorStack[len(operatorStack)-1]
					if isOperator(top.Kind) && precedencia[top.Kind] >= precedencia[token.Kind] {
						rhs := outputQueue[len(outputQueue)-1]
						lhs := outputQueue[len(outputQueue)-2]
						outputQueue = outputQueue[:len(outputQueue)-2]
						switch top.Literal {
						case "+":
							outputQueue = append(outputQueue, &BinaryOperation{Rhs: rhs.(Expression), Lhs: lhs.(Expression), Op: Add})
						case "-":
							outputQueue = append(outputQueue, &BinaryOperation{Rhs: rhs.(Expression), Lhs: lhs.(Expression), Op: Sub})
						case "*":
							outputQueue = append(outputQueue, &BinaryOperation{Rhs: rhs.(Expression), Lhs: lhs.(Expression), Op: Mul})
						case "/":
							outputQueue = append(outputQueue, &BinaryOperation{Rhs: rhs.(Expression), Lhs: lhs.(Expression), Op: Div})
						}

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
		rhs := outputQueue[len(outputQueue)-1]
		lhs := outputQueue[len(outputQueue)-2]
		outputQueue = outputQueue[:len(outputQueue)-2]

		switch op.Literal {
		case "+":
			outputQueue = append(outputQueue, &BinaryOperation{Rhs: rhs.(Expression), Lhs: lhs.(Expression), Op: Add})
		case "-":
			outputQueue = append(outputQueue, &BinaryOperation{Rhs: rhs.(Expression), Lhs: lhs.(Expression), Op: Sub})
		case "*":
			outputQueue = append(outputQueue, &BinaryOperation{Rhs: rhs.(Expression), Lhs: lhs.(Expression), Op: Mul})
		case "/":
			outputQueue = append(outputQueue, &BinaryOperation{Rhs: rhs.(Expression), Lhs: lhs.(Expression), Op: Div})
		}
		operatorStack = operatorStack[:len(operatorStack)-1]
	}

	return outputQueue
}
