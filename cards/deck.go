package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"math/rand"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

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

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func getDeckFromFile(fileName string) deck {
	byteSlice, error := ioutil.ReadFile(fileName)

	if error != nil {
		fmt.Println("No deck found. Generating a new one")
		cards := newDeck()
		cards.saveToFile(fileName)
		
		return getDeckFromFile(fileName)
	}

	str := strings.Split(string(byteSlice), ",")

	return deck(str)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for i := range d {
		newPos := random.Intn(len(d) - 1)

		d[i], d[newPos] = d[newPos], d[i]
	}
}