package mango

func Eval(source string) (result []MangoData, err error) {
	tokens, err := Tokenize(source)
	if err != nil {
		return nil, err
	}

	statements, err := Parse(tokens)
	if err != nil {
		return nil, err
	}

	result, err = Interpret(statements)
	if err != nil {
		return nil, err
	}
	return result, nil
}
