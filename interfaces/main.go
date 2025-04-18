package main

import "fmt"

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)


	triangle := triangle{
		height: 53.31,
		base: 87.33,
	}

	square := square{
		sideLength: 324,
	}

	printArea(triangle)
	printArea(square)
}

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func (eb englishBot) getGreeting() string {
	return "Hello There"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
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
	fmt.Print(s.getArea())
} 