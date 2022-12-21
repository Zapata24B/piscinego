package main

import (
	"fmt"
	"piscine"
)

func main() {
	table := []int{1, 2, 3}
	ac := 93
	piscine.FoldInt(piscine.Add, table, ac)
	piscine.FoldInt(piscine.Mul, table, ac)
	piscine.FoldInt(piscine.Sub, table, ac)
	fmt.Println()

	table = []int{0}
	piscine.FoldInt(piscine.Add, table, ac)
	piscine.FoldInt(piscine.Mul, table, ac)
	piscine.FoldInt(piscine.Sub, table, ac)
}
