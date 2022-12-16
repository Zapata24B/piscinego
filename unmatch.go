package piscine

func Unmatch(a []int) int {
	for i := range a {
		for j := range a {
			if i != j && a[i] == a[j] {
				break
			} else if j == len(a)-1 && a[i] != a[j] {
				return a[i]
			}
		}
	}
	return -1
}
