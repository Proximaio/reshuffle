package main

import "fmt"

type DeckEndpoints struct{}

func (d DeckEndpoints) NewDeck(cards uint) string {
	return fmt.Sprintf("http://deckofcardsapi.com/api/deck/new/draw/?count=%d")
}

func main() {
	deck, err := GetDeck(DeckOpts{
		Shuffle:   true,
		Cards:     52,
		Endpoints: DeckEndpoints{},
	})

	if err == nil {
		fmt.Println(deck)
	}

	StartServer(ServerOpts{})
}
