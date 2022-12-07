package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	isUpper := false
	for _, v := range arguments {
		if v == "--upper" {
			isUpper = true
		}
	}
	for _, arg := range arguments {
		number := 0
		number = miniAtoi(arg, number)
		if number >= 1 && number <= 26 {
			if !isUpper {
				z01.PrintRune(rune(number + 96))
			} else {
				z01.PrintRune(rune(number + 64))
			}
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func miniAtoi(arg string, number int) int {
	for _, v := range arg {
		number = number*10 + int(v-'0')
	}
	return number
}
