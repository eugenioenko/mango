package mango

import (
	"fmt"
	"strconv"
)

type Interpreter struct {
	Scope *Scope
}

func NewInterpreter() *Interpreter {
	interpreter := Interpreter{}
	interpreter.Scope = NewScope(nil)
	return &interpreter
}

func Interpret(statements []Statement) (result []MangoData, err error) {
	defer func() {
		if r := recover(); r != nil {
			result = nil
			err = fmt.Errorf("%s", r)
			return
		}
	}()

	if len(statements) == 0 {
		return nil, nil
	}

	interpreter := NewInterpreter()
	result = interpreter.Interpret(statements)
	return result, err
}

func (interpreter *Interpreter) Interpret(statements []Statement) []MangoData {
	result := []MangoData{}
	for _, statement := range statements {
		result = append(result, interpreter.Execute(statement))

	}
	return result
}

func (interpreter *Interpreter) Evaluate(expr Expression) MangoData {
	return expr.Accept(interpreter)
}

func (interpreter *Interpreter) Execute(stmt Statement) MangoData {
	return stmt.Accept(interpreter)
}

func (interpreter *Interpreter) Error(errorMessage string) {
	panic("[Runtime Error] " + errorMessage)
}

func (interpreter *Interpreter) VisitStatementExpression(stmt *StatementExpression) MangoData {
	return interpreter.Evaluate(stmt.Value)
}

func (interpreter *Interpreter) VisitStatementPrint(stmt *StatementPrint) MangoData {
	value := interpreter.Evaluate(stmt.Value)
	fmt.Println(value.GetValue())
	return value
}

func (interpreter *Interpreter) VisitStatementBlock(stmt *StatementBlock) MangoData {
	interpreter.Interpret(stmt.Statements)
	return NewMangoNull()
}

func (interpreter *Interpreter) VisitStatementIf(stmt *StatementIf) MangoData {
	condition := interpreter.Evaluate(stmt.Condition)
	if condition.ToBoolean() {
		return interpreter.Execute(stmt.Then)
	} else if stmt.Else != nil {
		return interpreter.Execute(stmt.Else)
	}
	return NewMangoNull()
}

func (interpreter *Interpreter) VisitExpressionBinary(expr *ExpressionBinary) MangoData {
	left := interpreter.Evaluate(expr.Left)
	right := interpreter.Evaluate(expr.Right)
	operator := expr.Operator.Literal

	if operator == "==" || operator == "!=" {
		equals := left.GetType() == right.GetType() && left.GetValue() == right.GetValue()
		if operator == "!=" {
			equals = !equals
		}
		return NewMangoBoolean(equals)
	}

	if right.GetType() != left.GetType() {
		interpreter.Error("Type mismatch")
	}

	if left.GetType() == MangoTypeInteger {
		var result int64
		if operator == "+" {
			result = left.ToInteger() + right.ToInteger()
		} else if operator == "-" {
			result = left.ToInteger() - right.ToInteger()
		} else if operator == "*" {
			result = left.ToInteger() * right.ToInteger()
		} else if operator == "/" {
			result = left.ToInteger() / right.ToInteger()
		}
		return NewMangoInteger(result)
	}

	if left.GetType() == MangoTypeFloat {
		var result float64
		if operator == "+" {
			result = left.ToFloat() + right.ToFloat()
		} else if operator == "-" {
			result = left.ToFloat() - right.ToFloat()
		} else if operator == "*" {
			result = left.ToFloat() * right.ToFloat()
		} else if operator == "/" {
			result = left.ToFloat() / right.ToFloat()
		}
		return NewMangoFloat(result)
	}

	if left.GetType() == MangoTypeString && operator == "+" {
		return NewMangoString(left.ToString() + right.ToString())
	}

	return NewMangoNull()
}

func (interpreter *Interpreter) VisitExpressionPrimary(expr *ExpressionPrimary) MangoData {
	if expr.Value.Type == TokenTypeNumber {
		value, err := strconv.ParseInt(expr.Value.Literal, 10, 64)
		if err != nil {
			interpreter.Error(fmt.Sprintf("%s is not a valid integer", expr.Value.Literal))
		}
		return NewMangoInteger(value)
	}

	if expr.Value.Type == TokenTypeFloat {
		value, err := strconv.ParseFloat(expr.Value.Literal, 64)
		if err != nil {
			interpreter.Error(fmt.Sprintf("%s is not a valid float", expr.Value.Literal))
		}
		return NewMangoFloat(value)
	}

	if expr.Value.Type == TokenTypeString {
		return NewMangoString(expr.Value.Literal)
	}

	interpreter.Error("unknown token type")
	return nil
}

func (interpreter *Interpreter) VisitExpressionUnary(expr *ExpressionUnary) MangoData {
	interpreter.Error("unimplemented")
	return nil
}

func (interpreter *Interpreter) VisitExpressionVariable(expr *ExpressionVariable) MangoData {
	value, ok := interpreter.Scope.Get(expr.Name.Literal)
	if !ok {
		return NewMangoNull()
	}
	return value
}

func (interpreter *Interpreter) VisitExpressionGrouping(expr *ExpressionGrouping) MangoData {
	return interpreter.Evaluate(expr.Group)
}

func (interpreter *Interpreter) VisitExpressionPrint(expr *StatementPrint) MangoData {
	result := interpreter.Evaluate(expr.Value)
	fmt.Println(result.ToString())
	return result
}

func (interpreter *Interpreter) VisitExpressionAssign(expr *ExpressionAssign) MangoData {
	value := interpreter.Evaluate(expr.Value)
	interpreter.Scope.Set(expr.Name.Literal, value)
	return nil
}
