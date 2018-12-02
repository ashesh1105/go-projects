package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// We define a new type called deck which will be slice of string
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// Two loops to construct the deck and return
	// Notice the use of underscore here instead of i, j. Whenever we declare and not use
	// a variable in GO, we can repalce them with underscore. Otherwise it will flag error!
	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

/**
Here, we are saying that any variable of type deck (d in this case) will have a method called print()
and when this method is called, it will be iterated over and every element of deck type, i.e., every string
in this variable will get printed along with its index in the slice (deck type). We are saying d is a variable of
type deck and iterate on it. Another important thing is here, variable of a type (d in this case) is defined
as first 1-2 chars of type, i.e., d is used for variable name of deck type.

GO is not an Object Oriented so we can get around that by defining a type of our own and then define as many
function we need it for any instance (variable) of that type to call on. These functions are called "Receivers" in GO
*/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// A function in GO can return multiple values. Here's an example
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Notice the Join function from strings package. It concatenates slice of string elements
// with a given delimiter
func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

// This function needs converting a slice of string (our deck) to byte slice ([]byte)
// Notice how we convert our deck type back to slice of string and then to string and
// then finally to a byte slice ([]byte)! toString on deck is defined above.
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func readDeckFromFile(fileName string) deck {
	byteSlice, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(byteSlice), ","))
}

// Iterate through slice, for each element generate a random number between 0 and len-1
// and swap this element with element at the new number (index) generated.
func (d deck) shuffle() {

	// Below code allows a unique seed to be passed to random function, so result is really random
	// source is passed as seed. Unique source is comuted each time by passing unique Int64 value
	// as computed by time.Now().UnixNano() function.
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// Very interesting way to swap slice elements at two positions!
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
