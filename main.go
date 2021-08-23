package main

import "fmt"

func main() {

	// var colors map[string][string]

	colors := map[string]string{
		"first":  "#ffff",
		"second": "#ffff2",
		"third":  "#ffff3",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
