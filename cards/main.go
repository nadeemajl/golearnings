package main

import (
	"fmt"
)

func main() {

	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
	cards.print()
	fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	cardsFromFile := newDeckFromFile("my_cards")
	cardsFromFile.shuffle()
	cardsFromFile.print()
	oddandeven()
}

func oddandeven() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range ints {
		if value%2 == 0 {
			fmt.Println("value ", value, " is even")
		} else {
			fmt.Println("value ", value, " is odd")
		}
	}
}
