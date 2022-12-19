package piscine

func Divisors(n int) int {
	count := 2
	for i := 2; i < n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}
