package main

import "fmt"

type bot interface {
	getGreeting() string
}

// type bot interface
// our program has a new type called 'bot'

// getGreeting() string
// if you are a type in this program with a function getGreeting
// and you return a string, you are now associated with type bot
// you can now call this specific function defined

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello"
}
func (sb spanishBot) getGreeting() string {
	return "Hola"
}

// concrete vs interface types

// interfaces are NOT generic types
// interfaces are implicit - can make dificult to see if your type implements which interfaces
// interfaces help us manage types - NOT a test for your code - Garbage in => garbage out
