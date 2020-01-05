package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"

    "github.com/defgadget/cardgames/pkg/deck"
    "github.com/defgadget/cardgames/pkg/game"
)
// Setup Players
// Deal 2 cards to each player, including dealer
// Players cards are visible, only 2nd card of delear is visible.
// Allow players other then Dealer to hit/stay
// 21 or Bust automatically ends player's turn
// Compare users score to Dealer score
// If dealer has highest score, round over Players lose
// If dealer has lower score, dealer hits until over 17

func getChoice(msg string) string {
    trim := strings.TrimSpace
    lower := strings.ToLower
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(msg)
    text, _ := reader.ReadString('\n')
    text = lower(trim(text))
    return text
}
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

func main() {
    players := setupPlayers(2)
    d := deck.New()
    d.Shuffle()
    game.Deal(&d, 2, players)
    dealer := players[len(players)-1]
    playerFin := false
    playerScore := 0
    dealerScore := scoreHand(dealer.Hand)
    choice := ""
    for i, player := range players {
        fmt.Printf("--- %v ---\n----------------\n", player.Name)
        playerScore = scoreHand(players[i].Hand)
        for !playerFin {
            fmt.Printf("%v\n\n", players[i].Hand)
            if checkNoChoice(playerScore) {
                break
            }
            if !player.IsDealer {
                choice = getChoice("Would you like to Hit or Stay? ")
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
        if !player.IsDealer && playerScore > 21 || playerScore < dealerScore {
            fmt.Println("You lost")
            fmt.Printf("Dealer Score: %v\n%v Score: %v\n", dealerScore, player.Name, playerScore)
            break
        }
    }

    // See who won
    for _, player := range players {
        if player.IsDealer {
            break
        }
        fmt.Println(player.Hand)
        playerScore = scoreHand(player.Hand)
        switch {
        case playerScore == dealerScore:
            fmt.Println("Push")
            fmt.Println("Dealer Score:", dealerScore)
            fmt.Println("Your Score:", playerScore)
        case playerScore < dealerScore:
            fmt.Println("You Lost")
            fmt.Println("Dealer Score:", dealerScore)
            fmt.Println("Your Score:", playerScore)
        case playerScore > dealerScore:
            fmt.Println("You Won!")
            fmt.Println("Dealer Score:", dealerScore)
            fmt.Println("Your Score:", playerScore)
        }
    }
}
