package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff000",
	// 	"green": "#324r3",
	// }

	// var colors map[string]string

	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	colors["black"] = "#110000"

	delete(colors, "black")

	fmt.Println(colors)
}
