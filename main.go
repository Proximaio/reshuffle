package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	client := &http.Client{}
	logger := &log.Logger{}

	deck, err := GetDeck(DeckOpts{
		Shuffle: true,
		Cards:   52,
	}, client, logger)

	if err == nil {
		fmt.Println(deck)
	}
}
