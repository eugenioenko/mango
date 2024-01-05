package mango

func Eval(source string) (result []MangoData, err error) {
	tokens, err := Tokenize(source)
	if err != nil {
		return nil, err
	}

	expressions, err := Parse(tokens)
	if err != nil {
		return nil, err
	}

	result, err = Interpret(expressions)
	if err != nil {
		return nil, err
	}
	return result, nil
}
