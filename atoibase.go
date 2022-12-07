package piscine

func AtoiBase(s string, base string) int {
	baseLen := len(base)
	if ValidateBase(baseLen, base) {
		number := 0
		factor := 1
		numberLen := len(s)
		for i := numberLen - 1; i >= 0; i-- {
			number += Index(base, string(s[i])) * factor
			factor *= baseLen
		}
		return number
	}
	return 0
}
