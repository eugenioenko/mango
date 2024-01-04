package mango

import (
	"log"
	"os"
)

type Interpreter struct {
	Root    *Scope
	Scope   *Scope
	Runtime *Scope
}

func NewInterpreter() *Interpreter {
	interpreter := Interpreter{}
	interpreter.Runtime = &Scope{Parent: nil, Values: map[string]MangoData{}}
	interpreter.Root = NewScope(nil)
	interpreter.Scope = NewScope(interpreter.Root)
	return &interpreter
}

func (interpreter *Interpreter) Interpret(statements []Expression) (result MangoData) {
	for _, statement := range statements {
		result = interpreter.Evaluate(statement)
	}
	return result
}

func (interpreter *Interpreter) Evaluate(expr Expression) MangoData {
	return expr.Accept(interpreter)
}

func (interpreter *Interpreter) Error(errorMessage string) {
	log.Fatal("[Runtime Error] " + errorMessage)
	os.Exit(1)
}

func (*Interpreter) VisitExpressionBinary(expr *ExpressionBinary) MangoData {
	panic("unimplemented")
}

func (*Interpreter) VisitExpressionPrimary(expr *ExpressionPrimary) MangoData {
	panic("unimplemented")
}

func (*Interpreter) VisitExpressionUnary(expr *ExpressionUnary) MangoData {
	panic("unimplemented")
}

func (*Interpreter) VisitExpressionVariable(expr *ExpressionVariable) MangoData {
	panic("unimplemented")
}
