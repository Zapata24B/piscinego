package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	// Sort numerical args only
	for i := range arguments {
		runes := []rune(arguments[i])
		j := 0
		if runes[j] >= '0' && runes[j] <= '9' {
			for _, j := range arguments[i] {
				z01.PrintRune(j)
				z01.PrintRune('\n')
			}
		}
	}
	// Sort uppercase args only
	for i := range arguments {
		runes := []rune(arguments[i])
		j := 0
		if runes[j] >= 'A' && runes[j] <= 'Z' {
			for _, j := range arguments[i] {
				z01.PrintRune(j)
				z01.PrintRune('\n')
			}
		}
	}
	// Sort lowercase args only
	for i := range arguments {
		runes := []rune(arguments[i])
		j := 0
		if runes[j] >= 'a' && runes[j] <= 'z' {
			for _, j := range arguments[i] {
				z01.PrintRune(j)
				z01.PrintRune('\n')
			}
		}
	}
}
