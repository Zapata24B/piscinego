package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.LeapYear(2020))
	fmt.Println(piscine.LeapYear(2021))
	fmt.Println(piscine.LeapYear(2022))
	fmt.Println(piscine.LeapYear(-10))
}
