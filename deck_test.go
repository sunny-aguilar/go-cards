package main

import (
	"os"
	"testing"
)

// to run test, type in terminal: go test

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 9 {
		t.Errorf("Expected deck length of 9 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Space, but got %v", d[0])
	}

	if d[len(d)-1] != "Three of Hearts" {
		t.Errorf("Expected last card of Three of Hearts, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 9 {
		t.Errorf("Expecting 9 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
