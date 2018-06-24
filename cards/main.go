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
}
