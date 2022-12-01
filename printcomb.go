package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				if i != j && i != k && i < j && j < k {
					z01.PrintRune(rune(48 + i))
					z01.PrintRune(rune(48 + j))
					z01.PrintRune(rune(48 + k))
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
