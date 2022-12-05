package piscine

func Sqrt(nb int) int {
	if nb <= 1 {
		return 0
	} else {
		for i := 2; i <= nb/2; i++ {
			if i*i == nb {
				return i
			}
		}
		return 0
	}
}
