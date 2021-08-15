package main

import "fmt"

func main() {

	cards := newDeckFromFile("my_cards")
	fmt.Println(cards)
}
