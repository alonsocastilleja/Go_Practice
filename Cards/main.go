package main

func main() {

	// Created a new deck of playing cards, essentially an array
	// of strings. The contents of the deck of cards is then logged
	// out to the console.
	cards := newDeck()

	cards.print()
}
