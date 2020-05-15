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

//DealerString func
func (h Hand) DealerString() string {
	return h[0].String() + ", --HIDDEN--"
}

//MinScore func
func (h Hand) MinScore() int {

	score := 0

	for _, c := range h {
		score += min(int(c.Rank), 10)
	}

	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	cards := deck.New(deck.Deck(3), deck.Shuffle)

	var card deck.Card

	var player, dealer Hand

	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = draw(cards)

			*hand = append(*hand, card)
		}
	}

	var input string

	for input != "s" {
		fmt.Println("Player:", player)
		fmt.Println("Dealer:", dealer.DealerString())
		fmt.Println("What wiil you do? (h)it, (s)tand")
		fmt.Scanf("%s\n", &input)

		switch input {
		case "h":
			card, cards = draw(cards)
			player = append(player, card)
		}
	}

	fmt.Println("--Final Hands--")
	fmt.Println("Player:", player, "\nScore:", player.MinScore())
	fmt.Println("Dealer:", dealer, "\nScore:", dealer.MinScore())
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}
