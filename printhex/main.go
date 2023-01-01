package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		if nb, err := strconv.Atoi(os.Args[1]); err == nil && nb >= 0 {
			base := "0123456789abcdef"
			numbers := []rune(nil)
			for nb != 0 {
				numbers = append(numbers, rune(base[nb%16]))
				nb /= 16
			}
			for i := 0; i < len(numbers)/2; i++ {
				numbers[i], numbers[len(numbers)-i-1] = numbers[len(numbers)-i-1], numbers[i]
			}
			printStr(string(numbers))
		} else {
			printStr("ERROR")
		}
	}
}

func printStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
