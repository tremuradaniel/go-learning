package deck

import (
	"errors"
	"math/rand/v2"
	"slices"
	"strconv"
)

type Deck struct {
	cards []Card
}

const ACE_RANK = "A"
const KING_RANK = "K"
const QUEEN_RANK = "Q"
const JACK_RANK = "J"

var suites = [4]string{"♦️", "❤️", "♣️", "♠️"}
var values = func() [13]int {
	v := [13]int{}

	for i := range len(v) {
		v[i] = i + 1
	}

	return v
}()

func (d *Deck) New() *Deck {
	for _, r := range values {
		for _, s := range suites {
			d.cards = append(d.cards, newCard(s, r))
		}
	}

	return d
}

func (d Deck) PrintCurrentDeck() {
	for _, c := range d.cards {
		c.Print()
	}
}

func (d *Deck) Shuffle() {
	shuffledDeck := Deck{}

	for len(d.cards) > 0 {
		randomCardIndex := randRange(1, len(d.cards)) - 1
		shuffledDeck.cards = append(shuffledDeck.cards, d.cards[randomCardIndex])
		d.cards = slices.Delete(d.cards, randomCardIndex, randomCardIndex+1)
	}

	d.cards = shuffledDeck.cards
}

func randRange(min, max int) int {
	return min + rand.IntN(max-min+1)
}

func (d *Deck) DrawNCards(n int) (*[]Card, error) {
	var cardsLeft = len(d.cards)
	cardsDrawn := make([]Card, n)

	if n > cardsLeft {
		return nil, errors.New("Asked for " + strconv.FormatInt(int64(n), 10) +
			" cards, but " + strconv.FormatInt(int64(cardsLeft), 10) + " left!")
	}

	copy(cardsDrawn, d.cards[:n])

	d.cards = d.cards[n:]

	return &cardsDrawn, nil
}
