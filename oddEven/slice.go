package main

import "fmt"

type slice []int

func newSlice() slice {
	s := slice{}
	number := 1

	for number < 11 {
		s = append(s, number)

		number++
	}

	return s
}

func (s slice) iterate() {
	fmt.Println("Iterating slice...")

	for _, num := range s {
		remainder := num % 2

		if remainder != 0 {
			fmt.Println("num", num, "is odd")
		} else {
			fmt.Println("num", num, "is even")
		}
	}
}
