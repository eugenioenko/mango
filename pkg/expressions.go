package mango

type Expression interface {
	Accept(visitor VisitorExpression) MangoData
}

type VisitorExpression interface {
	VisitExpressionBinary(expr *ExpressionBinary) MangoData
	VisitExpressionUnary(expr *ExpressionUnary) MangoData
	VisitExpressionPrimary(expr *ExpressionPrimary) MangoData
	VisitExpressionVariable(expr *ExpressionVariable) MangoData
}

type ExpressionBinary struct {
	left     Expression
	operator Token
	right    Expression
}

func NewExpressionBinary(left Expression, operator Token, right Expression) *ExpressionBinary {
	return &ExpressionBinary{left, operator, right}
}

func (expr *ExpressionBinary) Accept(visitor VisitorExpression) MangoData {
	return visitor.VisitExpressionBinary(expr)
}

type ExpressionUnary struct {
	operator Token
	right    Expression
}

func NewExpressionUnary(operator Token, right Expression) *ExpressionUnary {
	return &ExpressionUnary{operator, right}
}

func (expr *ExpressionUnary) Accept(visitor VisitorExpression) MangoData {
	return visitor.VisitExpressionUnary(expr)
}

type ExpressionPrimary struct {
	value Token
}

func NewExpressionPrimary(value Token) *ExpressionPrimary {
	return &ExpressionPrimary{value}
}

func (expr *ExpressionPrimary) Accept(visitor VisitorExpression) MangoData {
	return visitor.VisitExpressionPrimary(expr)
}

type ExpressionVariable struct {
	name Token
}

func NewExpressionVariable(name Token) *ExpressionVariable {
	return &ExpressionVariable{name}
}

func (expr *ExpressionVariable) Accept(visitor VisitorExpression) MangoData {
	return visitor.VisitExpressionVariable(expr)
}
