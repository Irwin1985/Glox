package lox

import (
	"Glox/interpreter"
	"Glox/parser"
	"Glox/scanner"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// hadError will ensure don't try to execute code that has a known error.
var HadError = true

func RunFile(path string, i *interpreter.Interpreter) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("file does not exist: %s", path)
		os.Exit(65)
	}
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	run(string(bytes), i)
	if HadError {
		os.Exit(65)
	}
}

func RunPrompt(i *interpreter.Interpreter) {
	mode := "console"
	//mode := "debug"
	if mode != "debug" {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Welcome to Glox, an GoLang implementation of Lox.")
		fmt.Printf("Data and Time: %s\n", time.Now().Format(time.Stamp))
		fmt.Println("Type 'quit' to exit")
		for {
			fmt.Print("> ")
			scanned := scanner.Scan()
			if !scanned {
				return
			}
			line := scanner.Text()
			if len(line) <= 0 {
				continue
			}
			if line == "quit" {
				break
			}
			run(line, i)
		}
	} else {
		line := `
		// Your first Lox program!
		var name = "Irwin";
		print name;
		`
		run(line, i)
	}
}

func run(source string, i *interpreter.Interpreter) bool {
	s := scanner.NewScanner(source)
	tokens := s.ScanTokens()
	if len(s.Errors()) > 0 {
		printErrors(s.Errors())
		return false
	}

	p := parser.NewParser(tokens)
	statements := p.Parse()
	if len(p.Errors()) > 0 {
		printErrors(p.Errors())
		return false
	}

	i.Interpret(statements)

	return true
}

func printErrors(errors []string) {
	for _, msg := range errors {
		fmt.Println(msg)
	}
}
