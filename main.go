package main

// import "fmt"

func main() {
	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	cards := newDeck()
	cards.saveToFile("my_cards")
	// fmt.Println(cards.toString())

	// type conversion
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
}
