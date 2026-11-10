package deck

import (
	"fmt"
	"strconv"
)

type Card struct {
	suite string
	rank  string
	value int
}

func newCard(s string, v int) Card {
	return Card{
		suite: s,
		rank:  determineRank(v),
		value: v,
	}
}

func determineRank(v int) string {
	switch v {
	case 1:
		return ACE_RANK
	case 11:
		return JACK_RANK
	case 12:
		return QUEEN_RANK
	case 13:
		return KING_RANK
	default:
		return strconv.FormatInt(int64(v), 10)
	}
}

func (c Card) GetSuite() string {
	return c.suite
}

func (c Card) GetRank() string {
	return c.rank
}

func (c Card) Print() {
	fmt.Println(c.rank + " " + c.suite)
}
