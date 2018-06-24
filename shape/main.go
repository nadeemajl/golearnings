package main

import (
	"fmt"
)

type triangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

type shape interface {
	getArea() float64
}

func main() {

	t := triangle{base: 12, height: 10}
	s := square{side: 5}

	printArea(t)
	printArea(s)

}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (sq square) getArea() float64 {
	return sq.side * sq.side
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
