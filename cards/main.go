package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"   -- long way
	// card := newCard()

	cards := newDeck()
	cards = append(cards, "Joker")

	hand, remainingDeck := deal(cards, 5)

	fmt.Println()
	hand.print()
	fmt.Println()
	remainingDeck.print()

	remainingDeck.saveToFile("remaining_deck")

	newRemainingDeck := newDeckFromFile("remaining_deck")
	fmt.Println()
	newRemainingDeck.print()
}
