package main

import "fmt"

func main() {
	tri := triangle{base: 10, height: 25}
	squ := square{sideLength: 10}
	printArea(tri)
	printArea(squ)
}

type shape interface {
	area() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) area() float64 {
	return t.base * t.height * 0.5
}

func printArea(s shape) {
	fmt.Println(s.area())
}
