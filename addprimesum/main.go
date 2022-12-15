package main

import (
	"fmt"
	"os"
)

func main() {
	numArg := len(os.Args)
	sum := 0
	if numArg == 2 {
		num := BasicAtoi(os.Args[1])
		for i := num - 1; i > 0; i-- {
			if IsPrime(i) {
				sum += i
			}
		}
	}
	fmt.Print(sum)
}

func BasicAtoi(s string) int {
	number := 0
	factor := 1
	for i := len(s) - 1; i >= 0; i-- {
		number += (int(s[i]) - 48) * factor
		factor = factor * 10
	}
	return number
}

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	for i := 3; i <= nb/3; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
