package main

import "github.com/roneetkumar/blackjack-ai/blackjack"

// //Hand type
// type Hand []deck.Card

// func (h Hand) String() string {
// 	str := make([]string, len(h))

// 	for i := range h {
// 		str[i] = h[i].String()
// 	}
// 	return strings.Join(str, ", ")
// }

// //DealerString func
// func (h Hand) DealerString() string {
// 	return h[0].String() + ", --HIDDEN--"
// }

// //Score func
// func (h Hand) Score() int {
// 	minScore := h.MinScore()

// 	if minScore > 11 {
// 		return minScore
// 	}

// 	for _, c := range h {
// 		if c.Rank == deck.Ace {
// 			// ace is currently worth 1
// 			return minScore + 10
// 		}
// 	}
// 	return minScore
// }

// //MinScore func
// func (h Hand) MinScore() int {

// 	score := 0

// 	for _, c := range h {
// 		score += min(int(c.Rank), 10)
// 	}

// 	return score
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// //Shuffle func
// func Shuffle(gs GameState) GameState {
// 	state := clone(gs)
// 	state.Deck = deck.New(deck.Deck(3), deck.Shuffle)
// 	return state
// }

// //Deal func
// func Deal(gs GameState) GameState {
// 	state := clone(gs)

// 	state.Player = make(Hand, 0, 5)
// 	state.Dealer = make(Hand, 0, 5)

// 	var card deck.Card

// 	for i := 0; i < 2; i++ {
// 		card, state.Deck = draw(state.Deck)
// 		state.Player = append(state.Player, card)
// 		card, state.Deck = draw(state.Deck)
// 		state.Dealer = append(state.Dealer, card)
// 	}

// 	state.State = StatePlayerTurn

// 	return state
// }

// // Hit func
// func Hit(gs GameState) GameState {
// 	state := clone(gs)

// 	hand := state.CurrentPlayer()

// 	var card deck.Card

// 	card, state.Deck = draw(state.Deck)
// 	*hand = append(*hand, card)

// 	if hand.Score() > 21 {
// 		return Stand(state)
// 	}

// 	return state
// }

// //Stand func
// func Stand(gs GameState) GameState {
// 	state := clone(gs)
// 	state.State++
// 	return state
// }

// //EndHand func
// func EndHand(gs GameState) GameState {

// 	state := clone(gs)

// 	pScore, dScore := state.Player.Score(), state.Dealer.Score()

// 	fmt.Println("--Final Hands--")
// 	fmt.Println("Player:", state.Player, "\nScore:", pScore)
// 	fmt.Println("Dealer:", state.Dealer, "\nScore:", dScore)

// 	switch {
// 	case pScore > 21:
// 		fmt.Println("You Busted")
// 	case dScore > 21:
// 		fmt.Println("Dealer Busted")
// 	case pScore > dScore:
// 		fmt.Println("You Wins")
// 	case pScore < dScore:
// 		fmt.Println("You Lose")
// 	case pScore == dScore:
// 		fmt.Println("Draw")
// 	}

// 	fmt.Println()

// 	state.Player = nil
// 	state.Dealer = nil

// 	return state
// }

func main() {

	game := blackjack.New()

	game.Play(blackjack.HumanAI())

	// var gs GameState
	// gs = Shuffle(gs)

	// for i := 0; i < 10; i++ {
	// 	gs = Deal(gs)

	// 	var input string

	// 	for gs.State == StatePlayerTurn {
	// 		fmt.Println("Player:", gs.Player)
	// 		fmt.Println("Dealer:", gs.Dealer.DealerString())
	// 		fmt.Println("What wiil you do? (h)it, (s)tand")
	// 		fmt.Scanf("%s\n", &input)

	// 		switch input {
	// 		case "h":
	// 			gs = Hit(gs)
	// 		case "s":
	// 			gs = Stand(gs)
	// 		default:
	// 			fmt.Println("Invalid option", input)
	// 		}
	// 	}

	// 	for gs.State == StateDealerTurn {
	// 		if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
	// 			gs = Hit(gs)
	// 		} else {
	// 			gs = Stand(gs)
	// 		}
	// 	}

	// 	gs = EndHand(gs)
	// }
}

// func draw(cards []deck.Card) (deck.Card, []deck.Card) {
// 	return cards[0], cards[1:]
// }

// // State type
// type State int8

// const (
// 	StatePlayerTurn State = iota
// 	StateDealerTurn
// 	StateHandOver
// )

// //GameState struct
// type GameState struct {
// 	Deck   []deck.Card
// 	State  State
// 	Player Hand
// 	Dealer Hand
// }

// func (gs *GameState) CurrentPlayer() *Hand {
// 	switch gs.State {
// 	case StatePlayerTurn:
// 		return &gs.Player
// 	case StateDealerTurn:
// 		return &gs.Dealer

// 	default:
// 		panic("It isn't currently any player's turn")

// 	}
// }

// func clone(gs GameState) GameState {
// 	state := GameState{
// 		Deck:   make([]deck.Card, len(gs.Deck)),
// 		State:  gs.State,
// 		Player: make(Hand, len(gs.Player)),
// 		Dealer: make(Hand, len(gs.Dealer)),
// 	}

// 	copy(state.Deck, gs.Deck)
// 	copy(state.Player, gs.Player)
// 	copy(state.Dealer, gs.Dealer)

// 	return state
// }
