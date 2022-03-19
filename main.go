package main

import (
	"fmt"
)

func main() {
	cards := deck{newCard(), newCard()}
	cards = append(cards, "Ace of Diamonds")

	cards.print()

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
