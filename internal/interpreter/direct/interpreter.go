// Package direct implements the direct string interpreter for the World language.
//
// The direct interpreter reads source code line by line and executes instructions
// by matching known string patterns. It does not use tokenization or parsing.
//
// This implementation exists primarily for educational purposes, demonstrating
// the simplest possible way to execute a programming language.
//
// It serves as the foundation and historical first implementation of World.
package direct

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"world/internal/errors"
)

type Interpreter struct{}

func New() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) Run(code string) {
	lines := strings.Split(code, "\n")

	for lineNumber, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "display ") {
			i.executeDisplay(line, lineNumber)
			continue
		}

		errors.SyntaxError(lineNumber, line,
			"I don't understand this instruction.",
			"Please check your syntax on line number: "+strconv.Itoa(lineNumber+1))

		fmt.Println()
		os.Exit(0)
	}
}

func (i *Interpreter) executeDisplay(line string, lineNumber int) {
	rest, ok := strings.CutPrefix(line, "display ")
	if !ok {
		errors.SyntaxError(lineNumber, line,
			"This display statement looks incorrect.",
			"Try using: display \"Hello\"")
	}

	rest = strings.TrimSpace(rest)

	if len(rest) < 2 {
		errors.SyntaxError(lineNumber, line,
			"display requires quoted text",
			"Example: display \"Hello, World\"")
	}

	if !strings.HasPrefix(rest, "\"") || !strings.HasSuffix(rest, "\"") {
		errors.SyntaxError(lineNumber, line,
			"display needs text inside quotes.",
			"Example: display \"Hello\"",
		)
	}
	text := rest[1 : len(rest)-1]
	fmt.Println(text)
}
