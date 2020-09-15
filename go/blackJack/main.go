package main

import (
	"fmt"
)

var suite = map[rune]int{'A': 11, 'K': 10, 'Q': 10, 'J': 10, '9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2}

func handTotal(hand string) int {
	total := 0
	for _, i := range hand {
		var n int
		if v, found := suite[i]; found {
			n = v
		}
		total = total + n
	}
	fmt.Println(total)
	return total
}

// PlayerWins determines if the player has a winning blackjack hand.
func PlayerWins(playerHand string, dealerHand string) bool {
	playerHandTotal := handTotal(playerHand)
	dealerHandTotal := handTotal(dealerHand)
	if playerHandTotal > dealerHandTotal {
		return true
	}
	return false

}

func main() {

	if !PlayerWins("A9", "J8") {
		panic("Fail 1")
	}
	fmt.Println("pass 1")
	if PlayerWins("54X", "KQ") {
		panic("Fail 2")
	}
	fmt.Println("pass 2")
	if !PlayerWins("279", "368") {
		panic("fail 3")
	}
	fmt.Println("pass 3")
}
