package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	baseLen := len(base)
	if ValidateBase(baseLen, base) {
		number := ""
		isNegative := false
		isMinInt := false
		nbr, isMinInt, isNegative = handleNegative(nbr, isMinInt, isNegative)
		if nbr != 0 {
			for nbr != 0 {
				mod := nbr % baseLen
				number += string(base[mod])
				nbr /= baseLen
			}
		} else {
			number = string(base[0])
		}
		number = StrRev(number)
		if isNegative {
			z01.PrintRune('-')
		}
		numberLen := len(number)
		if isMinInt {
			mod := 8 % baseLen
			_number := []rune(number)
			_number[numberLen-1] = rune(base[mod])
			number = string(_number)
		}
		for i := 0; i < numberLen; i++ {
			z01.PrintRune(rune(number[i]))
		}
	} else {
		z01.PrintRune('N')
		z01.PrintRune('V')
	}
}

func handleNegative(nbr int, isMinInt bool, isNegative bool) (int, bool, bool) {
	if nbr < 0 {
		if nbr == -9223372036854775808 {
			isMinInt = true
			nbr = -(nbr + 1)
		} else {
			nbr = -nbr
		}
		isNegative = true
	}
	return nbr, isMinInt, isNegative
}

func ValidateBase(baseLen int, base string) bool {
	if baseLen <= 1 {
		return false
	} else if string(base[0]) == "-" || string(base[0]) == "+" {
		return false
	} else {
		occurence := 0
		for i := 0; i < baseLen-1; i++ {
			occurence = 0
			for j := 0; j < baseLen; j++ {
				if base[i] == base[j] {
					occurence += 1
					if occurence == 2 {
						break
					}
				}
			}
			if occurence == 2 {
				break
			}
		}
		return occurence != 2
	}
}
