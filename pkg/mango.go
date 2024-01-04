package mango

import (
	"log"
)

func Eval(source string) (result MangoData) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("[Runtime Error] Oops! Unhandled Error")
			result = NewMangoException("Unhandled Exception")
		}
	}()

	tokens, err := Tokenize(source)
	if err != nil {
		log.Fatal(err)
	}

	parser := NewParser()
	expressions := parser.Parse(tokens)

	interpreter := NewInterpreter()
	result = interpreter.Interpret(expressions)
	return result
}
