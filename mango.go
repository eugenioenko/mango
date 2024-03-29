package main

import (
	"fmt"
	mango "mango/pkg"
	"os"
)

var (
	version = "0.3.0"
)

func main() {

	args := os.Args

	if len(args) == 1 {
		help()
		return
	}

	if args[1] == "--version" {
		fmt.Println(version)
	}

	if args[1] == "--help" {
		help()
		return
	}

	mango.Run(args[1])
}

func help() {
	fmt.Print(`
Mango is a minimal expression parser and interpreter

Usage:
  mango [options] [file] [arguments]

	--version	prints mango version
	--help		prints this message

`)
}
