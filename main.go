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

//Score func
func (h Hand) Score() int {
	minScore := h.MinScore()

	if minScore > 11 {
		return minScore
	}

	for _, c := range h {
		if c.Rank == deck.Ace {
			// ace is currently worth 1
			return minScore + 10
		}
	}
	return minScore
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

	for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
		card, cards = draw(cards)
		dealer = append(dealer, card)
	}

	pScore, dScore := player.Score(), dealer.Score()

	fmt.Println("--Final Hands--")
	fmt.Println("Player:", player, "\nScore:", pScore)
	fmt.Println("Dealer:", dealer, "\nScore:", dScore)

	switch {
	case pScore > 21:
		fmt.Println("You Busted")
	case dScore > 21:
		fmt.Println("Dealer Busted")
	case pScore > dScore:
		fmt.Println("You Wins")
	case pScore < dScore:
		fmt.Println("You Lose")
	case pScore == dScore:
		fmt.Println("Draw")
	}
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}
