package main

import (
	"fmt"
)

/**
* Array: Fixed length list of things
* Slices: An array that can grow or shrink
* Every element in a slice must be of the same type
 */

func main() {
	cards := newDeck()
	
	hand, remainingCards := deal(cards, 12)
	hand.print()
	fmt.Println()
	remainingCards.print()
	fmt.Println(cards.toString())
	hand.saveToFile("my_cards")
	fmt.Println("cards being read below:")
	readCards := readFromFile("my_cards")
	readCards.print()

	fmt.Println("randomized deck:")
	cards.shuffle()
	cards.print()
}

