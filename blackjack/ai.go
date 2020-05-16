package blackjack

import (
	"fmt"

	deck "github.com/roneetkumar/shuffle"
)

//AI interface
type AI interface {
	Bet(shuffled bool) int
	Play(hand []deck.Card, dealer deck.Card) Move
	Results(hand [][]deck.Card, dealer []deck.Card)
}

// HumanAI struct
type dealerAI struct {
}

//Bet func
func (ai dealerAI) Bet(shuffled bool) int {
	// nothing
	return 1
}

func (ai dealerAI) Play(hand []deck.Card, dealer deck.Card) Move {
	dScore := Score(hand...)
	if dScore <= 16 || dScore == 17 && Soft(hand...) {
		return MoveHit
	}
	return MoveStand
}

// Results func
func (ai dealerAI) Results(hand [][]deck.Card, dealer []deck.Card) {
	// nothing
}

//HumanAI func
func HumanAI() AI {
	return humanAI{}
}

// HumanAI struct
type humanAI struct {
}

//Bet func
func (ai humanAI) Bet(shuffled bool) int {

	if shuffled {
		fmt.Println("The deck was just shuffled!")
	}

	fmt.Println("What would you like to bet?")
	var bet int
	fmt.Scanf("%d\n", &bet)
	return bet
}

// Play func
func (ai humanAI) Play(hand []deck.Card, dealer deck.Card) Move {

	for {
		fmt.Println("Player:", hand)
		fmt.Println("Dealer:", dealer)
		fmt.Println("What wiil you do? (h)it, (s)tand")
		var input string
		fmt.Scanf("%s\n", &input)

		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		default:
			fmt.Println("Invalid option", input)
		}
	}
}

// Results func
func (ai humanAI) Results(hand [][]deck.Card, dealer []deck.Card) {
	fmt.Println("--Final Hands--")
	fmt.Println("Player:", hand)
	fmt.Println("Dealer:", dealer)
}
