package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	// option := args[0]
	start := Atoi(args[1]) - 1
	files := args[2:]
	for i, fileName := range files {
		content, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Printf("open %v: no such file or directory\n", fileName)
			continue
		}
		toPrint := string(content)
		toPrint = toPrint[len(toPrint)-start:]
		if len(files) == 1 {
			fmt.Printf("%v\n", toPrint)
		} else {
			if i >= 1 {
				fmt.Printf("\n")
			}
			fmt.Printf("==> %v <==\n%v\n", fileName, toPrint)
		}
	}
}

func Atoi(s string) int {
	number := 0
	factor := 1
	for i := len(s) - 1; i >= 0; i-- {
		number += (int(s[i]) - 48) * factor
		factor = factor * 10
	}
	return number
}
