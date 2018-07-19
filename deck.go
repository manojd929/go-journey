package main

import "fmt"

// Create a new type called 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	
	cardSuits := []string{ "Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit + " of " + value)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// here word 'deck' is a reference to the type that we want to attach this function to
// d is the actual copy of the deck we're working with is available in function
// one or two letter abbreviation by conventions 'd'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
