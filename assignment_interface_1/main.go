package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height, base float64
}
type square struct {
	sideLength float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func main() {
	t := triangle{}
	s := square{}

	s.sideLength = 10

	t.base = 10
	t.height = 10

	printArea(t)
	printArea(s)

}
