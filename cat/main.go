package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		input, err := io.ReadAll(os.Stdin)
		printStr(string(input))
		if err != nil {
			printStr(err.Error())
		}
	} else if len(args) == 1 {
		fileName := args[0]
		content, err := os.ReadFile(fileName)
		if err != nil {
			printStr("ERROR: open " + fileName + ": No such file or directory")
			z01.PrintRune('\n')
			return
		}
		printStr(string(content))
	} else {
		for _, fileName := range args {
			content, err := os.ReadFile(fileName)
			if err != nil {
				printStr("ERROR: " + fileName + ": No such file or directory")
				z01.PrintRune('\n')
				break
			}
			printStr(string(content))
		}
	}
}