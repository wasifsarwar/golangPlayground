package main

import "fmt"

func main() {

	triangle := triangle{
		height: 10,
		base: 8,
	}

	square := square{
		sideLength: 12,
	}

	printArea(triangle)
	printArea(square)
}
/**
Exercise: Write a program that creates two custom struct types triangle and square
Square struct has field with side length of type float64
Triangle struct has fields called height and base, both float64
Both types should have a function called getArea that returns the calculated area of the square or the triangle
Add a shape interface that calculates the area of the given shape, so printArea can print both triangle/square areas
*/



type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface  {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return (0.5 * t.height * t.base)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
} 