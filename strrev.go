package piscine

func StrRev(s string) string {
	reverse := []rune(s)
	j := 0
	for i := len(s) - 1; i >= 0; i-- {
		reverse[j] = rune(s[i])
		j++
	}
	return string(reverse)
}
