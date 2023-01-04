package main

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	_, _ = deal(remainingCards, 3)
}

// Delcaring Variables
// var card string
// var card string = "queen"
// card := newCard() // will infer the type

// you have to declare the return type from your functions
// func newCard() string {
// 	return "Five of Diamonds"
// }

// Array - Fixed length list of items. Must delcare the length at initiation
// Slice - An array that can grow or shrink. Every element must be of the same type
// cards := []string{"Ace of Diamonds", newCard()}
// cards = append(cards, "Six of Diamonds") // does not modify the cards slice but returns a new one

// For Loop iteration over a slice
/*
for index, item := range items {
	fmt.Println(index, item)
}
*/

// slice is delcared with []type
// cards := deck{"Ace of Diamonds", newCard()}
// cards = append(cards, "Six of Spades") // does not modify the cards slice. returns a new slice
