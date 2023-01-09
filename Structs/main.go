package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// order of fields in structs matters if you dont use property names!
	// brandon := person{"Brandon", "Hexsel"}
	// brandon := person{firstName: "Brandon", lastName: "Hexsel"}
	var brandon person
	brandon.firstName = "Brandon"
	brandon.lastName = "Hexsel"
	fmt.Printf("%+v", brandon)
}

// structs. a data structure. collection of related properties. Similar to a plain object in JS.
// For our playing card example. Suit -> String, Value -> String

// used widely throughout the language
// Printf with %+v will print out fieldnames and values!
