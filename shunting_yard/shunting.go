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
			outputQueue = append(outputQueue, Number{Value: int64(n)})
		
		case primitives.OpenParen:
			operatorStack = append(operatorStack, token)

		case primitives.CloseParen:
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1].Kind != primitives.OpenParen {
				top := operatorStack[len(operatorStack)-1]
				operatorStack = operatorStack[:len(operatorStack)-1]

				rhs := outputQueue[len(outputQueue)-1]
				lhs := outputQueue[len(outputQueue)-2]
				outputQueue = outputQueue[:len(outputQueue) -2]

				switch top.Literal {
				case "+":
					outputQueue = append(outputQueue, BinaryOperation{rhs: rhs.(Expression), lhs: lhs.(Expression), op: add} )
				case "-":
					outputQueue = append(outputQueue, BinaryOperation{rhs: rhs.(Expression), lhs: lhs.(Expression), op: sub} )
				case "*":
					outputQueue = append(outputQueue, BinaryOperation{rhs: rhs.(Expression), lhs: lhs.(Expression), op: mul} )
				case "/": 
					outputQueue = append(outputQueue, BinaryOperation{rhs: rhs.(Expression), lhs: lhs.(Expression), op: div} )
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
						outputQueue = outputQueue[:len(outputQueue) -2]
						switch top.Literal {
						case "+":
							outputQueue = append(outputQueue, BinaryOperation{rhs: rhs.(Expression), lhs: lhs.(Expression), op: add} )
						case "-":
							outputQueue = append(outputQueue, BinaryOperation{rhs: rhs.(Expression), lhs: lhs.(Expression), op: sub} )
						case "*":
							outputQueue = append(outputQueue, BinaryOperation{rhs: rhs.(Expression), lhs: lhs.(Expression), op: mul} )
						case "/": 
							outputQueue = append(outputQueue, BinaryOperation{rhs: rhs.(Expression), lhs: lhs.(Expression), op: div} )
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
		outputQueue = outputQueue[:len(outputQueue) -2]
	
		switch op.Literal {
		case "+":
			outputQueue = append(outputQueue, BinaryOperation{rhs: rhs.(Expression), lhs: lhs.(Expression), op: add} )
		case "-":
			outputQueue = append(outputQueue, BinaryOperation{rhs: rhs.(Expression), lhs: lhs.(Expression), op: sub} )
		case "*":
			outputQueue = append(outputQueue, BinaryOperation{rhs: rhs.(Expression), lhs: lhs.(Expression), op: mul} )
		case "/": 
			outputQueue = append(outputQueue, BinaryOperation{rhs: rhs.(Expression), lhs: lhs.(Expression), op: div} )
		}
		operatorStack = operatorStack[:len(operatorStack)-1]
	}

	return outputQueue
}

func extractTokensBetween(tokens []*primitives.Token, open string, close string) []*primitives.Token {
	depth := 0
	start := -1
	for i, token := range tokens {
		if token.Literal == open {
			if depth == 0 {
				start = i + 1
			}
			depth++
		} else if token.Literal == close {
			depth--
			if depth == 0 {
				return tokens[start:i]
			}
		}
	}
	return nil
}

func hasElse(tokens []*primitives.Token) bool {
	for i, token := range tokens {
		if token.Literal == "else" {
			return true
		}
		// Assume que '}' fecha o bloco e depois disso pode vir o else
		if token.Literal == "}" && i+1 < len(tokens) && tokens[i+1].Literal == "else" {
			return true
		}
	}
	return false
}

func extractTokensAfterElse(tokens []*primitives.Token) []*primitives.Token {
	for i, token := range tokens {
		if token.Literal == "else" && i+1 < len(tokens) && tokens[i+1].Literal == "{" {
			return extractTokensBetween(tokens[i:], "{", "}")
		}
	}
	return nil
}

func splitByComma(tokens []*primitives.Token) [][]*primitives.Token {
	var args [][]*primitives.Token
	start := 0
	depth := 0
	for i, tok := range tokens {
		if tok.Literal == "(" {
			depth++
		} else if tok.Literal == ")" {
			depth--
		} else if tok.Literal == "," && depth == 0 {
			args = append(args, tokens[start:i])
			start = i + 1
		}
	}
	args = append(args, tokens[start:])
	return args
}
