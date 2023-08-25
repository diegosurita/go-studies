package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// Creates a new deck of cards but it's not attached to deck type because it's not receiving any deck.
func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Create a function that print all the cards of the deck.
// Every function that receives a variable of type deck, is like a method in OOP approach.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
