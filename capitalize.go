package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	isFirstLetterWord := false
	isOnWord := false
	for i := 0; i < len(s); i++ {
		if IsAlpha(string(s[i])) {
			if isFirstLetterWord && isOnWord {
				isFirstLetterWord = false
			} else if !isFirstLetterWord && !isOnWord {
				isFirstLetterWord = true
			}
			isOnWord = true
		} else {
			isOnWord = false
		}
		if isFirstLetterWord {
			if !IsNumeric(string(s[i])) && IsLower(string(s[i])) {
				runes[i] = rune(ToUpper(string(runes[i]))[0])
			}
			isFirstLetterWord = false
		} else if isOnWord {
			runes[i] = rune(ToLower(string(runes[i]))[0])
		}
	}
	return string(runes)
}
