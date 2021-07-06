package main

import (
	"Glox/lox"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Usage: jlox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		lox.RunFile(os.Args[0])
	} else {
		lox.RunPrompt()
	}
}
