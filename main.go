package main

import (
	"fmt"

	"github.com/roneetkumar/blackjack-ai/blackjack"
)

func main() {

	game := blackjack.New()

	winnings := game.Play(blackjack.HumanAI())

	fmt.Println(winnings)

}
