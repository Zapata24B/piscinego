package piscine

func Unmatch(a []int) int {
	for i := range a {
		for j := range a {
			if i != j {
				if a[i] == a[j] {
					break
				}
				if j == len(a)-1 {
					return a[i]
				}
			}
			if j == len(a)-1 {
				return a[i]
			}
		}
	}
	return -1
}
