package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	baseLen := len(base)
	occurence := 0
	for i := 0; i < baseLen-1; i++ {
		occurence = 0
		for j := 0; j < baseLen; j++ {
			if base[i] == base[j] {
				occurence += 1
				if occurence >= 2 {
					break
				}
			}
		}
		if occurence >= 2 {
			break
		}
	}
	if baseLen <= 1 || occurence >= 2 || string(base[0]) == "-" || string(base[0]) == "+" {
		z01.PrintRune('N')
		z01.PrintRune('V')
	} else {
		numbers := ""
		isNegative := false
		if nbr < 0 {
			nbr = -nbr
			isNegative = true
		}
		if nbr != 0 {
			for nbr != 0 {
				mod := nbr % baseLen
				numbers += string(base[mod])
				nbr = nbr / baseLen
			}
		} else {
			numbers = string(base[0])
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
