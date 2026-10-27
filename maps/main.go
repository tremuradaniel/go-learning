package main

import "fmt"

func main() {

	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	colors["black"] = "#110000"
	colors["red"] = "#e91919"

	printMap(colors)

	fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color + " has hex " + hex)
	}
}
