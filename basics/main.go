package main

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	cards := newDeck()
	cards.print()
	cards.shuffle()
	cards.print()
}

/*
func newCard() string {
	return "Five of Diamonds"
}
*/
