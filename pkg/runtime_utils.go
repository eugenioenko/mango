package mango

func WF(name string, function Callable) *MangoCallable {
	return &MangoCallable{Kind: MangoTypeCallable, Name: name, Function: function}
}

func EvalParams(interpreter *Interpreter, expressions []Expression) []MangoData {
	params := make([]MangoData, len(expressions))
	for index, expression := range expressions {
		params[index] = interpreter.Evaluate(expression)
	}
	return params
}

func RuntimeDebug(interpreter *Interpreter, expressions []Expression) MangoData {
	return interpreter.Evaluate(expressions[0])
}
