package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

// Create a new type called 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// here word 'deck' is a reference to the type that we want to attach this function to
// d is the actual copy of the deck we're working with and it is available in function
// one or two letter abbreviation. ( by convention ) 'd'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	err := ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	if err != nil {
		return err
	}
	return nil

	// return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option 1 - log the error and return a call to newDeck()
		// Option 2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	ss := strings.Split(string(bs), ",")
	return []string(ss) // return deck(ss)
}
