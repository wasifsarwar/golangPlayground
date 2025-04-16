package main

import "fmt"

/**
Exercise: Write a program that creates two custom struct types triangle and square
Square struct has field with side length of type float64
Triangle struct has fields called height and base, both float64
Both types should have a function called getArea that returns the calculated area of the square or the triangle
Add a shape interface that calculates the area of the given shape, so printArea can print both triangle/square areas
*/

type Square struct {
	sideLength float64
}

type Triangle struct {
	height float64
	base   float64
}

type Shape interface  {
	getArea() float64
}

func (t Triangle) getArea() float64 {
	return (0.5 * t.height * t.base)
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func PrintArea(s Shape) {
	fmt.Print(s.getArea())
}