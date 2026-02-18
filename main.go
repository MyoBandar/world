package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
	for line := range strings.SplitSeq(code, "\n") {
		line = strings.TrimSpace(line)

		if text, ok := strings.CutPrefix(line, "display "); ok {
			text = strings.Trim(text, "\"")
			fmt.Println(text)
		}
	}
}

func isWorldFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".w" || ext == ".world"
}
