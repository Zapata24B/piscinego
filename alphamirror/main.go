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
	mirror := "zyxwvutsrqponmlkjihgfedcba"
	runes := []rune(s)
	for i, c := range runes {
		if c >= 'a' && c <= 'z' {
			runes[i] = rune(mirror[int(c)-97])
		} else if c >= 'A' && c <= 'Z' {
			runes[i] = rune(mirror[int(c)-65]-32)
		}
	}
	return string(runes)
}
