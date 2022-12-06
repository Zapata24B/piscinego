package piscine

func IsUpper(s string) bool {
	runes := []rune(s)
	isUpper := true
	for i := 0; i < len(s); i++ {
		if runes[i] <= 'A' || runes[i] >= 'Z' {
			isUpper = false
		}
	}
	return isUpper
}
