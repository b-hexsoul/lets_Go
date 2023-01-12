package main

import "fmt"

func main() {

}

type shape interface {
	area() int
}

type square struct {
	sideLength int
}

func (s square) area() int {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.area())
}
