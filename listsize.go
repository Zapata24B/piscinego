package piscine

func ListSize(l *List) int {
	count := 0
	actual := l.Head
	for actual.Next != nil {
		count++
		actual = actual.Next
	}
	return count + 1
}
