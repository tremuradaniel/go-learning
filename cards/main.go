package main

func main() {
	// cards := newDeck()

	// cards.saveToFile("my_cards")

	// cards = append(cards, "Six of Spades")

	// hand, cards := deal(cards, 2)

	// hand.print()

	// cards.print()

	deck := createFromFile("my_cards")

	deck.print()
}
