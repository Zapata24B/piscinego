package main

import (
	"fmt"
	"piscine"
)

func main() {
	result := piscine.ReverseBits(byte(10))
	fmt.Printf("%08b", result)
}
