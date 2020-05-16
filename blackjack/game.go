package blackjack

import (
	"fmt"

	deck "github.com/roneetkumar/shuffle"
)

const (
	statePlayerTurn state = iota
	stateDealerTurn
	stateHandOver
)

// State type
type state int8

//New func
func New() Game {
	return Game{
		state:    statePlayerTurn,
		dealerAI: dealerAI{},
		balance:  0,
	}
}

//Game struct
type Game struct {
	//unexported fields
	deck     []deck.Card
	state    state
	player   []deck.Card
	dealer   []deck.Card
	dealerAI AI
	balance  int
}

func (g *Game) currentHand() *[]deck.Card {
	switch g.state {
	case statePlayerTurn:
		return &g.player
	case stateDealerTurn:
		return &g.dealer

	default:
		panic("It isn't currently any player's turn")

	}
}

//Deal func
func deal(g *Game) {

	g.player = make([]deck.Card, 0, 5)
	g.dealer = make([]deck.Card, 0, 5)

	var card deck.Card

	for i := 0; i < 2; i++ {
		card, g.deck = draw(g.deck)
		g.player = append(g.player, card)
		card, g.deck = draw(g.deck)
		g.dealer = append(g.dealer, card)
	}

	g.state = statePlayerTurn

}

//Play func
func (g *Game) Play(ai AI) int {
	g.deck = deck.New(deck.Deck(3), deck.Shuffle)

	for i := 0; i < 10; i++ {
		deal(g)

		for g.state == statePlayerTurn {
			hand := make([]deck.Card, len(g.player))
			copy(hand, g.player)
			move := ai.Play(hand, g.dealer[0])
			move(g)
		}

		for g.state == stateDealerTurn {
			hand := make([]deck.Card, len(g.player))
			copy(hand, g.dealer)
			move := g.dealerAI.Play(hand, g.dealer[0])
			move(g)
		}

		endHand(g, ai)
	}
	return 0
}

//Move func
type Move func(*Game)

// MoveHit func
func MoveHit(g *Game) {
	hand := g.currentHand()

	var card deck.Card

	card, g.deck = draw(g.deck)
	*hand = append(*hand, card)

	if Score(*hand...) > 21 {
		MoveStand(g)
	}
}

//MoveStand func
func MoveStand(g *Game) {
	g.state++
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func endHand(g *Game, ai AI) {

	pScore, dScore := Score(g.player...), Score(g.dealer...)
	//TODO: keep track of win/lose
	switch {
	case pScore > 21:
		fmt.Println("You Busted")
		g.balance--
	case dScore > 21:
		fmt.Println("Dealer Busted")
		g.balance++
	case pScore > dScore:
		fmt.Println("You Wins")
		g.balance++
	case pScore < dScore:
		fmt.Println("You Lose")
		g.balance--
	case pScore == dScore:
		fmt.Println("Draw")
	}
	fmt.Println()

	ai.Results([][]deck.Card{g.player}, g.dealer)

	g.player = nil
	g.dealer = nil
}

//Score func
func Score(hand ...deck.Card) int {
	minScore := minScore(hand...)

	if minScore > 11 {
		return minScore
	}

	for _, c := range hand {
		if c.Rank == deck.Ace {
			// ace is currently worth 1
			return minScore + 10
		}
	}
	return minScore
}

//Soft func
func Soft(hand ...deck.Card) bool {
	minScore := minScore(hand...)
	score := Score(hand...)
	return minScore != score
}

func minScore(hand ...deck.Card) int {
	score := 0

	for _, c := range hand {
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
