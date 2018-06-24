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
	batook.updateName("rahil")
	batook.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
