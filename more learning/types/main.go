package main

func main() {

}

// Structures
// Allow data to be stored in groups
// Each data point in a structure is called a field
// It is possible to associate functionality with structures to organize code/data

/*
type Sample struct {
	field string
	a, b  int
}
*/

// Any fields not indicated during instantiation will have default values

// Anonymous structures
// it's possible to create anonymous/inline structures inside of a function
// useful when working with library functions or shipping data across a network

// Anonymous structures example
/*
var Sample struct {
	field string
	a, b  int
}
OR
sample := struct {
	field string
	a, b  int
}{
	"hello",
	1, 2
}
*/

// *********** ARRAYS *************
// Store multiple pieces of the same kind of data
// Arrays are fixed size in Go and cannot be resized
/*
myArray := [3]int{1,2,3}
myArray := [...]int{1,4,3} best way to initialize an array
*/
