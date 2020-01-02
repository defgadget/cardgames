package deck

import (
	"fmt"
	"testing"
)

func TestAddJokers(t *testing.T) {
	deck := New()
	qty := 2
	deck.AddJokers(qty)
	if len(deck) != 52 + qty {
		t.Error("The deck is not the correct length", len(deck))
	}
	jokersFound := 0
	for _, card := range deck {
		if card.Value == Joker {
			jokersFound++
		}
	}
	if jokersFound != qty {
		t.Error("All the jokers weren't added", jokersFound)
	}
}
func TestNewDeckSize(t *testing.T){
	expected := 52
	actual := len(New())
	if actual != expected {
		t.Errorf("Test Failed - Expected %v and received %v", expected, actual)
	}
}

func TestNewDeckOrder(t *testing.T){
	suits := [...]string{
		"Hearts",
		"Diamonds",
		"Clubs",
		"Spades"}

	values := [...]string{
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
	index := 0
	deck := New()
	for _, s := range suits {
		for _, v := range values {
			card := deck[index]
			if fmt.Sprint(card.Suit) != s && fmt.Sprint(card.Value) != v {
				t.Errorf("Expected: %v-%v\nReceived %v-%v", card.Value, card.Suit, v, s)
			}
			index++
		}
	}
}

func TestRefreshDeck(t *testing.T){
	deck := New()
	freshdeck := New()
	deck.Remove(2,3,4,5)

	deck.Refresh()
	if len(deck) != len(freshdeck) {
		t.Error("The decks should be the same length")
	}
	for i, card := range freshdeck {
		if card.Value != deck[i].Value && card.Suit != deck[i].Suit {
			t.Error("The decks should match order", card, deck[i])
		}
	}
}

func TestRemoveTwos(t *testing.T){
	deck := New()
	deck.Remove(2)
	for _, card := range deck {
		if card.Value == 2 {
			t.Error("Twos should have been removed", len(deck))
		}
	}
}

func TestRemoveMany(t *testing.T){
	deck := New()
	deck.Remove(2,3,4,5,6)
	for _, card := range deck {
		if card.Value == 2 || card.Value == 3 || card.Value == 4 || card.Value == 5 || card.Value == 6{
			t.Error("Card should have been removed:", card.Value)
		}
	}
}

func TestNewShuffle(t *testing.T){
	// Not sure how to effectively test the shuffle
	numOfDifferences := 0
	new := New()
	deck := New()
	deck.Shuffle()
	for i, card := range new {
		if card.Value != deck[i].Value || card.Suit != deck[i].Suit {
			numOfDifferences++
		}
	}
	if numOfDifferences < 10 {
		t.Error("Less than 10% variance between shuffled and unshuffled")
	}
}

func TestRemovedCardsShuffle(t *testing.T){
	// Not sure how to effectively test the shuffle
	numOfDifferences := 0
	new := New()
	deck := New()
	new.Remove(2,3,4,5)
	deck.Remove(2,3,4,5)
	deck.Shuffle()
	for i, card := range new {
		if card.Value != deck[i].Value || card.Suit != deck[i].Suit {
			numOfDifferences++
		}
	}
	if numOfDifferences < 10 {
		t.Error("Less than 10% variance between shuffled and unshuffled")
	}
}

func TestWithJokersShuffle(t *testing.T){
	// Not sure how to effectively test the shuffle
	numOfDifferences := 0
	new := New()
	deck := New()
	new.AddJokers(2)
	deck.AddJokers(2)
	deck.Shuffle()
	for i, card := range new {
		if card.Value != deck[i].Value || card.Suit != deck[i].Suit {
			numOfDifferences++
		}
	}
	if numOfDifferences < 10 {
		t.Error("Less than 10% variance between shuffled and unshuffled")
	}
}