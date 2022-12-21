package piscine

import "fmt"

func FoldInt(f func(int, int) int, a []int, n int) {
	for _, el := range a {
		n = f(n, el)
	}
	fmt.Println(n)
}

func Add(x, y int) int {
	return x + y
}

func Mul(x, y int) int {
	return x * y
}

func Sub(x, y int) int {
	return x - y
}
