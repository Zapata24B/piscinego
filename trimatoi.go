package piscine

func TrimAtoi(s string) int {
	numbers, isNegative := Itoa(s)
	n := Atoi(string(numbers))
	if isNegative {
		return -n
	}
	return n
}

func Itoa(s string) ([]rune, bool) {
	numbers := []rune{}
	isNegative := false
	firstDigit := false
	for i := 0; i < len(s); i++ {
		if !firstDigit {
			if s[i] == '-' {
				isNegative = true
			}
		}
		if IsNumeric(string(s[i])) {
			numbers = append(numbers, rune(s[i]))
			firstDigit = true
		}
	}
	return numbers, isNegative
}
