package main

import (
	"fmt"
	"os"
	"strings"
)

func asciiArt(input string, lines []string) string {
	if input == "" {
		return ""
	}
	if input == "\\n" {
		return "\n"
	}

	var sb strings.Builder
	parts := strings.Split(input, "\\n")
	for _, part := range parts {
		if part == "" {
			sb.WriteString("\n")
		} else {
			for row := 0; row < 8; row++ {
				for _, ch := range part {
					startIndex := (int(ch)-32)*9 + 1
					sb.WriteString(lines[startIndex+row])
				}
				sb.WriteString("\n")
			}
		}
	}
	return sb.String()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go \"your text\"")
		os.Exit(1)
	}

	input := os.Args[1]

	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")
	fmt.Print(asciiArt(input, lines))
}
