package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	isFirstLetterWord := true
	isOnWord := true
	for i := 0; i < len(s); i++ {
		if IsAlpha(string(s[i])) {
			if isFirstLetterWord && isOnWord {
				isFirstLetterWord = false
				isOnWord = true
			} else if !isFirstLetterWord && !isOnWord {
				isFirstLetterWord = true
				isOnWord = true
			}
		} else {
			isOnWord = false
		}
		if isFirstLetterWord && IsLower(string(s[i])) {
			runes[i] = rune(ToUpper(string(runes[i]))[0])
			isFirstLetterWord = false
		}
	}
	return string(runes)
}
