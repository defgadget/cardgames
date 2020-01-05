package deck

import (
	"sort"
	"time"
	"math/rand"
	"fmt"

)

type suit int
type value int

// These are the suits of a playing card
const (
	Heart suit = iota * 15
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
	if c.Value == Joker {
		return "Joker"
	}
	return fmt.Sprintf("%v of %v", c.Value, c.Suit)
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
// BySuit used to sort Deck by suit
type BySuit Deck
// ByValue used to sort Deck by value
type ByValue Deck

func (d Deck) Len() int           { return len(d) }
func (d Deck) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d Deck) Less(i, j int) bool { return int(d[i].Value) + int(d[i].Suit) < int(d[j].Value) + int(d[j].Suit) }

func (bs BySuit) Less(i, j int) bool { return int(bs[i].Suit) <  int(bs[j].Suit)}
func (bs BySuit) Len() int           { return len(bs) }
func (bs BySuit) Swap(i, j int)      { bs[i], bs[j] = bs[j], bs[i] }

func (bv ByValue) Less(i, j int) bool { return int(bv[i].Value) < int(bv[j].Value) }
func (bv ByValue) Len() int           { return len(bv) }
func (bv ByValue) Swap(i, j int)      { bv[i], bv[j] = bv[j], bv[i] }

// AddJokers will add Jokers to the Deck - Jokers have no suit and
func (d *Deck) AddJokers(qty int) {
	joker := Card{-1, Joker}
	for i := 0; i < qty; i++ {
		*d = append(*d, joker)
	}
}

// Draw returns one Card from front/top of the Deck
func (d *Deck) Draw() Card {
	card, deck := (*d)[0], (*d)[1:]
	*d = deck
	return card
}

// DrawCards returns number of cards (qty) from the front/top of the deck. Those cards are then removed from the deck
func (d *Deck) DrawCards(qty int) []Card {
	cardsDrew, deck := (*d)[:qty], (*d)[qty:]
	*d = deck
	return cardsDrew
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

// CustomSort takes a user-defined comparison function (i, j int) bool that will order the cards
func (d Deck) CustomSort(fn func(i, j int) bool) {
	sort.Slice(d, fn)
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

// Sort should sort the deck by suit and value
func (d Deck) Sort() {
	sort.Sort(ByValue(d))
	sort.Sort(BySuit(d))
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

// NewMultiple will return a deck containing the number (n) of decks combined
func NewMultiple(n int) Deck {
	deck := make(Deck, 0)
	for i := 0; i < n; i++ {
		deck = append(deck, New()...)
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
