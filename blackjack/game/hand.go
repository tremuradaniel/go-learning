package game

import "blackjack/deck"

const DISPLAY_ALL = -1
const DISPLAY_NONE = 0

type Hand struct {
	cards []deck.Card
}

func (h *Hand) add(cards []deck.Card) {
	for _, c := range cards {
		h.cards = append(h.cards, c)
	}
}

func (h *Hand) discard() {
	clear(h.cards)
}

func (h Hand) show(showFirstNCards int) string {
	cs := "["
	for _, c := range h.cards {
		if showFirstNCards <= DISPLAY_NONE {
			cs += "??"
		}

		if showFirstNCards == DISPLAY_ALL || showFirstNCards > 0 {
			cs += c.GetToDisplay()
			showFirstNCards--
		}
	}

	cs += "]"

	return cs
}
