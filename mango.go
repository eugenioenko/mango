package main

import (
	"fmt"
	mango "mango/pkg"
	"os"
)

func main() {

	args := os.Args
	if args[1] == "help" {
		help()
		return
	}

	mango.Eval(args[1])

}

func help() {
	fmt.Print(`
Mango is a minimal expression parser and interpreter

Usage:

  mango [expression]: evaluates the expression
	mango help: prints this message

`)
}
