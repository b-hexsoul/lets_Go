package main

import "fmt"

func main() {
	// colors := make(map[string]string)
	// colors["red"] = "warning"
	colors := map[string]string{
		"red":    "error",
		"blue":   "info",
		"orange": "warning",
	}
	// delete(colors, "red")
	printMap(colors)
	fmt.Println(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}

// Maps are key vale pair. In Go, all keys must be of the same type in a map. All values must be of the same type
// variable := map[key type]value type
// To initialize a map, use the built in make function:
// colors := make(map[string]string)

// From Go Docs:
// var m map[string]int

// Map types are reference types, like pointers or slices, and so the value of m above is nil; it doesn't point to an initialized map.
// A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic; don't do that.
// To initialize a map, use the built in make function:

// m = make(map[string]int)

// The make function allocates and initializes a hash map data structure and returns a map value that points to it.
// The specifics of that data structure are an implementation detail of the runtime and are not specified by the language itself.

// Iteration over a Map
/* func printMap(c map[type]type){
	for key, value := range c {

	}
}
*/
