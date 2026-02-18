package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	Reset      = "\033[0m"
	ErrorColor = "\033[38;5;203m" // soft red
	InfoColor  = "\033[38;5;221m" // warm yellow
	HintColor  = "\033[38;5;114m" // soft green
	LabelColor = "\033[38;5;245m" // neutral gray
	Bold       = "\033[1m"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: world <file.w | file.world>")
		return
	}

	filename := os.Args[1]

	if !isWorldFile(filename) {
		fmt.Println("Error: file must have .w or .world extension")
		return
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	run(string(data))
}

func run(code string) {
	lines := strings.Split(code, "\n")

	for lineNumber, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "display ") {
			executeDisplay(line, lineNumber)
			continue
		}

		syntaxError(lineNumber, line,
			"I don't understand this instruction.",
			"Please check your syntax on line number: "+strconv.Itoa(lineNumber+1))
	}
}

func isWorldFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".w" || ext == ".world"
}

func executeDisplay(line string, lineNumber int) {
	rest, ok := strings.CutPrefix(line, "display ")
	if !ok {
		syntaxError(lineNumber, line,
			"This display statement looks incorrect.",
			"Try using: display \"Hello\"")
	}

	rest = strings.TrimSpace(rest)

	if len(rest) < 2 {
		syntaxError(lineNumber, line,
			"display requires quoted text",
			"Example: display \"Hello, World\"")
	}

	if !strings.HasPrefix(rest, "\"") || !strings.HasSuffix(rest, "\"") {
		syntaxError(lineNumber, line,
			"display needs text inside quotes.",
			"Example: display \"Hello\"",
		)
	}
	text := rest[1 : len(rest)-1]
	fmt.Println(text)
}

func syntaxError(lineNumber int, line, message, hint string) {
	fmt.Println()

	fmt.Printf("%s%sWorld Syntax Error%s\n",
		Bold, ErrorColor, Reset)

	fmt.Printf("%sline:%s %d\n",
		LabelColor, Reset, lineNumber+1)

	fmt.Printf("%sproblem:%s %s\n",
		LabelColor, Reset, message)

	fmt.Printf("%scode:%s\n",
		LabelColor, Reset)

	fmt.Printf("  %s\n", line)

	fmt.Printf("%shint:%s %s\n",
		LabelColor, Reset, HintColor+hint+Reset)

	fmt.Println()

	os.Exit(1)
}
