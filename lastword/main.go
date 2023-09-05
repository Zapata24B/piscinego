package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	words := splitWords(input)

	if len(words) > 0 {
		lastWord := words[len(words)-1]
		PrintStr(lastWord)
	}
}

func splitWords(input string) []string {
	var words []string
	start := -1

	for i, char := range input {
		if char == ' ' {
			if start != -1 {
				words = append(words, input[start:i])
				start = -1
			}
		} else {
			if start == -1 {
				start = i
			}
		}
	}

	if start != -1 {
		words = append(words, input[start:])
	}

	return words
}

func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
