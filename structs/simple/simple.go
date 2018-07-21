package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (p person) print() {
	fmt.Println(p)
	fmt.Printf("%+v \n %#v \n\n", p, p)
}

func main() {
	/* different ways of declaring values for type struct */
	alex := person{"alex", "ferguson"}
	alex.print()

	craig := person{firstName: "craig", lastName: "mcmillan"}
	craig.print()

	var kane person
	kane.print()
	kane.firstName = "kane"
	kane.lastName = "anderson"
	kane.print()
}
