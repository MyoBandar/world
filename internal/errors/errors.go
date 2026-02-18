// Package errors defines error types and formatting for the World language.
//
// It provides structured error types and beginner-friendly error messages,
// including learning-focused formatting and helpful hints.
//
// This package separates error representation from interpreter logic,
// allowing consistent and educational feedback across all World components.
package errors

import (
	"fmt"

	"world/internal/colors"
)

func SyntaxError(lineNumber int, line, message, hint string) {
	fmt.Println()

	fmt.Printf("%s%sWorld Syntax Error%s\n",
		colors.Bold, colors.ErrorColor, colors.Reset)

	fmt.Printf("%sline:%s %d\n",
		colors.LabelColor, colors.Reset, lineNumber+1)

	fmt.Printf("%sproblem:%s %s\n",
		colors.LabelColor, colors.Reset, message)

	fmt.Printf("%scode:%s\n",
		colors.LabelColor, colors.Reset)

	fmt.Printf("  %s\n", line)

	fmt.Printf("%shint:%s %s\n",
		colors.LabelColor, colors.Reset, colors.HintColor+hint+colors.Reset)
}
