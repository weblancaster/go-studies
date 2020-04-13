package main

import (
	"fmt"
)

func main() {
	fileName := "new-deck.txt"
	cardsFromFile := getDeckFromFile(fileName)
	// hand, remainingDeck := deal(cards, 5)

	// hand.print();
	// remainingDeck.print();

	// fmt.Println(cards.toString())

	fmt.Println(cardsFromFile)
}