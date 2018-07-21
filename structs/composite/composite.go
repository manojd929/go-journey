package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type employee struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	/* embedding one struct in another */
	jim := employee{
		firstName: "Jim",
		lastName:  "Carry",
		contact: contactInfo{
			email:   "jim@jmail.com",
			zipcode: 94000,
		},
	}

	// Pass By Value
	jim.updateName("Jimmy")
	jim.print()

	// Pass By Reference
	jimPointer := &jim
	jimPointer.updateNameThroughPointer("Jimmy")
	jim.print()

	// Pass By Reference - Go allows this even though method is accepting pointer through a person
	jim.updateNamePointerShortcut("Jean")
	jim.print()
}

func (e employee) updateName(newFirstName string) {
	e.firstName = newFirstName
}

func (e *employee) updateNameThroughPointer(newFirstName string) {
	(*e).firstName = newFirstName
	// e.firstName = newFirstName works fine (Pointer Shortcut)
}

func (e *employee) updateNamePointerShortcut(newFirstName string) {
	(*e).firstName = newFirstName
}

func (e employee) print() {
	fmt.Printf("%+v \n", e)
}

/*
This is also valid definition

type employee struct {
	firstName string
	lastName  string
	contactInfo
}
*/
