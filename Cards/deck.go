package main // still inside our main package...

import "fmt"

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Custom Type Declaration
// we want to "extend" a base type and add some functionality to it

// we will add function specifically for this "extended" type

// Make functions with the new "extended" deck type as a Receiver
// A function with a "Receiver" is like a method - a function that belongs to an instance
// func (d deck) print() // any variable of type deck now gets the func print...
