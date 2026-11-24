package game

import (
	"blackjack/deck"
)

type Dealer struct {
	hand Hand
	deck *deck.Deck
}

func NewDealer() *Dealer {
	return &Dealer{
		deck: getNewShuffledDeck(),
	}
}

// dealer-ul ar trebui sa fie cel care orchestreaza jocul:
// amesteca pachetul
// scoate carti din pachet
// da cartile jucatorului sau in mana sa
// genereaza un nou pachet cand ramane fara

func (d *Dealer) DrawNCards(n int) *[]deck.Card {
	cards, err := d.deck.DrawNCards(n)

	// TODO : maybe a better way to catch specific errors
	if nil != err {
		d.deck = getNewShuffledDeck()
		cards, _ := d.deck.DrawNCards(n)
		return cards
	}

	return cards
}

func getNewShuffledDeck() *deck.Deck {
	d := deck.Deck{}
	d = *d.New()
	d.Shuffle()

	return &d
}
