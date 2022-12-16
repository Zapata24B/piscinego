package piscine

func Compact(ptr *[]string) int {
	newPtr := []string(nil)
	count := 0
	for _, str := range *ptr {
		if str == "" {
			count++
		} else {
			newPtr = append(newPtr, str)
		}
	}
	*ptr = newPtr
	return count
}
