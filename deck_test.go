package main

import "testing"

// to run test, type in terminal: go test

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 9 {
		t.Errorf("Expected deck length of 9 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Space, but got %v", d[0])
	}

	if d[len(d) - 1] != "Three of Hearts" {
		t.Errorf("Expected last card of Three of Hearts, but got %v", d[len(d) - 1])
	}
}
