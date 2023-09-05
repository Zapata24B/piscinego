package main

import (
	"fmt"
	"os"
)

func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return ""
	}

	const charset = "0123456789ABCDEF"
	var result string

	if value < 0 {
		result = "-"
		value = -value
	}

	for value > 0 {
		remainder := value % base
		result = string(charset[remainder]) + result
		value /= base
	}

	return result
}

func main() {
	if len(os.Args) == 3 {

		valueStr := os.Args[1]
		baseStr := os.Args[2]

		value := Atoi(valueStr)

		base := Atoi(baseStr)
		if base < 2 || base > 16 {
			fmt.Println("Invalid base. Must be between 2 and 16.")
			return
		}

		result := ItoaBase(value, base)
		fmt.Println(result)
	}
}

func Atoi(s string) int {
	number := 0
	if len(s) > 0 {
		factor := 1
		for i := len(s) - 1; i >= 0; i-- {
			if (s[i] < '0' || s[i] > '9') && (i != 0 || (s[0] != '-' && s[0] != '+')) {
				return 0
			}
			if s[i] != '-' && s[i] != '+' {
				number += (int(s[i]) - 48) * factor
				factor = factor * 10
			}
		}
		if s[0] == '-' {
			return -number
		}
	}
	return number
}
