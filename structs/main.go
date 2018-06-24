package main

import (
	"fmt"
)

type contactinfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contactinfo
}

func main() {

	batook := person{
		firstName: "Ayan",
		lastName:  "Lalani",
		contactinfo: contactinfo{
			email:   "batook@batook.in",
			zipcode: 5037,
		},
	}

	batook.print()
	batook.updateName("Rahil")
	batook.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
