package piscine

func SplitWhiteSpaces(s string) []string {
	arr := []string(nil)
	lastWhiteSpace := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			arr = append(arr, s[lastWhiteSpace:i])
			lastWhiteSpace = i + 1
		}
	}
	return arr
}
