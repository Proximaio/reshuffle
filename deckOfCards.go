package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CardRecord represents the JSON data sent from server representing a card
type CardRecord struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
}

func (card CardRecord) String() string {
	suit := strings.Title(strings.ToLower(card.Suit))
	value := strings.Title(strings.ToLower(card.Value))
	return fmt.Sprintf("%s of %s", value, suit)
}

// DeckRecord represents the JSON data sent from server representing a deck of cards
type DeckRecord struct {
	Success   bool         `json:"success"`
	DeckID    string       `json:"deck_id"`
	Shuffled  bool         `json:"shuffled"`
	Remaining int          `json:"remaining"`
	Cards     []CardRecord `json:"cards"`
}

func (deck DeckRecord) String() string {
	out := fmt.Sprintf("Remaining: %v\n", deck.Remaining)
	out += fmt.Sprintf("Cards: %v\n", len(deck.Cards))
	for i, c := range deck.Cards {
		out += fmt.Sprintf("\t%d: %v\n", i, c)
	}
	return out
}

// DeckOpts is the configuration for the deck of cards we want
type DeckOpts struct {
	// Should we shuffle?
	Shuffle bool
	// How many decks are we drawing from
	Cards uint8
}

// Doer is the Wrapper for object that is used to actually hit the deck of cards API
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

// Logger is the wrapper for logging
type Logger interface {
	Fatal(v ...interface{})
}

// GetDeck fetches a deck of cards from API
func GetDeck(opts DeckOpts, d Doer, l Logger) (DeckRecord, error) {
	deck := DeckRecord{}
	deck.Success = false

	url := fmt.Sprintf("http://deckofcardsapi.com/api/deck/new/draw/?count=%d", opts.Cards)
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := d.Do(req)

	if err != nil {
		l.Fatal("Error fecthing deck", err)
		return deck, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&deck); err != nil {
		l.Fatal("Error decoding deck", err)
		return deck, err
	}

	return deck, nil
}
