package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	lastCard := 0
	for i := 0; i < 4; i++ {
		fmt.Print("Player ", i+1, ": ")
		for i := lastCard; i < lastCard+3; i++ {
			fmt.Print(deck[i])
			if i != lastCard+2 {
				fmt.Print(", ")
			}
		}
		fmt.Println()
		lastCard += 3
	}
}
