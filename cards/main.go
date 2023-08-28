package main

func main() {
	// We can initialize a variable with a value.
	// var card string = "Ace of Spades"

	// Or just initialize a variable without a value.
	// var card string

	// But the most commont way to initialize a variable is:
	// card := newCard()
	// Go strongly infer the type of the variable as string, because newCard returns a string.

	// After the variable to be initialized, we can reassign a variable.
	// card = "Five of Diamonds"

	// We can also create a slice of strings.
	// Slices are like arrays but with a dynamic size. Arrays have a fixed size.
	cards := newDeck()

	// We can append a new element to the slice.
	// cards = append(cards, "Six of Spades")

	// We can iterate over the slice.
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// Let's print the cards using the print function created on deck.go.
	// cards.print()

	cards.saveToFile("my_cards")
}
