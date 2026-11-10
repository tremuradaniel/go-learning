package main

import (
	"blackjack/deck"
	"fmt"
	"log"
	"slices"
	"strings"
)

func main() {
	var w1 string
	var loop bool = true
	var allowed_inputs = [2]string{"y", "n"}

	var d = deck.Deck{}

	d.New()

	d.PrintCurrentDeck()

	for loop {
		fmt.Println("Stop loop[y/n]:")

		n, err := fmt.Scanln(&w1)

		if err != nil {
			log.Fatal(err)
			break
		}

		if !slices.Contains(allowed_inputs[:], w1) {
			fmt.Println(w1, " is invalid! allowed: ", strings.Join(allowed_inputs[:], " or "))
			continue
		}

		if w1 == "y" {
			loop = false
			return
		}

		fmt.Println("number of items read: ", n)
		fmt.Printf("read line: %s \n", w1)
	}

}
