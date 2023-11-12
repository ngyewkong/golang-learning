package main

import "fmt"

func main() {
	// variable assignment - long formal way
	// var variableName variableType = value
	// variableType bool string int float64
	// var card string = "Ace of Spades"
	// short hand as go compiler is able to infer
	// variableName := value (only for initialising new variable)
	card := "Ace of Spades"
	// reassignment
	card = "Five of Diamonds"
	fmt.Println(card)
}
