package main

import "fmt"

type contact struct {
	email string
	zip   int
}

type person struct {
	firstName   string
	lastName    string
	age         int
	contactInfo contact
}

func main() {
	dan := person{
		firstName: "dan",
		lastName:  "ban",
		contactInfo: contact{
			zip:   23123123,
			email: "test@email.com",
		},
	}

	fmt.Printf("%+v", dan)
}
