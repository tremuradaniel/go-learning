package main

func main() {

	cards := deck{newCard(), newCard()}
	cards = append(cards, "Another card")

	cards.print()

}

func newCard() string {
	return "Card name"
}
