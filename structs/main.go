package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	// alex := person{"John", "Doe"}
	// alex := person{lastName: "John", firstName: "Doe"}

	var alex person

	fmt.Println(alex)

	fmt.Printf("%+v", alex)
}
