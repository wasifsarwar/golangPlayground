package main

import "fmt"

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

type bot interface{
	getGreeting() string
}

type englishBot struct {}

type spanishBot struct {}

func (eb englishBot) getGreeting() string {
	return "Hello There"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}


func (sb spanishBot) getGreeting() string {
	return "Hola"
}