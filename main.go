package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"world/internal/interpreter/direct"
)

func main() {
	// Using direct string interpreter
	interp := *direct.New()

	// Check file type
	if len(os.Args) < 2 {
		fmt.Println("Usage: world <file.w | file.world>")
		return
	}

	filename := os.Args[1]

	if !isWorldFile(filename) {
		fmt.Println("Error: file must have .w or .world extension")
		return
	}

	// Load code data
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Run the code
	interp.Run(string(data))
}

// Checks if the file is a valid World file
func isWorldFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".w" || ext == ".world"
}
