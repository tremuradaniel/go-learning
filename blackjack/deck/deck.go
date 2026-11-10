package deck

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

func (d *Deck) New() {
	for _, r := range values {
		for _, s := range suites {
			d.cards = append(d.cards, newCard(s, r))
		}
	}
}

func (d Deck) PrintCurrentDeck() {
	for _, c := range d.cards {
		c.Print()
	}
}
