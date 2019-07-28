package main

import (
	"fmt"
)

func main() {
	s := square{sideLength: 0.5}
	s.printArea()

	t := triangle{height: 0.4, base: 0.6}
	t.printArea()
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}
type shape interface {
	printArea()
}

func (t triangle) printArea() {
	fmt.Println(t.height * t.base / 2)
}

func (s square) printArea() {
	fmt.Println(s.sideLength * s.sideLength)
}
