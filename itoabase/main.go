package main

import (
	"fmt"
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
	value := -42
	base := 16
	result := ItoaBase(value, base)
	fmt.Printf("%d in base %d is %s\n", value, base, result)
}
