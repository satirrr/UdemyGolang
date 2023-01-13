package main

import (
	"fmt"
)

func main() {
	deck := newDeck()
	myHand := Deck{}
	fmt.Println("deck : ", deck)
	fmt.Println("### draw one card from the old deck ###")
	myHand, deck = drawAHandfulCard(myHand, deck, 1)
	fmt.Println("deck : ", deck)
	fmt.Println("my hand : ", myHand)

	fmt.Println("### draw another card from old deck to my hand ###")
	myHand, deck = drawAHandfulCard(myHand, deck, 1)

	fmt.Println("deck : ", deck)
	fmt.Println("my hand : ", myHand)

	fmt.Println("#### NEW GAME ####")
	deck = newDeck()
	myHand = nil
	fmt.Println("deck : ", deck)

	numberOfCard := 3
	myHand, deck = drawAHandfulCard(myHand, deck, 3)
	fmt.Println("### draw a handful of ", numberOfCard, " cards from the deck ###")
	fmt.Println("my hand : ", myHand)
	fmt.Println("deck : ", deck)

}
