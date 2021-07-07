package parser

import (
	"Glox/ast"
	"Glox/token"
	"fmt"
	"strconv"
)

type Parser struct {
	tokens  []token.Token
	current int
	errors  []string
}

func NewParser(tokens []token.Token) *Parser {
	p := &Parser{
		tokens:  tokens,
		current: 0,
		errors:  []string{},
	}
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) Parse() []ast.Statement {
	statements := []ast.Statement{}
	for !p.isAtEnd() {
		statements = append(statements, p.declaration())
	}
	return statements
}

func (p *Parser) statement() ast.Statement {
	if p.match(token.PRINT) {
		return p.printStatement()
	}
	if p.match(token.LEFT_BRACE) {
		return &ast.BlockStmt{Statements: p.Block()}
	}
	return p.expressionStatement()
}

func (p *Parser) Block() []ast.Statement {
	statements := []ast.Statement{}
	for !p.isAtEnd() && !p.check(token.RIGHT_BRACE) {
		statements = append(statements, p.declaration())
	}
	p.consume(token.RIGHT_BRACE, "Expect '}' after block.")
	return statements
}

func (p *Parser) printStatement() ast.Statement {
	value := p.expression()
	p.consume(token.SEMICOLON, "Expect ';' after value.")
	return &ast.PrintStmt{Expression: value}
}

func (p *Parser) varDeclaration() ast.Statement {
	name := p.consume(token.IDENTIFIER, "Expect variable name.")
	var initializer ast.Expression
	if p.match(token.EQUAL) {
		initializer = p.expression()
	}
	p.consume(token.SEMICOLON, "Expect ';' after variable declaration.")

	return &ast.VarStmt{Initializer: initializer, Name: name}
}

func (p *Parser) expressionStatement() ast.Statement {
	expr := p.expression()
	p.consume(token.SEMICOLON, "Expect ';' after expression.")
	return &ast.ExpressionStmt{Expression: expr}
}

func (p *Parser) expression() ast.Expression {
	return p.assignment()
}

func (p *Parser) assignment() ast.Expression {
	expr := p.equality()
	if p.match(token.EQUAL) {
		equals := p.previous()
		value := p.assignment()

		if e, ok := expr.(*ast.Variable); ok {
			name := e.Name
			return &ast.Assign{Name: name, Value: value}
		}
		p.errors = append(p.errors, fmt.Sprintf("%v invalid assignment target.", equals))
	}
	return expr
}

func (p *Parser) declaration() ast.Statement {
	if p.match(token.VAR) {
		return p.varDeclaration()
	}
	return p.statement()
}

func (p *Parser) equality() ast.Expression {
	expr := p.comparison()

	for p.match(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = &ast.Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr
}

func (p *Parser) comparison() ast.Expression {
	expr := p.term()
	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = &ast.Binary{Left: expr, Operator: operator, Right: right}
	}
	return expr
}

func (p *Parser) term() ast.Expression {
	expr := p.factor()

	for p.match(token.PLUS, token.MINUS) {
		operator := p.previous()
		right := p.factor()
		expr = &ast.Binary{Left: expr, Operator: operator, Right: right}
	}
	return expr
}

func (p *Parser) factor() ast.Expression {
	expr := p.unary()

	for p.match(token.STAR, token.SLASH) {
		operator := p.previous()
		right := p.unary()
		expr = &ast.Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr
}

func (p *Parser) unary() ast.Expression {
	if p.match(token.MINUS, token.BANG) {
		operator := p.previous()
		right := p.unary()
		return &ast.Unary{Operator: operator, Right: right}
	}
	return p.primary()
}

func (p *Parser) primary() ast.Expression {
	if p.match(token.FALSE) {
		return &ast.Literal{Value: false}
	}
	if p.match(token.TRUE) {
		return &ast.Literal{Value: true}
	}
	if p.match(token.NIL) {
		return &ast.Literal{Value: nil}
	}
	if p.match(token.NUMBER, token.STRING) {
		tok := p.previous()
		if tok.Type == token.NUMBER {
			str := tok.Lexeme
			val, err := strconv.ParseFloat(str, 64)
			if err != nil {
				p.errors = append(p.errors, "could not parse token to float")
			}
			return &ast.Literal{Value: val}
		}
		return &ast.Literal{Value: tok.Lexeme}
	}
	if p.match(token.IDENTIFIER) {
		return &ast.Variable{Name: p.previous()}
	}
	if p.match(token.LEFT_PAREN) {
		expr := p.expression()
		p.consume(token.RIGHT_PAREN, "Expect ')' after expression.")
		return &ast.Grouping{Expression: expr}
	}
	p.errors = append(p.errors, "Expect expression.")
	return nil
}

func (p *Parser) consume(t token.TokenType, msg string) token.Token {
	if p.check(t) {
		return p.advance()
	}
	p.errors = append(p.errors, msg)

	return token.Token{}
}

func (p *Parser) match(types ...token.TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(t token.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == t
}

func (p *Parser) advance() token.Token {
	if !p.isAtEnd() {
		p.current += 1
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF
}

func (p *Parser) peek() token.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() token.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().Type == token.SEMICOLON {
			return
		}

		switch p.peek().Type {
		case token.CLASS:
		case token.FUN:
		case token.VAR:
		case token.FOR:
		case token.IF:
		case token.WHILE:
		case token.PRINT:
		case token.RETURN:
			return
		}
		p.advance()
	}
}
