package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		arguments := os.Args[1:]
		for _, word := range arguments {
			if Index(word, "01") != -1 || Index(word, "galaxy") != -1 || Index(word, "galaxy 01") != -1 {
				fmt.Println("Alert!!!")
			}
		}
	}
}

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return -1
	}
	for i := 0; i <= len(s)-len(toFind); i++ {
		subs := s[i : i+len(toFind)]
		if subs == toFind {
			return i
		}
	}
	return -1
}
