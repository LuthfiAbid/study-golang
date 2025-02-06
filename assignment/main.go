package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	square := square{4}
	triangle := triangle{4, 5}

	printArea("result area of square :", square)
	printArea("result area of triangle :", triangle)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func printArea(typeArea string, s shape) {
	fmt.Println(typeArea, s.getArea())
}
