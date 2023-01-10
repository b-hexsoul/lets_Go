package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// order of fields in structs matters if you dont use property names!
	// brandon := person{"Brandon", "Hexsel"}
	// brandon := person{firstName: "Brandon", lastName: "Hexsel", contact: contactInfo{ email:"b@b.com", zipcode: 12345, }}
	var brandon person
	brandon.firstName = "Brandon"
	brandon.lastName = "Hexsel"
	brandon.contact = contactInfo{email: "b@b.com", zipcode: 12345}

	// brandonPointer := &brandon
	// unnecessary to use the above. can simply pass in the struct
	brandon.updateFirstName("Brenda")
	brandon.print()
}

func (pointerToPerson *person) updateFirstName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p person) print() {
	fmt.Printf("My name is %v %v", p.firstName, p.lastName)
}

// structs. a data structure. collection of related properties. Similar to a plain object in JS.
// For our playing card example. Suit -> String, Value -> String

// used widely throughout the language
// Printf with %+v will print out fieldnames and values!

// A struct can be embedded inside another struct
// Structs can also have receiver functions

// Go is a Pass By Value language!
// Use pointers with structs to solve pass by value issue

// * and & operators...

// &variableName. Give us access to the memory address this variable is pointing at
// eg. jimPointer := &jim

// *pointer. Within a function, Give us the value this memory address is pointing to
// func (p *person) - a asterisk in front of a type is a type description meaning we are working with a pointer for that type
