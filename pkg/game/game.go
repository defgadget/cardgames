package game

import "github.com/defgadget/cardgames/pkg/deck"

type Hand struct {
	Cards []Card
	Score int
}

type Player struct {
	Name string
	PlayerHand Hand
}

Players []Player

// ScoreHand will be implemented by the individual games, based on their rules.
// it will take a PlayerHand and will Score the hand based on the cards contained
type ScoreHand func (hand Hand)