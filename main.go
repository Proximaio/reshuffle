package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Card is JSON record from server
type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
}

// Deck of cards and metadata
type Deck struct {
	Success   bool   `json:"success"`
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards"`
}

func (card Card) String() string {
	suit := strings.Title(strings.ToLower(card.Suit))
	value := strings.Title(strings.ToLower(card.Value))
	return fmt.Sprintf("[%s of %s]", value, suit)
}

// CreateDeck will grab a new deck of cards from API
func CreateDeck(numberOfCards uint) (Deck, error) {
	var deck Deck

	url := fmt.Sprintf("http://deckofcardsapi.com/api/deck/new/draw/?count=%d", numberOfCards)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return deck, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return deck, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&deck); err != nil {
		log.Println(err)
	}

	return deck, nil
}

func main() {
	numberOfCards := uint(52)
	deck, err := CreateDeck(numberOfCards)

	if err == nil {
		fmt.Println("Success", deck.Success)
		fmt.Println("DeckID", deck.DeckID)
		fmt.Println("Shuffled", deck.Shuffled)
		fmt.Println("Remaining", deck.Remaining)
		fmt.Println("Cards", deck.Cards)
	}
}
