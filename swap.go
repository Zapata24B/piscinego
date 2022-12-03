package piscine

func Swap(a *int, b *int) {
	// Python like swapping
	// *a, *b = *b, *a
	c := *a
	*a = *b
	*b = c
}
