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

	bobby.print()
	bobbyPointer := &bobby
	bobbyPointer.updateName("Bobbies")
	bobby.print()

	//Still works! Without direct pointer call.
	alex.updateName("Alexander")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// Go copies input variables.
// To update a specific object we need pointers.
func (p *person) updateName(name string) {
	(*p).firstName = name
}

//Gotchas:
// 1. Structs(strings, ints etc.) are pass-by values, they are copied inside funcs
// 2. Slices are NOT copied, they are modified.
