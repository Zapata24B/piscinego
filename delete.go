package piscine

func Delete(ints []int, position int) []int {
	position--
	if position < 0 || position > len(ints) {
		return ints
	}
	return append(ints[:position], ints[position+1:]...)
}
