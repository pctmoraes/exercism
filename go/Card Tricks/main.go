package main

import (
	"fmt"
	"cards/cards"
)

func main() {
	cards1 := cards.FavoriteCards()
	fmt.Println(cards1)
	// Output: [2 6 9]

	card2 := cards.GetItem([]int{1, 2, 4, 1}, 2) // card == 4
	fmt.Println(card2)

	card2 = cards.GetItem([]int{1, 2, 4, 1}, 10) // card == -1
	fmt.Println(card2)

	index := 2
	newCard := 6
	cards3 := cards.SetItem([]int{1, 2, 4, 1}, index, newCard)
	fmt.Println(cards3)
	// Output: [1 2 6 1]

	slice := []int{3, 2, 6, 4, 8}
	cards4 := cards.PrependItems(slice, 5, 1)
	fmt.Println(cards4)
	// Output: [5 1 3 2 6 4 8]

	slice = []int{3, 2, 6, 4, 8}
	cards4 = cards.PrependItems(slice)
	fmt.Println(cards4)
	// Output: [3 2 6 4 8]

	cards5 := cards.RemoveItem([]int{3, 2, 6, 4, 8}, 2)
	fmt.Println(cards5)
	// Output: [3 2 4 8]


}
