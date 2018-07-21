package main

import "fmt"

/*
	In Go, we do not manually have to say that custom type satisfies some interface
	we do not explicitly write anywhere saying type implements this interface
  -------------------------------------------------------------------------------
	Our Program has a new type called 'bot'

	If you are a type in this program with a function called 'getGreeting' and you return a string
	Congrats, you are now an honorary member of type 'bot'

	Now that you are an honorary member of type 'bot', you can call 'printGreeting'

	Here, both englishBot and spanishBot has a function called 'getGreeting' that returns a string
	so, they are member of 'bot'
	Hence, printGreeting can be called
*/
type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	printGreeting(eb)

	sb := spanishBot{}
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// Custom logic specific for english Bot
	return "Hello"
}

func (sb spanishBot) getGreeting() string {
	// Custom logic specific for spanish Bot
	return "Hola"
}

/*
func (englishBot) getGreeting() string {
	we can omit 'eb' here as we are not using it
}
*/
