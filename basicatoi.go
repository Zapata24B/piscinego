package piscine

func BasicAtoi(s string) int {
	number := 0
	d := 1
	
	for i := len(s) - 1; i >= 0; i-- {
		number += (int(rune(s[i])) - 48) * d
		d = d * 10 
	}

	return number
}
