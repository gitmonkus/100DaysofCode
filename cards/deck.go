package main

import "fmt"

// Create a new type of deck
// which is a slice of strings
// deck is equal to a slice of string, like an alias
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// This is a receiver function
// (d deck) is a receiver.  d is of type deck
// for index (0,1,2,3) item number
// loop through the cards slice passed in and save in card variable
// print out the index number and each item in card from the cards slice
// Anything of type deck can call this function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
