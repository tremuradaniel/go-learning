package game

import (
	"fmt"
	"log"
	"slices"
	"strings"
)

const DEFAULT_CARDS_FIRST_DRAWN = 2

const FIRST_DRAW_PHASE = 0
const PLAYER_PHASE = 1
const DEALER_PHASE = 2

type game struct {
	player    HumanPlayer
	dealer    Dealer
	gamePhase int
}

func InitGame() {
	g := game{
		player:    HumanPlayer{},
		dealer:    *NewDealer(),
		gamePhase: FIRST_DRAW_PHASE,
	}
	// var d = deck.Deck{}

	// d.New()

	// d.PrintCurrentDeck()

	for true {
		if !run(&g) {
			break
		}
	}
}

func run(g *game) bool {

	start := startNewGame()

	if !start {
		return false
	}

	g.playRound()

	return true
}

func startNewGame() bool {
	var loop = true

	for loop {
		var w1 string
		var allowed_inputs = [2]string{"y", "n"}

		fmt.Println("Begin new game? [y/n]")

		n, err := fmt.Scanln(&w1)

		if err != nil {
			log.Fatal(err)
			break
		}

		if !slices.Contains(allowed_inputs[:], w1) {
			fmt.Println(w1, " is invalid! allowed: ", strings.Join(allowed_inputs[:], " or "))
			continue
		}

		if w1 == "n" {
			loop = false
			break
		}

		fmt.Println("number of items read: ", n)
		fmt.Printf("read line: %s \n", w1)
		return true
	}

	return false
}

func (g *game) playRound() {
	g.drawFirstCards()
	g.displayCards(FIRST_DRAW_PHASE)
	// initial draw stage
	// player turn
	// dealer turn

	// for _, card := range *cards {
	// 	card.Print()
	// }
}

func (g *game) drawFirstCards() {
	g.player.AddCardsToHand(*g.dealer.DrawNCards(DEFAULT_CARDS_FIRST_DRAWN))
	g.dealer.AddCardsToHand(*g.dealer.DrawNCards(DEFAULT_CARDS_FIRST_DRAWN))
}

func (g *game) displayCards(phase int) {
	fmt.Println("Dealer Hand:")
	fmt.Println(g.dealer.hand.Show(getCardsDisplayedByDealer(phase)))

	fmt.Println("Player Hand:")
	fmt.Println(g.player.hand.Show(DISPLAY_ALL))
}

func getCardsDisplayedByDealer(phase int) int {
	if phase != DEALER_PHASE {
		return DISPLAY_FIRST_CARD
	}

	return DISPLAY_ALL
}
