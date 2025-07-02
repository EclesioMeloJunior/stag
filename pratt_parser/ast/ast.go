package ast

import (
	"bytes"
	"stag/primitives"
)

type Node interface {
	TokenLiteral() string
	String() string
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
	Token *primitives.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) StatementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token *primitives.Token
	Value string
}

type ReturnStatement struct {
	Token       primitives.Token
	ReturnValue Expression
}

func (i *Identifier) ExpressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (rs *ReturnStatement) StatementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

type ExpressionStatement struct {
	Token      	primitives.Token
	Expression 	Expression
}

func (es *ExpressionStatement) StatementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
}
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (i *Identifier) String() string {
	return  i.Value
}
