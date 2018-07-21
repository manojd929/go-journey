package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type contactInfo struct {
	email   string
	zipcode int
}

type employee struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p person) print() {
	fmt.Println(p)
	fmt.Printf("%+v \n %#v \n\n", p, p)
}

func (e employee) updateName(newFirstName string) {
	e.firstName = newFirstName
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

	/* embedding one struct in another */
	jim := employee{
		firstName: "Jim",
		lastName:  "Carry",
		contact: contactInfo{
			email:   "jim@jmail.com",
			zipcode: 94000,
		},
	}
	jim.updateName("Jimmy")
	fmt.Printf("%+v", jim)
}

/*
This is also valid definition

type employee struct {
	firstName string
	lastName  string
	contactInfo
}
*/
