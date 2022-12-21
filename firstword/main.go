package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		word := []rune(nil)
		for _, char := range arg {
			if char != ' ' {
				word = append(word, char)
			} else {
				break
			}
		}
		PrintStr(string(word))
	}
}

func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
