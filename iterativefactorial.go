package piscine

func IterativeFactorial(nb int) int {
	if nb > 1 && nb < 2147483647 {
		result := 1
		for i := nb; i >= 1; i-- {
			result *= i
			if result > 2147483647 {
				result = 0
				break
			}
		}
		return result
	} else if nb == 0 || nb == 1 {
		return 1
	} else {
		return 0
	}
}
