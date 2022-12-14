package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		actual := string(s[i])
		if i == 0 {
			if IsAlpha(actual) && !IsNumeric(actual) && IsLower(actual) {
				runes[i] = rune(ToUpper(string(runes[i]))[0])
			}
		} else {
			previous := string(s[i-1])
			if IsAlpha(actual) && !IsAlpha(previous) {
				if !IsNumeric(actual) && IsLower(actual) {
					runes[i] = rune(ToUpper(string(runes[i]))[0])
				}
			} else {
				runes[i] = rune(ToLower(string(runes[i]))[0])
			}
		}
	}
	return string(runes)
}
