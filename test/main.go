package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Delete([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(piscine.Delete([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(piscine.Delete([]int{1, 2, 3, 4, 5}, 1))
}