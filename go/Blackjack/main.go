package main

import (
	"fmt"
	"blackjack/blackjack"
)

func main() {

	// value := blackjack.ParseCard("ace")
	// fmt.Println(value)
	// // Output: 11
	value := blackjack.FirstTurn("ace", "ace", "jack")
	fmt.Println(value) // "P"

	value = blackjack.FirstTurn("ace", "king", "ace")
	fmt.Println(value) // "S"

	value = blackjack.FirstTurn("five", "queen", "ace")
	fmt.Println(value) // "H"

	value = blackjack.FirstTurn("ace", "ten", "two")
	fmt.Println(value) // "W"

}