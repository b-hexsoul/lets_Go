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

// ******* SLICE *******
// slice[a:b]
// a = start index (inclusive)
// b = end index (exclusive)
// adding items to a slice
// numbers := []int{1,2,3,4}
// numbers = append(numbers, 4)

// You can combine two slices
// part1 := []int{1,2,3}
// part2 := []int{4,5,6}
// combined := append(part1, part2...)

// Preallocation
// Useful when number of elements are known, but their values are still uknown
// slice := make([]int, 10)

// ******* MAPS ********
// key-value pairs
// myMap := make(map[string]int)
// myMap := map[string]int{
// 	"Item 1": 1,
// 	"Item 2": 2,
// 	"Item 3": 3,
// }

// CHECKING EXISTENCE IN A MAP
// price, found := myMap["price"]
// if !found {
// 	fmt.Println("Not found")
// }

// for key, value := range myMap {}
