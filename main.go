package main

import "fmt"

func main() {
	cards := deck{"Five of Diamonds", newCard()}
	// fmt.Println(card)
	cards = append(cards, "six of Spades")
	for i, card := range cards {
		fmt.Println(i, card)

	}
}

func newCard() string {
	return "Ace of Spades"
}
