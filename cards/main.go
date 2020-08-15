package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"   -- long way
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
