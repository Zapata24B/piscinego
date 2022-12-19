package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	game := make([][]int, 4)
	current := 0
	for i := 0; i < len(deck); i++ {
		game[current] = append(game[current], deck[i])
		current = (current + 1) % 4
	}
	for i := 0; i < 4; i++ {
		fmt.Print("Player ", i+1, ": ")
		for j := 0; j < len(game[i]); j++ {
			fmt.Print(game[i][j])
			if j != len(game[i])-1 {
				fmt.Print(", ")
			} else {
				fmt.Println()
			}
		}
	}
}
