package main

// import "fmt"

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	cards := newDeckFromFile("my_cards")
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	cards.saveToFile("my_cards")
}

/* func newCard() string {
	return "Five of Diamonds"
}
*/
