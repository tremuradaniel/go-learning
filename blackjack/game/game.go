package game

import (
	"fmt"
	"log"
	"slices"
	"strings"
)

type game struct {
	player HumanPlayer
	dealer Dealer
}

const DEFAULT_CARDS_FIRST_DRAWN = 2

func InitGame() {
	g := game{
		player: HumanPlayer{},
		dealer: *NewDealer(),
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

	// startNewGame()

	playRound(g)

	return false
}

func startNewGame() {
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
			return
		}

		fmt.Println("number of items read: ", n)
		fmt.Printf("read line: %s \n", w1)
	}
}

func playRound(g *game) {
	cards := g.dealer.DrawNCards(2)

	for _, card := range *cards {
		card.Print()
	}
}
