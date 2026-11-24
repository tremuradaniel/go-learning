package game

import "blackjack/deck"

type HumanPlayer struct {
	hand Hand
}

func (hp *HumanPlayer) AddCardsToHand(cards []deck.Card) {
	hp.hand.add(cards)
}
