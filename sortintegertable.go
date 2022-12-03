package piscine

// Bubble Sort
func SortIntegerTable(table []int) {
	for i := range table {
		for j := 0; j < len(table); j++ {
			if table[i] < table[j] {
				Swap(&table[i], &table[j])
			}
		}
	}
}
