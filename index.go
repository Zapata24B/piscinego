package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return -1
	}
	for i := 0; i < len(s)-len(toFind); i++ {
		for j := 0; j < len(toFind); j++ {
			if s[i+j] != toFind[j] {
				break
			} else if j == len(toFind)-1 {
				return i
			}
		}
	}
	return -1
}
