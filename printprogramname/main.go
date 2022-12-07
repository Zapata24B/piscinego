package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	filePath := os.Args[0]
	fileName := ""
	for i := len(filePath) - 1; i >= 0; i-- {
		if filePath[i] == '/' {
			fileName = filePath[i+1:]
			break
		}
	}
	for i := 0; i < len(fileName); i++ {
		z01.PrintRune(rune(fileName[i]))
	}
	z01.PrintRune(rune('\n'))
}
