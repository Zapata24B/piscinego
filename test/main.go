package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsMultiple(3))
	fmt.Println(piscine.IsMultiple(7))
	fmt.Println(piscine.IsMultiple(8))
	fmt.Println(piscine.IsMultiple(9))
	fmt.Println(piscine.IsMultiple(-1))
}
