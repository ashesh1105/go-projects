package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	someDeck := newDeck()
	numCardsExpected := 52
	numCardsActual := len(someDeck)
	if numCardsActual != numCardsExpected {
		t.Errorf("Expected %v elemments in new deck but found %v", numCardsExpected, numCardsActual)
	}

	expectedFirstElement := "Ace of Spades"
	actualFirstElement := someDeck[0]
	if actualFirstElement != expectedFirstElement {
		t.Errorf("Expected first element of new deck as \"%v\", but found \"%v\"", expectedFirstElement, actualFirstElement)
	}

	expectedLastElement := "King of Clubs"
	actualLastElement := someDeck[len(someDeck)-1]
	if expectedLastElement != actualLastElement {
		t.Errorf("Expected last element of new deck as: \"%v\" but found: \"%v\"", expectedLastElement, actualLastElement)
	}
}

func TestSaveToFileAndReadDeckFromFile(t *testing.T) {

	someDeck := newDeck()
	fileName := "_decktesting"

	// Remove the file, if exists already
	os.Remove(fileName)
	err := someDeck.saveToFile(fileName)
	if err == nil {
		deckFromFile := readDeckFromFile(fileName)
		if len(deckFromFile) != len(someDeck) {
			t.Errorf("File could not be read successfully from disk")
		}
	}
	// os.Remove(fileName)
}
