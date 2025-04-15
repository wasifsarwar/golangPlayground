package main

import (
	"os"
	"testing"
)

/**
testing new deck
- Code to make sure that a deck is created with x number of cards
- Code to make sure that the first card is an Ace of Spades
- Code to make sure the last card is a King of Clubs
*/
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 48 {
		t.Errorf("Expected deck of length of 48, but got %v", len(d))
	}

	if d[0] != "Ace of Spades"  {
		t.Errorf("Expected Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d) - 1])
	}
}

/**
Testing save to file
- Delete any files in current working directory with the name "_decktesting"
- Create a deck
- save to file "_decktesting"
- load from a file
- assert deck length
- Delete any files in current working directory with the name "_decktesting"
**/

func TestWriteToFileAndReadFromFile(t *testing.T) {
	os.Remove("_decktesting")
	
	deck := newDeck()
	deck.saveToFile("_decktesting")
	
	readDeck := readFromFile("_decktesting")
	if len(readDeck) != 48 {
		t.Errorf("Expected 48 cards in deck, got %v cards", len(deck))
	}

	os.Remove("_decktesting")
}


func TestDeckDeal(t *testing.T) {
	deck := newDeck()

	hand, remaining := deal(deck, 12)
	if len(hand) != 12 {
		t.Errorf("Expected 12 cards in hand, got %v cards in hand", len(hand))
	}

	if len(remaining) != 36 {
		t.Errorf("Expected 36 cards remaining, got %v remaining cards", len(remaining))
	}
}

func TestShuffle(t *testing.T) {
	deck := newDeck()
	deck.shuffle()

	if deck[0] == "Ace of Spades" {
		t.Errorf("Wasn't expecting first card to be Ace of Spades, and got %v", deck[0])
	}

	if deck[len(deck) - 1] == "King of Clubs" {
		t.Errorf("Wasn't expecting last card to be King of Clubs, but got %v", deck[len(deck) - 1])
	}
}