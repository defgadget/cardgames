package blackjack

import (
	"fmt"

	"github.com/defgadget/cardgames/pkg/deck"
	"github.com/defgadget/cardgames/pkg/game"
)
func scoreHand(cards []deck.Card) int {
    score := 0
    hasAce := false
    for _, card := range cards {
        switch {
        case card.Value >= deck.Ten :
            score += 10
        case card.Value > deck.Ace && card.Value < deck.Ten :
            score += int(card.Value)
        case card.Value == deck.Ace :
            score += 1
            hasAce = true
        }
    }
    if hasAce && score + 10 <= 21 {
        score += 10
    }
    return score
}

func scoreRound(players game.Players) {
    // See who won
    dealer := players[len(players)-1]
    dealerScore := scoreHand(dealer.Hand)
    busted := false
    for _, player := range players {
        playerScore := scoreHand(player.Hand)
        if player.IsDealer {
            break
        }
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
        if i + 1 == numPlayers {
            player.Name = " Dealer "
            player.IsDealer = true
        }
        player.Hand = make([]deck.Card, 0)
        players[i] = player
    }
    return players
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

func playRound(players game.Players, d *deck.Deck) {
    dealer := players[len(players)-1]
    playerBusted := false
    playerFin := false
    playerScore := 0
    dealerScore := scoreHand(dealer.Hand)
    choice := ""
    for i, player := range players {
        fmt.Printf("--- %v ---\n----------------\n", player.Name)
        playerScore = scoreHand(players[i].Hand)
        if player.IsDealer && playerBusted {
            fmt.Printf("%v\n\n", players[i].Hand)
            break
        }
        for !playerFin {
            fmt.Printf("%v\n\n", players[i].Hand)
            if checkNoChoice(playerScore) {
                break
            }
            if !player.IsDealer {
                choice = game.GetChoice("Would you like to Hit or Stay? ")
            }
            if player.IsDealer {
                dealerScore  = playerScore
                if dealerScore < 17 {
                    choice = "hit"
                } else {
                    choice  = "stay"
                }
            }
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
            if player.IsDealer {
                dealerScore = playerScore
            }
        }
        playerFin = false
        if !player.IsDealer && playerScore > 21 {
            playerBusted = true
        }
    }
}

func Play() {
    d := deck.New()
    d.Shuffle()
    playing := true
    for playing {
        if len(d) < 15 {
            fmt.Println("Reshuffling Deck")
            d = deck.New()
            d.Shuffle()
        }
        players := setupPlayers(2)
        game.Deal(&d, 2, players)
        playRound(players, &d)
        scoreRound(players)
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
