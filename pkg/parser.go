package mango

import (
	"log"
	"os"
)

type Parser struct {
	current     int
	tokens      []Token
	expressions []Expression
}

func Parse(tokens []Token) ([]Expression, error) {
	parser := NewParser()
	expressions := parser.Parse(tokens)
	return expressions, nil
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

func (parser *Parser) Match(tokenTypes ...TokenType) bool {
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

func (parser *Parser) Check(tokenTypes ...TokenType) bool {
	for _, tokenType := range tokenTypes {
		currentType := parser.Peek().Type
		if currentType == tokenType {
			return true
		}
	}
	return false
}

func (parser *Parser) Consume(errorMessage string, tokenTypes ...TokenType) Token {
	if parser.Check(tokenTypes...) {
		return parser.Advance()
	}
	parser.Error(parser.Peek(), errorMessage)
	return parser.Peek()
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

func (parser *Parser) Error(token Token, errorMessage string) {
	log.Fatal("[Syntax Error] " + errorMessage)
	os.Exit(1)
}

// ------------------------------------------------------------------------------
// AST STARTS HERE
// ------------------------------------------------------------------------------
func (parser *Parser) Statement() Expression {
	return parser.Expression()
}

func (parser *Parser) Expression() Expression {
	return parser.Addition()
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

	if parser.Match(TokenTypeNumber) {
		return NewExpressionPrimary(parser.Previous())
	}
	if parser.Match(TokenTypeIdentifier) {
		return NewExpressionVariable(parser.Previous())
	}

	parser.Error(parser.Previous(), "Unexpected end of file")
	return nil
}
