package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{"Alex", "Trotskiy"}
	alex := person{firstName: "Alex",
		lastName: "Trotskiy",
		contact: contactInfo{
			email:   "some@email.here",
			zipCode: 12531,
		},
	}
	fmt.Println(alex)

	var bobby person
	bobby.firstName = "Bobby"
	bobby.lastName = "Kotik"
	bobby.contact.email = "kotik@email.com"

	fmt.Println(bobby)
}
