package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	file, err := os.ReadFile(args[1])

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Printf("%v", string(file))

}
