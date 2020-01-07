package blackjack

import (
	"fmt"

	"github.com/defgadget/cardgames/pkg/deck"
	"github.com/defgadget/cardgames/pkg/game"
)

type BlackJack struct {
    Players game.Players
    Deck deck.Deck
    Dealer game.Player
}

type Options struct {
    NumPlayers int
    NumDecks int
    NumHands int
}

func New(opt Options) BlackJack {
    players := setupPlayers(opt.NumPlayers)
    d := deck.NewMultiple(opt.NumDecks)
    d.Shuffle()
    dealer := game.Player{Name: "Dealer", Hand: make([]deck.Card, 0)}
    return BlackJack{players, d, dealer}
}

func scoreHand(cards []deck.Card) int {
    score := 0
    for _, card := range cards {
        switch {
        case card.Value >= deck.Ten :
            score += 10
        case card.Value > deck.Ace && card.Value < deck.Ten :
            score += int(card.Value)
        case card.Value == deck.Ace :
            score += 1
        }
    }
    if softAce(cards) {
        score += 10
    }
    return score
}

func scoreRound(dealer game.Player, players game.Players) {
    // See who won
    dealerScore := scoreHand(dealer.Hand)
    busted := false
    for _, player := range players {
        playerScore := scoreHand(player.Hand)
        if scoreHand(player.Hand) > 21 {
            busted = true
        }
        switch {
        case playerScore == dealerScore:
            fmt.Println("Push")
            fmt.Println("Dealer Score:", dealerScore)
            fmt.Printf("%v Score: %v\n", player.Name, playerScore)
        case busted || playerScore < dealerScore && dealerScore <= 21:
            fmt.Println("You Lost")
            fmt.Println("Dealer Score:", dealerScore)
            fmt.Printf("%v Score: %v\n", player.Name, playerScore)
        case playerScore > dealerScore || dealerScore > 21:
            fmt.Println("You Won!")
            fmt.Println("Dealer Score:", dealerScore)
            fmt.Printf("%v Score: %v\n", player.Name, playerScore)
        }
    }
}

func setupPlayers(numPlayers int) game.Players {
    players := make(game.Players, numPlayers)
    player := game.Player{}
    for i := 0; i < numPlayers; i++ {
        player.Name = fmt.Sprintf("Player %v", i + 1)
        player.Hand = make([]deck.Card, 0)
        players[i] = player
    }
    return players
}

func softAce(hand []deck.Card) bool {
    hasAce := false
    for _, card := range hand {
        if card.Value == deck.Ace {
            hasAce = true
        }
    }
    if hasAce && scoreHand(hand) + 10 <= 21 {
        return true
    }
    return false
}

func checkNoChoice(score int) bool {
    finished := false
    if score == 21 {
        fmt.Println("You got 21!!")
        fmt.Println()
        finished = true
    }
    if score > 21 {
        fmt.Println("Bust", score)
        fmt.Println()
        finished = true
    }
    return finished
}

func playRound(dealer *game.Player, players game.Players, d *deck.Deck) {
    playerFin := false
    playerScore := 0
    choice := ""
    for i, player := range players {
        fmt.Printf("--- %v ---\n----------------\n", player.Name)
        playerScore = scoreHand(players[i].Hand)
        for !playerFin {
            fmt.Printf("%v\n\n", players[i].Hand)
            if checkNoChoice(playerScore) {
                break
            }
            choice = game.GetChoice("Would you like to Hit or Stay? ")
            switch choice {
            case "hit":
                newCard := d.Draw()
                players[i].Hand = append(players[i].Hand, newCard)
                fmt.Printf("\n**Hit** %v **Hit**\n\n", newCard)
            case "stay":
                fmt.Printf("\n**Stay** %v **Stay**\n\n", playerScore)
                playerFin = true
            default:
                fmt.Printf("\nYou can only choose Hit or Stay\n\n")
            }
            playerScore = scoreHand(players[i].Hand)
        }
        playerFin = false
    }
    dealerScore := scoreHand(dealer.Hand)
    dealerStay := false
    for !dealerStay {
        fmt.Printf("--- %v ---\n----------------\n", dealer.Name)
        fmt.Printf("%v\n\n", dealer.Hand)
        if dealerScore < 17 || softAce(dealer.Hand) && dealerScore == 17 {
            choice = "hit"
        } else {
            choice  = "stay"
        }
        switch choice {
        case "hit":
            newCard := d.Draw()
            dealer.Hand = append(dealer.Hand, newCard)
            fmt.Printf("\n**Hit** %v **Hit**\n\n", newCard)
        case "stay":
            fmt.Printf("\n**Stay** %v **Stay**\n\n", dealerScore)
            dealerStay = true
        default:
            fmt.Printf("\nYou can only choose Hit or Stay\n\n")
        }
        dealerScore = scoreHand(dealer.Hand)
        if checkNoChoice(dealerScore) {
            break
        }
    }
}

func (bj *BlackJack) Play() {
    playing := true
    for playing {
        game.Deal(&bj.Dealer, bj.Players, &bj.Deck, 2)
        playRound(&bj.Dealer, bj.Players, &bj.Deck)
        scoreRound(bj.Dealer, bj.Players)
        for gotAnswer := false; !gotAnswer; {
            another := game.GetChoice("Would you like to play another round? (Yes/No): ")
            switch another {
            case "yes":
                playing = true
                gotAnswer = true
            case "no":
                playing = false
                gotAnswer = true
            default:
                fmt.Println("I didn't understand")
            }
        }
    }
}
