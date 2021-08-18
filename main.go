package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person
	fmt.Println(alex)
	fmt.Printf("%+v", alex)

	var alexa person
	alexa.firstName = "Alexa"
	alexa.lastName = "Stan"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
