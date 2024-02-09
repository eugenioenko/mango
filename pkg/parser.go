package mango

import (
	"fmt"
)

type Parser struct {
	current    int
	tokens     []Token
	statements []Statement
}

func Parse(tokens []Token) (stms []Statement, err error) {
	defer func() {
		if r := recover(); r != nil {
			stms = nil
			err = fmt.Errorf("%s", r)
			return
		}
	}()

	if len(tokens) == 0 {
		return nil, nil
	}

	parser := NewParser()
	stms = parser.Parse(tokens)
	return stms, err
}

func (parser *Parser) Parse(tokens []Token) []Statement {
	parser.current = 0
	parser.statements = make([]Statement, 0)
	parser.tokens = tokens
	for !parser.Eof() {
		stmt := parser.Statement()
		parser.statements = append(parser.statements, stmt)
	}
	return parser.statements
}

func NewParser() *Parser {
	return &Parser{}
}

func (parser *Parser) MatchType(tokenTypes ...int) bool {
	for _, tokenType := range tokenTypes {
		if parser.Peek().Type == tokenType {
			parser.Advance()
			return true
		}
	}
	return false
}

func (parser *Parser) Match(tokenType int, symbols ...string) bool {
	for _, symbol := range symbols {
		next := parser.Peek()
		if next.Type == tokenType && next.Literal == symbol {
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
func (parser *Parser) Statement() Statement {
	if parser.Match(TokenTypeReserved, "if") {
		return parser.IfStatement()
	}
	if parser.Match(TokenTypeReserved, "while") {
		return parser.WhileStatement()
	}
	if parser.Match(TokenTypeSymbol, "{") {
		return parser.Block()
	}
	if parser.Match(TokenTypeReserved, "print") {
		return parser.Print()
	}
	return parser.ExpressionStatement()
}

func (parser *Parser) WhileStatement() Statement {
	condition := parser.Expression()
	loop := parser.Statement()

	return NewStatementWhile(condition, loop)
}

func (parser *Parser) Block() Statement {
	statements := []Statement{}
	for parser.Peek().Literal != "}" && !parser.Eof() {
		statements = append(statements, parser.Statement())
	}

	parser.ConsumeSymbol("}", "Expected close brace '}' after block statement")
	return NewStatementBlock(statements)
}

func (parser *Parser) IfStatement() Statement {
	condition := parser.Expression()
	thenStmt := parser.Statement()
	var elseStmt Statement
	if parser.Match(TokenTypeReserved, "else") {
		elseStmt = parser.Statement()
	}

	return NewStatementIf(condition, thenStmt, elseStmt)
}

func (parser *Parser) Print() Statement {
	expr := parser.Expression()
	return NewStatementPrint(expr)
}

func (parser *Parser) ExpressionStatement() Statement {
	expr := parser.Expression()
	return NewStatementExpression(expr)
}

func (parser *Parser) Expression() Expression {
	return parser.Assignment()
}

func (parser *Parser) Assignment() Expression {
	expr := parser.Equality()
	if parser.Match(TokenTypeSymbol, ":=") {
		right := parser.Assignment()
		if _, ok := expr.(*ExpressionVariable); ok {
			return NewExpressionAssign(expr.(*ExpressionVariable).Name, right)
		}
	}
	return expr
}

func (parser *Parser) Equality() Expression {
	expr := parser.Addition()
	for parser.Match(TokenTypeSymbol, "==", "!=") {
		operator := parser.Previous()
		right := parser.Addition()
		expr = NewExpressionBinary(expr, operator, right)
	}
	return expr
}

func (parser *Parser) Addition() Expression {
	expr := parser.Multiplication()
	for parser.Match(TokenTypeSymbol, "+", "-") {
		operator := parser.Previous()
		right := parser.Multiplication()
		expr = NewExpressionBinary(expr, operator, right)
	}
	return expr
}

func (parser *Parser) Multiplication() Expression {
	expr := parser.Unary()
	for parser.Match(TokenTypeSymbol, "*", "/") {
		operator := parser.Previous()
		right := parser.Unary()
		expr = NewExpressionBinary(expr, operator, right)
	}
	return expr
}

func (parser *Parser) Unary() Expression {
	if parser.Match(TokenTypeSymbol, "-") {
		operator := parser.Previous()
		right := parser.Unary()
		return NewExpressionUnary(operator, right)
	}
	return parser.Primary()
}

func (parser *Parser) Primary() Expression {

	if parser.MatchType(TokenTypeNumber) {
		return NewExpressionPrimary(parser.Previous())
	}
	if parser.MatchType(TokenTypeIdentifier) {
		return NewExpressionVariable(parser.Previous())
	}
	if parser.MatchType(TokenTypeString) {
		return NewExpressionPrimary(parser.Previous())
	}
	if parser.Match(TokenTypeSymbol, "(") {
		expr := parser.Expression()
		parser.ConsumeSymbol(")", "closing ) required after group expression")
		return NewExpressionGrouping(expr)
	}
	if parser.Eof() {
		parser.Error("Unexpected end of file")
	}
	parser.Error(fmt.Sprintf("Invalid or unexpected token: `%s`", parser.Peek().Literal))

	return nil
}
