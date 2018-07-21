package main

import (
	"fmt"       // print function
	"io/ioutil" // ReadFromFile function
	"math/rand" // shuffle function
	"os"        // WriteToFle function
	"strings"   // ReadFromFile function
	"time"      // shuffle function
)

// Create a new type called 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// here word 'deck' is a reference to the type that we want to attach this function to
// d is the actual copy of the deck we're working with and it is available in function
// one or two letter abbreviation. ( by convention ) 'd'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
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

	s := strings.Split(string(bs), ",")
	return []string(s) // return deck(s)
}

/*
	for each index, card in cards
		generate a random number between 0 and len(cards)-1
		swap the current card and card at cards[randomNumber]
*/
func (d deck) shuffle() {
	// To generate new random number, generating a new seed every time
	source := rand.NewSource(time.Now().UnixNano()) // time.Now().UnixNano() generates new int64 every time
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
