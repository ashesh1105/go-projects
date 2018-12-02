package main

import (
	"fmt"
)

type ContanctInfo struct {
	phone   string
	address string
}

type Person struct {
	firstName string
	lastName  string
	// Below, you can just say ContactInfo and you will be fine, don't need to define a variable
	// Compiler will read it as ContactInfo ContactInfo
	contact ContanctInfo
}

// Set Receiever Functions for Person type
func (p Person) print() {
	fmt.Printf("%+v", p)
}

// Set first name of a person
func (p Person) setFirstName(firstName string) {
	p.firstName = firstName
}

// Set first name of a person
func (pointerToPerson *Person) setFirstNameWithPointer(firstName string) {
	(*pointerToPerson).firstName = firstName
}
