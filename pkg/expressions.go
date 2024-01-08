package mango

type Expression interface {
	Accept(visitor VisitorExpression) MangoData
}

type VisitorExpression interface {
	VisitExpressionAssign(expr *ExpressionAssign) MangoData
	VisitExpressionBinary(expr *ExpressionBinary) MangoData
	VisitExpressionGrouping(expr *ExpressionGrouping) MangoData
	VisitExpressionUnary(expr *ExpressionUnary) MangoData
	VisitExpressionPrimary(expr *ExpressionPrimary) MangoData
	VisitExpressionPrint(expr *ExpressionPrint) MangoData
	VisitExpressionVariable(expr *ExpressionVariable) MangoData
}

type ExpressionAssign struct {
	name  Token
	value Expression
}

func NewExpressionAssign(name Token, value Expression) *ExpressionAssign {
	return &ExpressionAssign{name, value}
}

func (expr *ExpressionAssign) Accept(visitor VisitorExpression) MangoData {
	return visitor.VisitExpressionAssign(expr)
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

type ExpressionGrouping struct {
	group Expression
}

func NewExpressionGrouping(group Expression) *ExpressionGrouping {
	return &ExpressionGrouping{group}
}

func (expr *ExpressionGrouping) Accept(visitor VisitorExpression) MangoData {
	return visitor.VisitExpressionGrouping(expr)
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

type ExpressionPrint struct {
	value Expression
}

func NewExpressionPrint(value Expression) *ExpressionPrint {
	return &ExpressionPrint{value}
}

func (expr *ExpressionPrint) Accept(visitor VisitorExpression) MangoData {
	return visitor.VisitExpressionPrint(expr)
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
