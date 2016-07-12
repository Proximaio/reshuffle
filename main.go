package main

import (
	"fmt"
	"net/http"
)

type DeckEndpoints struct{}

func (d DeckEndpoints) NewDeck(cards uint) string {
	return fmt.Sprintf("http://deckofcardsapi.com/api/deck/new/draw/?count=%d")
}

func main() {

	client := &http.Client{}

	deck, err := GetDeck(DeckOpts{
		Shuffle:   true,
		Cards:     52,
		Endpoints: DeckEndpoints{},
	}, client)

	if err == nil {
		fmt.Println(deck)
	}
}
