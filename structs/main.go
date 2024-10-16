package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipcode int
}

func main() {
	p := person{
		firstName: "John",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email:   "john.doe@gmail.com",
			zipcode: 12345,
		},
	}

	p.print()
	p.updateLastName("Ross")
	p.print()
}

func (p *person) updateLastName(newLastName string) {
	p.lastName = newLastName
}

func (p *person) print() {
	fmt.Printf("%+v\n", p)
}
