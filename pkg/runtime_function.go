package mango

/*

func RuntimeFunc(interpreter *Interpreter, expressions []Expression) MangoData {
	name := expressions[0].(*ExpressionAtom).Atom.Literal
	args := make([]string, len(expressions[1].(*ExpressionList).List))
	for index, token := range expressions[1].(*ExpressionList).List {
		args[index] = token.(*ExpressionAtom).Atom.Literal
	}
	function := NewMangoFunction(name, args, expressions[2:])
	interpreter.Scope.Set(name, function)
	return function
}

func RuntimeReturnFrom(interpreter *Interpreter, expressions []Expression) MangoData {
	result := interpreter.Evaluate(expressions[1])
	from := expressions[0].(*ExpressionAtom).Atom.Literal

	panic(NewMangoReturn(from, result))
}

func RuntimeReturn(interpreter *Interpreter, expressions []Expression) MangoData {
	result := interpreter.Evaluate(expressions[0])

	panic(NewMangoReturn("", result))
}
*/
