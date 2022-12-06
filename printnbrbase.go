package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	isValid := true
	baseLen := len(base)
	if baseLen <= 1 {
		isValid = false
		z01.PrintRune('N')
		z01.PrintRune('V')
	}
	for i := 0; i < baseLen-1; i++ {
		if base[i] == base[i+1] {
			isValid = false
			z01.PrintRune('N')
			z01.PrintRune('V')
		}
	}
	if isValid {
		numbers := ""
		isNegative := false
		if nbr < 0 {
			nbr = -nbr
			isNegative = true
		}
		for nbr != 0 {
			mod := nbr % baseLen
			numbers += string(base[mod])
			nbr = nbr / baseLen
		}
		numbers = StrRev(numbers)
		if isNegative {
			z01.PrintRune('-')
		}
		numbersLen := len(numbers)
		for i := 0; i < numbersLen; i++ {
			z01.PrintRune(rune(numbers[i]))
		}
	}
}
