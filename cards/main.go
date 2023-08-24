package main

import "fmt"

func main() {
	// We can initialize a variable with a value.
	// var card string = "Ace of Spades"

	// Or just initialize a variable without a value.
	// var card string

	// But the most commont way to initialize a variable is:
	card := newCard()
	// Go strongly infer the type of the variable.

	// After the variable to be initialized, we can reassign a variable.
	// card = "Five of Diamonds"

	fmt.Println(card)
}

// Functions have always to return a type.
func newCard() string {
	return "Five of Diamonds"
}
