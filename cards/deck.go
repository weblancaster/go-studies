package main

import (
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{
		"Spades",
		"Diamonds",
		"Hearts",
		"Clubs",
	}
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			formatted := value + " of " + suit

			cards = append(cards, formatted)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
	
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}