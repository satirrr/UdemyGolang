package main

import (
	"math/rand"
	"time"
)

//Deck create a new type of 'Deck'
//which is a slice of string
type Deck []string

func newDeck() Deck {
	deck := Deck{}
	return addCardToDeck(deck).shuffleCards()
}

func drawAHandfulCard(hand Deck, deck Deck, numberOfCard int) (Deck, Deck) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numberOfCard; i++ {
		coba := len(deck) - 1
		randomIndex := rand.Intn(coba)
		randomCard := deck[randomIndex]
		deck = removeCardFromDeck1(deck, randomCard)
		hand = append(hand, randomCard)
	}
	return hand, deck
}

func addCardToDeck(deck Deck) Deck {
	cardNumbers := []string{"Ace", "King"}
	cardSymbols := []string{"Spade", "Clover"}

	for _, number := range cardNumbers {
		for _, symbol := range cardSymbols {
			deck = append(deck, number+" of "+symbol)
		}
	}
	return deck
}

func removeCardFromDeck1(deck Deck, randomCard string) Deck {
	thisDeck := Deck{}
	for _, card := range deck {
		if card != randomCard {
			thisDeck = append(thisDeck, card)
		}
	}
	return thisDeck
}

func (cards Deck) shuffleCards() Deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	return cards
}
