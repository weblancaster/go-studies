package main

import (
	"bytes"
	"testing"
)

func TestNewDeck(t *testing.T) {
	expected := deck{
		"Ace of Spades",
		"Two of Spades",
		"Three of Spades",
		"Four of Spades",
		"Ace of Diamonds",
		"Two of Diamonds",
		"Three of Diamonds",
		"Four of Diamonds",
		"Ace of Hearts",
		"Two of Hearts",
		"Three of Hearts",
		"Four of Hearts",
		"Ace of Clubs",
		"Two of Clubs",
		"Three of Clubs",
		"Four of Clubs",
	}
	result := newDeck()

	if len(result) != len(expected) {
		t.Errorf("expected %v to be equals to %v", result, expected)
	}
}

func TestDeal(t *testing.T) {
	expectedStart := deck{
		"Ace of Spades",
		"Two of Spades",
	}
	expectedEnd := deck{
		"Three of Spades",
		"Four of Spades",
	}
	deck := deck{
		"Ace of Spades",
		"Two of Spades",
		"Three of Spades",
		"Four of Spades",
	}
	resultStart, resultEnd := deal(deck, 2)

	resultS := bytes.Compare([]byte(expectedStart.toString()), []byte(resultStart.toString()))
	resultE := bytes.Compare([]byte(expectedEnd.toString()), []byte(resultEnd.toString()))

	if resultS != 0 {
		t.Errorf("expected %v to be equals to %v", resultStart, expectedStart)
	}

	if resultE != 0 {
		t.Errorf("expected %v to be equals to %v", resultEnd, expectedEnd)
	}
}