package pratt_parser

import (
	"stag/lexer"
	"stag/pratt_parser/ast"
	"stag/primitives"
)

type Parser struct {
	l            *lexer.Lexer
	currentToken *primitives.Token
	peekToken    *primitives.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParserProgram() *ast.Program {
	return nil
}


func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.currentToken.Kind != primitives.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}


func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Kind {
	case primitives.Keyword:
		if p.currentToken.Literal == "let"{
			return p.parseLetStatement()
		}
		
	default:
		return nil
	}
	return nil
}


func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currentToken}
	if !p.expectPeek(primitives.Ident) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectPeek(primitives.Assign) {
		return nil
	}
	
	for !p.currentTokenIs(primitives.Semicolon) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) currentTokenIs(t primitives.TokenKind) bool {
	return p.currentToken.Kind == t
}

func (p *Parser) peekTokenIs(t primitives.TokenKind) bool {
	return p.peekToken.Kind == t
}

func (p *Parser) expectPeek(t primitives.TokenKind) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

