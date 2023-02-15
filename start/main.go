package main

// Types of Packages in Golang
// 1. Executable
// This type generates a file we can run

// package main is "sacred" and will spit out an executable. Must have a func called main

// 2. Resuable
// This type is for code used as "helpers".

import "fmt" // a part of the standard library in Go

func main() {
	for v := range [11]int{} {
		if v%2 == 0 {
			fmt.Printf("%v is even \n", v)
		} else {
			fmt.Printf("%v is odd \n", v)
		}
	}
}

// `go mod init project-name` and then `go mod tidy` in project root
