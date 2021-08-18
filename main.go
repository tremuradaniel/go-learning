package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{"Alex", "Tim"}
	// or
	alexa := person{firstName: "Alexa", lastName: "Stuart"}
	fmt.Println(alex)
	fmt.Println(alexa)
}
