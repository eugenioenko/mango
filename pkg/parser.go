package mango

import (
	"fmt"
)

type Parser struct {
	current     int
	tokens      []Token
	expressions []Expression
}

func Parse(tokens []Token) (exprs []Expression, err error) {
	defer func() {
		if r := recover(); r != nil {
			exprs = nil
			err = fmt.Errorf("%s", r)
			return
		}
	}()

	if len(tokens) == 0 {
		return nil, nil
	}

	parser := NewParser()
	exprs = parser.Parse(tokens)
	return exprs, err
}

func (parser *Parser) Parse(tokens []Token) []Expression {
	parser.current = 0
	parser.expressions = make([]Expression, 0)
	parser.tokens = tokens
	for !parser.Eof() {
		stmt := parser.Statement()
		parser.expressions = append(parser.expressions, stmt)
	}
	return parser.expressions
}

func NewParser() *Parser {
	return &Parser{}
}

func (parser *Parser) MatchToken(tokenTypes ...int) bool {
	for _, tokenType := range tokenTypes {
		if parser.Peek().Type == tokenType {
			parser.Advance()
			return true
		}
	}
	return false
}

func (parser *Parser) MatchSymbol(symbols ...string) bool {
	for _, symbol := range symbols {
		next := parser.Peek()
		if next.Type == TokenTypeSymbol && next.Literal == symbol {
			parser.Advance()
			return true
		}
	}
	return false
}

func (parser *Parser) Check(tokenType int) bool {
	return parser.Peek().Type == tokenType
}

func (parser *Parser) ConsumeToken(tokenType int, errorMessage string) Token {
	if parser.Check(tokenType) {
		return parser.Advance()
	}
	parser.Error(errorMessage)
	return parser.Peek()
}

func (parser *Parser) ConsumeSymbol(symbol string, errorMessage string) Token {
	next := parser.Peek()
	if next.Literal == symbol {
		return parser.Advance()
	}
	parser.Error(errorMessage)
	return next
}

func (parser *Parser) Advance() Token {
	if !parser.Eof() {
		parser.current += 1
	}
	return parser.Previous()
}

func (parser *Parser) Previous() Token {
	return parser.tokens[parser.current-1]
}

func (parser *Parser) Peek() Token {
	return parser.tokens[parser.current]
}

func (parser *Parser) Eof() bool {
	return parser.Peek().Type == TokenTypeEof ||
		parser.current >= len(parser.tokens)
}

func (parser *Parser) Error(errorMessage string) {
	panic("[Syntax Error] " + errorMessage)
}

// ------------------------------------------------------------------------------
// AST STARTS HERE
// ------------------------------------------------------------------------------
func (parser *Parser) Statement() Expression {
	return parser.Expression()
}

func (parser *Parser) Expression() Expression {
	return parser.Assignment()
}

func (parser *Parser) Assignment() Expression {
	expr := parser.Addition()
	if parser.MatchSymbol(":=") {
		right := parser.Assignment()
		if _, ok := expr.(*ExpressionVariable); ok {
			return NewExpressionAssign(expr.(*ExpressionVariable).name, right)
		}
	}
	return expr
}

func (parser *Parser) Addition() Expression {
	expr := parser.Multiplication()
	for parser.MatchSymbol("+", "-") {
		operator := parser.Previous()
		right := parser.Multiplication()
		expr = NewExpressionBinary(expr, operator, right)
	}
	return expr
}

func (parser *Parser) Multiplication() Expression {
	expr := parser.Unary()
	for parser.MatchSymbol("*", "/") {
		operator := parser.Previous()
		right := parser.Unary()
		expr = NewExpressionBinary(expr, operator, right)
	}
	return expr
}

func (parser *Parser) Unary() Expression {
	if parser.MatchSymbol("-") {
		operator := parser.Previous()
		right := parser.Unary()
		return NewExpressionUnary(operator, right)
	}
	return parser.Primary()
}

func (parser *Parser) Primary() Expression {

	if parser.MatchToken(TokenTypeNumber) {
		return NewExpressionPrimary(parser.Previous())
	}
	if parser.MatchToken(TokenTypeIdentifier) {
		return NewExpressionVariable(parser.Previous())
	}
	if parser.MatchSymbol("(") {
		expr := parser.Expression()
		parser.ConsumeSymbol(")", "closing ) required after group expression")
		return NewExpressionGrouping(expr)
	}
	if parser.Peek().Type == TokenTypeReserved && parser.Peek().Literal == "print" {
		parser.ConsumeToken(TokenTypeReserved, "print token")
		expr := parser.Expression()
		return NewExpressionPrint(expr)
	}

	if parser.Eof() {
		parser.Error("Unexpected end of file")
	}
	parser.Error(fmt.Sprintf("Invalid or unexpected token: `%s`", parser.Peek().Literal))

	return nil
}
