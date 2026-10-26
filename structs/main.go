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
	contact
}

func main() {
	dan := person{
		firstName: "dan",
		lastName:  "ban",
		contact: contact{
			zip:   23123123,
			email: "test@email.com",
		},
	}

	dan.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}
