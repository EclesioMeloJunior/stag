package pratt_parser

import (
	"fmt"
	"go/token"
	"stag/lexer"
	"stag/pratt_parser/ast"
	"stag/primitives"
)

const(
	_ int = iota
	LOWEST    		// 	-
	EQUALS 			//	==
	LESSGREATER		// 	> or <
	PLUS 			//	+
	STAR			// 	*
	PREFIX			//	-X or !X
	CALL			// myFunction(X)
)

type Parser struct {
	l            *lexer.Lexer
	errors       []string
	currentToken *primitives.Token
	peekToken    *primitives.Token

	prefixParseFns map[primitives.TokenKind]prefixParseFn
	infixParseFns  map[primitives.TokenKind]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.prefixParseFns= make(map[primitives.TokenKind]prefixParseFn)
	p.registerPrefix(primitives.Ident, p.parseIdentifier)

	p.nextToken()
	p.nextToken()
	return p
}

func(p *Parser) parseIdentifier() ast.Expression{
	return &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}


func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t primitives.TokenKind) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Kind)
	p.errors = append(p.errors, msg)

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
		if p.currentToken.Literal == "let" {
			return p.parseLetStatement()
		}
		if p.currentToken.Literal == "return" {
			return p.parseReturnStatement()
		}
	default:
		return p.parseExpressionStatement()
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

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: *p.currentToken}

	p.nextToken()
	for !p.currentTokenIs(primitives.Semicolon) {
		p.nextToken()
	}

	return stmt
}


func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt:= &ast.ExpressionStatement{Token: *p.currentToken}

	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(primitives.TokenKind(token.SEMICOLON)){
		p.nextToken()
	}
	return stmt
}
func (p *Parser) parseExpression(precedence int) ast.Expression{
	prefix:= p.prefixParseFns[p.currentToken.Kind]
	if prefix == nil {
		return nil
	}
	leftExp:= prefix()

	return leftExp
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
		p.peekError(t)
		return false
	}
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

func (p *Parser) registerPrefix(tokenType primitives.TokenKind, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
func (p *Parser) registerInfix(tokenType primitives.TokenKind, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
