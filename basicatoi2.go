package piscine

func BasicAtoi2(s string) int {
	number := 0
	d := 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		number += (int(rune(s[i])) - 48) * d
		d = d * 10
	}
	return number
}
