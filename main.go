package main

import "fmt"

type contractInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contractInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Patty",
		contractInfo: contractInfo{
			email:   "test@test.com",
			zipCode: 303212,
		},
	}

	jim.updateName("Jimmy")
	jim.print()

}

func (p person) updateName(newName string) {
	p.firstName = newName
}

func (p person) print() {
	fmt.Printf("%-v", p)
}
