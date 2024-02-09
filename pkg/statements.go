package mango

type Statement interface {
    Accept(visitor VisitorStatement) MangoData
}

type VisitorStatement interface {
	VisitStatementExpression(stmt *StatementExpression) MangoData
	VisitStatementPrint(stmt *StatementPrint) MangoData
	VisitStatementBlock(stmt *StatementBlock) MangoData
}

type StatementExpression struct {
    Value Expression
}

func NewStatementExpression(Value Expression) *StatementExpression {
	return &StatementExpression{Value}
}

func (stmt *StatementExpression) Accept (visitor VisitorStatement) MangoData {
	return visitor.VisitStatementExpression(stmt)
}

type StatementPrint struct {
    Value Expression
}

func NewStatementPrint(Value Expression) *StatementPrint {
	return &StatementPrint{Value}
}

func (stmt *StatementPrint) Accept (visitor VisitorStatement) MangoData {
	return visitor.VisitStatementPrint(stmt)
}

type StatementBlock struct {
    Statements []Statement
}

func NewStatementBlock(Statements []Statement) *StatementBlock {
	return &StatementBlock{Statements}
}

func (stmt *StatementBlock) Accept (visitor VisitorStatement) MangoData {
	return visitor.VisitStatementBlock(stmt)
}

