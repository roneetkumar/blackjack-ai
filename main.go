package main

import (
	"fmt"
	"strings"

	deck "github.com/roneetkumar/shuffle"
)

//Hand type
type Hand []deck.Card

func (h Hand) String() string {
	str := make([]string, len(h))

	for i := range h {
		str[i] = h[i].String()
	}
	return strings.Join(str, ", ")
}

func main() {
	cards := deck.New(deck.Deck(3), deck.Shuffle)

	var card deck.Card

	var player, dealer Hand

	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = cards[0], cards[1:]

			*hand = append(*hand, card)
		}
	}

	fmt.Println("Player:", player)
	fmt.Println("Dealer:", dealer)

}
