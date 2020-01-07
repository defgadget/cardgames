package main

import (
    "github.com/defgadget/cardgames/bin/blackjack"
)

func main() {
    game := blackjack.New(blackjack.Options{NumPlayers: 1, NumDecks: 1})
    game.Play()
}
