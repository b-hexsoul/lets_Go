package main // still inside our main package...

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardFaces := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardValues := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
	for _, face := range cardFaces {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+face)
		}
	}
	return cards
}

// specifying multiple return values with their types in parentheses
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// receiver function below
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Option 1: Log the error, and return a newDeck
		// Option 2: Log the error, and completely quit the program
		fmt.Println(err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

// Custom Type Declaration
// we want to "extend" a base type and add some functionality to it

// we will add function specifically for this "extended" type

// Make functions with the new "extended" deck type as a Receiver
// A function with a "Receiver" is like a method - a function that belongs to an instance
// func (d deck) print() // any variable of type deck now gets the func print...

// go has support for returning multiple values from a function

// values of type error. If nothing went wrong, it will have value of nil
