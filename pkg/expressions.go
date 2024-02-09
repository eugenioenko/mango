package mango

type Expression interface {
    Accept(visitor VisitorExpression) MangoData
}

type VisitorExpression interface {
	VisitExpressionAssign(expr *ExpressionAssign) MangoData
	VisitExpressionBinary(expr *ExpressionBinary) MangoData
	VisitExpressionEquality(expr *ExpressionEquality) MangoData
	VisitExpressionGrouping(expr *ExpressionGrouping) MangoData
	VisitExpressionUnary(expr *ExpressionUnary) MangoData
	VisitExpressionPrimary(expr *ExpressionPrimary) MangoData
	VisitExpressionVariable(expr *ExpressionVariable) MangoData
}

type ExpressionAssign struct {
    Name Token
    Value Expression
}

func NewExpressionAssign(Name Token, Value Expression) *ExpressionAssign {
	return &ExpressionAssign{Name, Value}
}

func (expr *ExpressionAssign) Accept (visitor VisitorExpression) MangoData {
	return visitor.VisitExpressionAssign(expr)
}

type ExpressionBinary struct {
    Left Expression
    Operator Token
    Right Expression
}

func NewExpressionBinary(Left Expression, Operator Token, Right Expression) *ExpressionBinary {
	return &ExpressionBinary{Left, Operator, Right}
}

func (expr *ExpressionBinary) Accept (visitor VisitorExpression) MangoData {
	return visitor.VisitExpressionBinary(expr)
}

type ExpressionEquality struct {
    Left Expression
    Right Expression
}

func NewExpressionEquality(Left Expression, Right Expression) *ExpressionEquality {
	return &ExpressionEquality{Left, Right}
}

func (expr *ExpressionEquality) Accept (visitor VisitorExpression) MangoData {
	return visitor.VisitExpressionEquality(expr)
}

type ExpressionGrouping struct {
    Group Expression
}

func NewExpressionGrouping(Group Expression) *ExpressionGrouping {
	return &ExpressionGrouping{Group}
}

func (expr *ExpressionGrouping) Accept (visitor VisitorExpression) MangoData {
	return visitor.VisitExpressionGrouping(expr)
}

type ExpressionUnary struct {
    Operator Token
    Right Expression
}

func NewExpressionUnary(Operator Token, Right Expression) *ExpressionUnary {
	return &ExpressionUnary{Operator, Right}
}

func (expr *ExpressionUnary) Accept (visitor VisitorExpression) MangoData {
	return visitor.VisitExpressionUnary(expr)
}

type ExpressionPrimary struct {
    Value Token
}

func NewExpressionPrimary(Value Token) *ExpressionPrimary {
	return &ExpressionPrimary{Value}
}

func (expr *ExpressionPrimary) Accept (visitor VisitorExpression) MangoData {
	return visitor.VisitExpressionPrimary(expr)
}

type ExpressionVariable struct {
    Name Token
}

func NewExpressionVariable(Name Token) *ExpressionVariable {
	return &ExpressionVariable{Name}
}

func (expr *ExpressionVariable) Accept (visitor VisitorExpression) MangoData {
	return visitor.VisitExpressionVariable(expr)
}

