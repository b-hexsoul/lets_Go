package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height, base float64
}

type square struct {
	sideLength float64
}

func printArea(s shape) {
	fmt.Println("The area is", s.getArea())
}

func main() {
	tri := triangle{height: 5, base: 3}
	sq := square{sideLength: 5}
	printArea(sq)
	printArea(tri)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.height * t.base * .5
}
