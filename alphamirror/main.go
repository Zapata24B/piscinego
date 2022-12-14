package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		strMirror := Mirror(os.Args[1])
		fmt.Print(strMirror)
	}
	fmt.Println()
}

func Mirror(s string) string {
	runes := []rune(s)
	for i, c := range runes {
		if c >= 'a' && c <= 'z' {
			shift := 25 - ((c - 97) * 2)
			if c/2 <= 110 {
				runes[i] += shift
			} else {
				runes[i] -= shift
			}
		} else if c >= 'A' && c <= 'Z' {
			shift := 25 - ((c - 65) * 2)
			if c/2 <= 78 {
				runes[i] += shift
			} else {
				runes[i] -= shift
			}
		}
	}
	return string(runes)
}
