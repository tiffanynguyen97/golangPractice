package main

import "fmt"

func main() {
	card := []string{"Ace of Diamonds", "Queen of Hearts", "King of Clubs", newCard()}

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
