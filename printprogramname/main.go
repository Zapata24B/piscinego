package main

import (
	"os"

	"piscine"
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
	piscine.PrintStr(fileName+("\n"))
}
