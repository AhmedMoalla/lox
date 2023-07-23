package main

import (
	"bufio"
	"fmt"
	"github.com/AhmedMoalla/lox/errors"
	"github.com/AhmedMoalla/lox/lexer"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Usage: lox [script]")
		os.Exit(64)
	} else if len(args) == 2 {
		runFile(args[1])
	} else {
		runPrompt()
	}
}

func runFile(file string) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		panic("could not read file at " + file)
	}

	run(string(bytes))
	if errors.HadError {
		os.Exit(65)
	}
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			break
		}
		run(line)
		errors.HadError = false
	}
}

func run(source string) {
	tokens := lexer.New(source).Tokenize()
	for _, token := range tokens {
		fmt.Println(token)
	}
}
