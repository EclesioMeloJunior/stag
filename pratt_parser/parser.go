package pratt_parser

import (
	"fmt"
	"stag/lexer"
	"stag/pratt_parser/ast"
	"stag/primitives"
	"strconv"
)

const (
	_           int = iota
	LOWEST          // 	-
	EQUALS          //	==
	LESSGREATER     // 	> or <
	SUM             //	+
	PRODUCT         // 	*
	PREFIX          //	-X or !X
	CALL            // myFunction(X)
)

var precedences = map[primitives.TokenKind]int{
	primitives.Equal:    EQUALS,
	primitives.NotEqual: EQUALS,
	primitives.Less:     LESSGREATER,
	primitives.Greater:  LESSGREATER,
	primitives.Plus:     SUM,
	primitives.Minus:    SUM,
	primitives.Slash:    PRODUCT,
	primitives.Star:     PRODUCT,
}

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

	p.prefixParseFns = make(map[primitives.TokenKind]prefixParseFn)
	p.registerPrefix(primitives.Ident, p.parseIdentifier)
	p.registerPrefix(primitives.Number, p.parseIntegerLiteral)
	p.registerPrefix(primitives.Bang, p.parsePrefixExpression)
	p.registerPrefix(primitives.Minus, p.parsePrefixExpression)

	p.infixParseFns = make(map[primitives.TokenKind]infixParseFn)
	p.registerInfix(primitives.Plus, p.parseInfixExpression)
	p.registerInfix(primitives.Minus, p.parseInfixExpression)
	p.registerInfix(primitives.Slash, p.parseInfixExpression)
	p.registerInfix(primitives.Star, p.parseInfixExpression)
	p.registerInfix(primitives.Equal, p.parseInfixExpression)
	p.registerInfix(primitives.NotEqual, p.parseInfixExpression)
	p.registerInfix(primitives.Less, p.parseInfixExpression)
	p.registerInfix(primitives.Greater, p.parseInfixExpression)

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    *p.currentToken,
		Operator: p.currentToken.Literal,
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    *p.currentToken,
		Operator: p.currentToken.Literal,
		Left:     left,
	}
	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}

func (p *Parser) parseIdentifier() ast.Expression {
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
	stmt := &ast.ExpressionStatement{Token: *p.currentToken}

	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(primitives.Semicolon) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) noPrefixParseFnError(t primitives.TokenKind) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.currentToken.Kind]
	if prefix == nil {
		p.noPrefixParseFnError(p.currentToken.Kind)
		return nil
	}
	leftExp := prefix()
	for !p.peekTokenIs(primitives.Semicolon) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Kind]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}
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

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: *p.currentToken}
	value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.currentToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value

	return lit
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Kind]; ok {
		return p
	}
	return LOWEST
}
func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.currentToken.Kind]; ok {
		return p
	}
	return LOWEST
}
