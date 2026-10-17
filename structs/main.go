package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{"John", "Doe"}
	alex := person{lastName: "John", firstName: "Doe"}

	fmt.Println(alex)
}
