package main

import "fmt"

type contractInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contractInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Patty",
		contact: contractInfo{
			email:   "test@test.com",
			zipCode: 303212,
		},
	}
	fmt.Printf("%-v", jim)
}
