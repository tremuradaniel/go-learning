package main

import "fmt"

func main() {

	// var colors map[string][string]

	colors := make(map[int]string)

	// colors := map[string]string{
	// 	"red": "#FF0000",
	// }

	colors[1] = "blue"

	fmt.Println(colors)

	delete(colors, 1)

	fmt.Println(colors)
}
