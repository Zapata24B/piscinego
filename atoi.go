package piscine

func Atoi(s string) int {
	number := 0
	d := 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] < '0' || s[i] > '9' {
			if i != 0 || (s[0] != '-' && s[0] != '+') {
				return 0
			}
		}
		if s[i] != '-' && s[i] != '+' {
			number += (int(rune(s[i])) - 48) * d
			d = d * 10
		}
	}
	if s[0] == '-' {
		return -number
	} else {
		return number
	}
}
