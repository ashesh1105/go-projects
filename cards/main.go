package main

import (
	"fmt"
)

func main() {

	// Commented line below is one way of declaring and initializing a variable
	// Non-commented one is another way. Remember := is needed only once
	// var card string = "Ace of Spades"
	card := "Ace of Spades"

	// Don't do := anymore since it is allowed only once for a variable
	card = "Some other card Type"

	// Call a function to initialize a variable
	anotherCard := newCard()

	fmt.Println(card)
	fmt.Println("\nAnother card:")
	fmt.Println(anotherCard)

	// This is how you declare a slice. [] string can be read as "slice of string"
	cards := []string{"King of Hearts", newCard()}
	fmt.Println("\ncards slice:")
	fmt.Println(cards)

	// You can append to a slice. Remember, GO never updates an existing slice, it generates a new slice
	cards = append(cards, "Jack of Spades")
	fmt.Println(append(cards, " : <= new cards slice after a string is appended to it"))

	// This is how you iterate through a slice
	fmt.Println("\nIterating though cards slice:")
	for i, card := range cards {
		fmt.Println(i, card)
	}

	myCards := deck{"Six of Spades", "Nine of Hearts", newCard()}
	fmt.Println("\nPrinting myCards slice using a function defined on our custom type:")

	// print() is our function defined in cards.go
	// and this can be called on all variables of type "deck"
	myCards.print()

	anotherDeck := newDeck()

	fmt.Println("\nPrinting custom type deck using a Receiver function as print() " +
		"defined in the deck type :")
	anotherDeck.print()

	// Cut a deal (hand). Notice how this function returns multiple values!
	hand, remainingCard := deal(anotherDeck, 5)
	fmt.Println("\nPrinting a hand of 5 cards:")
	hand.print()
	fmt.Println("\nPrinting remaining cards in the deck:")
	remainingCard.print()

	// Byte Slice: String to Bytes Slice conversion
	str := "Hi there!"
	fmt.Println("\nByte Slice out of string: \"" + str + "\":")
	fmt.Println([]byte(str))

	// Join the elements of a slice of strings with a delimiter
	fmt.Println("\nCalling custom toString on hand (deck):")
	fmt.Println(hand.toString())

	// Save hand deck to a file by calling our custom function
	firstHandFileName := "firstHand.txt"
	err := hand.saveToFile(firstHandFileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("\ndeck \"hand\" written to file:", firstHandFileName)
	}

	// Let's save entire deck as file too
	deckFileName := "my_deck.txt"
	anotherDeck.saveToFile(deckFileName)
	fmt.Println("\nEntire Deck written to file:", deckFileName)

	// Read Deck from File and print using our receiver function print() in it
	fmt.Println("\nReading Deck from File:", deckFileName, ":")
	readDeckFromFile(deckFileName).print()

	// Shuffle the deck using our receiver function shuffle()
	anotherDeck.shuffle()
	fmt.Println("\nShuffled the deck using our receiver function on deck type:")
	anotherDeck.print()
}

// Notice the return type being spefified here!
func newCard() string {
	return "Five of Diamonds"
}
