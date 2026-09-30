package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

// some sort of OOPish construct
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() { // receiver function

	fmt.Println("Start print")

	for i, card := range d {
		fmt.Println(i, card)
	}

	fmt.Println("End print")
}

func newCard() string {
	return "Queen of Hearts"
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}
