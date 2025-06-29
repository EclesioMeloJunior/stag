package ast

import (
	"stag/primitives"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	StatementNode()
}

type Expression interface {
	Node
	ExpressionNode()
}
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type LetStatement struct {
	Token      *primitives.Token
	Name       *Identifier
	Value      Expression
}

func (ls *LetStatement) StatementNode()  {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}


type Identifier struct {
	Token      *primitives.Token
	Value     string
}

func (i *Identifier) ExpressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

