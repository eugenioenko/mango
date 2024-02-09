package mango

import (
	"fmt"
	"os"
)

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

func Run(filename string) {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("[Mango CLI Error] %s", err)
		os.Exit(1)
	}
	source := string(file)
	_, err = Eval(source)
	if err != nil {
		fmt.Printf("%s", err)
	}
}
