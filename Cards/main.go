package main

func main() {

	// Created a new deck of playing cards, essentially an array
	// of strings. The contents of the deck of cards is then logged
	// out to the console.
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
