package piscine

func StrLen(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum++
	}
	return sum
}
