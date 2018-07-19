package main

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}

/* func newCard() string {
	return "Five of Diamonds"
}
*/
