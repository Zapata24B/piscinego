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

	firstWord := ""
	inWord := false

	for _, char := range input {
		if char == ' ' {
			if inWord {
				break
			}
		} else {
			firstWord += string(char)
			inWord = true
		}
	}
	if firstWord != "" {
		PrintStr(firstWord)
	}
}

func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
