package piscine

import "fmt"

func IsLower(s string) bool {
	runes := []rune(s)
	isLower := true
	for i := 0; i < len(s); i++ {
		if runes[i] < 'a' || runes[i] > 'z' {
			fmt.Println(string(runes[i]))
			isLower = false
		}
	}
	return isLower
}