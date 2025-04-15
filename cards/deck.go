package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

/**
Create a new type of deck, which is a slice of strins
*/

type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Jack", "Queen", "King"}

	for _, cardSuite := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuite)
		}
	}

	return cards

}

// Receiver function
// cards deck is a receiver
// receiver is usually a one letter abreviation of the type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
/**
- we didn't use receiver because we don't want to modify the underlying list of cards
**/
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/*
*
Byte slice
"Hi there" -> [72 105 32 116 ...]
Here, we are doing the following

	    DECK -> We have
		[]String
		String
		[]Byte -> We want

*
*/
func (d deck) toString() string {
	cards := strings.Join([]string(d), ",")
	return cards
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

/*
*
Byte slice
"Hi there" -> [72 105 32 116 ...]
Here, we are doing the following
		[]Byte -> We have
	    String
		[]String
		DECK -> We want

*
*/

func readFromFile(fileName string) deck {
	deckBytes, err := os.ReadFile(fileName)
	if err != nil {
		//Option #1 - log the error and return a call to newDeck()
		//Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}
	deckString := string(deckBytes)
	cards := strings.Split(deckString, ",")
	return deck(cards)
}

func (d deck) shuffle() {

	/**
		Random Number Generation in Go:
		Go's math/rand package requires a source of randomness
		Without a source, you'd get the same "random" sequence every time you run the program
		Why Two Variables?:
		source: Creates a source of randomness based on the current time (time.Now().Unix())
		r: Creates a random number generator using that source
	**/
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
