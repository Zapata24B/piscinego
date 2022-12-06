package piscine

func Concat(str1 string, str2 string) string {
	runes1 := []rune(str1)
	runes2 := []rune(str2)
	return string(runes1)+string(runes2)
}
