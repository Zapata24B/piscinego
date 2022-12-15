package piscine

import "fmt"

func ReverseBits(oct byte) byte {
	result := byte(0)
	fmt.Printf("%08b\n", oct)
	bit := oct
	for i := 0; i < 8; i++ {
		fmt.Printf("%08b\n", result)
		bit = (bit >> 1) & 1
		fmt.Printf("%b\n", bit)
		result = result | (bit << (7 - i))
	}
	return result
}
