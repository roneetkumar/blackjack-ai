package main

import (
	"fmt"

	"github.com/roneetkumar/blackjack-ai/blackjack"
)

func main() {
	opts := blackjack.Options{
		Decks:           3,
		Hands:           2,
		BlackJackPayout: 1.5,
	}

	game := blackjack.New(opts)

	winnings := game.Play(blackjack.HumanAI())

	fmt.Println(winnings)
}
