package main

// import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// cards := newDeckFromFile("my_cardss")
	// cards.print()
	// cards.saveToFile("my_cards")
	// fmt.Println(cards.toString())

	// type conversion
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
}
