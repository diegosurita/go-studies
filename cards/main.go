package main

import "fmt"

func main() {
	// We can initialize a variable with a value.
	// var card string = "Ace of Spades"

	// Or just initialize a variable without a value.
	// var card string

	// But the most commont way to initialize a variable is:
	card := "Ace of Spades"
	// Go strongly infer the type of the variable as string.

	// We can reassign a variable.
	card = "Five of Diamonds"

	fmt.Println(card)
}
