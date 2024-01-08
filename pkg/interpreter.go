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

func Interpret(statements []Expression) (result []MangoData, err error) {
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

func (interpreter *Interpreter) Interpret(statements []Expression) []MangoData {
	result := []MangoData{}
	for _, statement := range statements {
		result = append(result, interpreter.Evaluate(statement))

	}
	return result
}

func (interpreter *Interpreter) Evaluate(expr Expression) MangoData {
	return expr.Accept(interpreter)
}

func (interpreter *Interpreter) Error(errorMessage string) {
	panic("[Runtime Error] " + errorMessage)
}

func (interpreter *Interpreter) VisitExpressionBinary(expr *ExpressionBinary) MangoData {
	left := interpreter.Evaluate(expr.left)
	right := interpreter.Evaluate(expr.right)

	if right.GetType() != left.GetType() {
		interpreter.Error("Type mismatch")
	}

	if left.GetType() == MangoTypeInteger {
		var result int64
		if expr.operator.Literal == "+" {
			result = left.ToInteger() + right.ToInteger()
		} else if expr.operator.Literal == "-" {
			result = left.ToInteger() - right.ToInteger()
		} else if expr.operator.Literal == "*" {
			result = left.ToInteger() * right.ToInteger()
		} else if expr.operator.Literal == "/" {
			result = left.ToInteger() / right.ToInteger()
		}
		return NewMangoInteger(result)
	}

	if left.GetType() == MangoTypeFloat {
		var result float64
		if expr.operator.Literal == "+" {
			result = left.ToFloat() + right.ToFloat()
		} else if expr.operator.Literal == "-" {
			result = left.ToFloat() - right.ToFloat()
		} else if expr.operator.Literal == "*" {
			result = left.ToFloat() * right.ToFloat()
		} else if expr.operator.Literal == "/" {
			result = left.ToFloat() / right.ToFloat()
		}
		return NewMangoFloat(result)
	}

	if left.GetType() == MangoTypeString && expr.operator.Literal == "+" {
		return NewMangoString(left.ToString() + right.ToString())
	}

	return NewMangoNull()
}

func (interpreter *Interpreter) VisitExpressionPrimary(expr *ExpressionPrimary) MangoData {
	if expr.value.Type == TokenTypeNumber {
		value, err := strconv.ParseInt(expr.value.Literal, 10, 64)
		if err != nil {
			interpreter.Error(fmt.Sprintf("%s is not a valid integer", expr.value.Literal))
		}
		return NewMangoInteger(value)
	}

	if expr.value.Type == TokenTypeFloat {
		value, err := strconv.ParseFloat(expr.value.Literal, 64)
		if err != nil {
			interpreter.Error(fmt.Sprintf("%s is not a valid float", expr.value.Literal))
		}
		return NewMangoFloat(value)
	}

	if expr.value.Type == TokenTypeString {
		return NewMangoString(expr.value.Literal)
	}

	interpreter.Error("unknown token type")
	return nil
}

func (interpreter *Interpreter) VisitExpressionUnary(expr *ExpressionUnary) MangoData {
	interpreter.Error("unimplemented")
	return nil
}

func (interpreter *Interpreter) VisitExpressionVariable(expr *ExpressionVariable) MangoData {
	interpreter.Error("unimplemented")
	return nil
}

func (interpreter *Interpreter) VisitExpressionGrouping(expr *ExpressionGrouping) MangoData {
	return interpreter.Evaluate(expr.group)
}

func (interpreter *Interpreter) VisitExpressionAssign(expr *ExpressionAssign) MangoData {
	value := interpreter.Evaluate(expr.value)
	interpreter.Scope.Set(expr.name.Literal, value)
	return nil
}
