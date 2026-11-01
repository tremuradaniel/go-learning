package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGeeting(eb)
	printGeeting(sb)
}

func (eb englishBot) getGreeting() string {
	// VERY custom logic for generatin an english greeting DIFFERENT FROM SPANISH...
	return "Hi there"
}

func (sb spanishBot) getGreeting() string {
	// VERY custom logic for generatin a spanish greeting DIFFERENT FROM ENGLISH...
	return "Ola!"
}

func printGeeting(b bot) {
	fmt.Println(b.getGreeting())
}
