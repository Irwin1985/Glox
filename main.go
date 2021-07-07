package main

import (
	"Glox/interpreter"
	"Glox/lox"
	"fmt"
	"os"
)

func main() {
	i := interpreter.NewInterpreter()
	if len(os.Args) > 1 {
		fmt.Println("Usage: jlox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		lox.RunFile(os.Args[0], i)
	} else {
		lox.RunPrompt(i)
	}
}
