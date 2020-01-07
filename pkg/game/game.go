package game

import (
	"bufio"
	"fmt"
	"strings"
	"os"

	"github.com/defgadget/cardgames/pkg/deck"
)
type hand []deck.Card

type Player struct {
	Name string
	Hand hand
	IsDealer bool
}

func (h hand) String() string {
	str := "******HAND******\n"
	for _, card := range h {
		str += fmt.Sprintf("%v\n", card)
	}
	str += "****************"
	return str
}
type Players []Player

// ScoreHand will be implemented by the individual games, based on their rules.
// it will take a PlayerHand and will Score the hand based on the cards contained
type ScoreHand func (hand Player)

func Deal(dealer *Player, players Players, d *deck.Deck, numCards int) {
	for i := 0; i < numCards; i++ {
		for j := 0; j < len(players); j++ {
			newCard := d.Draw()
			fmt.Printf("%v dealt a %v\n", players[j].Name, newCard)
			players[j].Hand = append(players[j].Hand, newCard)
		}
		newCard := d.Draw()
		if i == 0 {
		    fmt.Printf("%v dealt a ******n", dealer.Name)
		}
		dealer.Hand = append(dealer.Hand, newCard)

	}
}

func GetChoice(msg string) string {
    trim := strings.TrimSpace
    lower := strings.ToLower
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(msg)
    text, _ := reader.ReadString('\n')
    text = lower(trim(text))
    return text
}
