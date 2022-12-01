package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	slc := []int{}
	i := 0
	if n < 0 {
		n = -1 * n
		z01.PrintRune('-')
	} else if n <= 0 {
		z01.PrintRune('0')
	}
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10
		i++
	}
	for i := len(slc) - 1; i >= 0; i-- {
		z01.PrintRune(rune(48 + slc[i]))
	}
}
