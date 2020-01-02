package deck

import (
	"time"
	"math/rand"
	"fmt"

)

type suit int
type value int

// These are the suits of a playing card
const (
	Heart suit = iota
	Diamond
	Club
	Spade
)
// These are the values of a playing card
const (
	Joker value = iota
	Ace 
	One 
	Two 
	Three 
	Four 
	Five 
	Six 
	Seven 
	Eight 
	Nine 
	Ten 
	Jack 
	Queen 
	King 
)
// Card represents a card in a normal deck of playing cards. Containing Suit (Club, Spade, Heart, Diamond)
// and a Value (Ace, 2 - 10, J, Q, K)
type Card struct {
	Suit suit
	Value value
}

func (c Card) String() string {
	return fmt.Sprintln(c.Value, "of", c.Suit)
}

func (s suit) String() string {
	suits := [...]string{
		"Hearts",
		"Diamonds",
		"Clubs",
		"Spades"}
	if s < Heart || s > Spade {
		return "Unknown"
	}
	return suits[s]
}

func (v value) String() string {
	values := [...]string{
		"Joker",
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King"}

		if v < Joker || v > King {
			return "Unknown"
		}
		return values[v]
}

// Deck is a collection of playing cards
type Deck []Card

// AddJokers will add Jokers to the Deck - Jokers have no suit and
func (d *Deck) AddJokers(qty int) {
	joker := Card{-1, Joker}
	for i := 0; i < qty; i++ {
		*d = append(*d, joker)
	}
}

// Refresh will return the deck to an unshuffled "New" state.
func (d *Deck) Refresh() {
	*d = New()
}

// Remove will remove cards from the deck. Remove does not assume of full deck
// If you have already removed cards those removals will remain for subsequent calls to Remove
// If you only whish for the current call to Remove to be applied you should call Refresh first. 
func (d *Deck) Remove(cards ...value) {
	newdeck := make(Deck, len(*d) - len(cards)*4)
	insert := 0
	removeCard := false
	for _, card := range *d {
		for _, cv := range cards {
			if card.Value == cv {
				removeCard = true
			}
		}
		if !removeCard {
			newdeck[insert] = card
			insert++
		}
		removeCard = false
	}
	*d = newdeck
}

// Shuffle will "shuffle" the cards in the deck to a random order
// the random Seed takes the (Minute * Hour * Day)/Second to try
// to ensure a non-deterministic order
func (d Deck) Shuffle() {
	seed := (time.Now().Minute() * time.Now().Hour() * time.Now().Day()) / time.Now().Second()
	rand.Seed(int64(seed))
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

// New returns a fresh deck which is unshuffled
func New() Deck {
	suits := [...]string{
		"Hearts",
		"Diamonds",
		"Clubs",
		"Spades"}

		values := [...]string{
		"Joker",
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King"}
	deck := make(Deck, 52)
	insert := 0
	for _, s := range suits {
		for _, v := range values {
			if v != "Joker" {
				deck[insert] = Card{getSuitValue(s), getValuesValue(v)}
				insert++
			}
		}
	}
	return deck
}

func getSuitValue(s string) suit {
	suits := [...]string{
		"Hearts",
		"Diamonds",
		"Clubs",
		"Spades"}
	for i, ss := range suits {
		if s == ss {
			return suit(i)
		}
	}
	return -1
}

func getValuesValue(s string) value {
	values := [...]string{
		"Joker",
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King"}
	for i, v := range values {
		if s == v {
			return value(i)
		}
	}
	return -1
}