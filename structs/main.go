package main

import "fmt"

type contact struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contact
}

func main() {
	// alex := person{"John", "Doe"}
	// alex := person{lastName: "John", firstName: "Doe"}

	var alex person

	alex.age = 33

	alex.contact.email = "alex.jognson@mail.com"

	fmt.Println(alex)

	fmt.Printf("%+v", alex)
}
