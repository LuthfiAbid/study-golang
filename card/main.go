package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	// cards := newDeckFromFile("myDeck")
	// cards.print()

	// cards := newDeck()

	// cards.saveToFile("myDeck")

	// Setting card in hand and remainingDeck
	// hand, remainingDeck := deal(cards, 5)

	// hand.print("Hand = ")
	// remainingDeck.print("Remaining Deck = ")
}
