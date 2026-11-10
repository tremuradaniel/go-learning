package deck

import "testing"

func TestGeneratingNewDeck(t *testing.T) {
	d := Deck{}
	const TOTAL_CARDS_EXPECTED = 52

	d.New()

	var actual_total = len(d.cards)

	if actual_total != TOTAL_CARDS_EXPECTED {
		t.Errorf("Expected deck to have %d cards, got %d", TOTAL_CARDS_EXPECTED, actual_total)
	}
}
