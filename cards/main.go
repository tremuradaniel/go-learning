package main

func main() {
	cards := newDeck()

	cards = append(cards, "Six of Spades")

	hand, cards := deal(cards, 2)

	hand.print()

	cards.print()
}
