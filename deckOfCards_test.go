package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type TestEndpoints struct {
	url string
}

func (t TestEndpoints) NewDeck(cards uint) string {
	return t.url
}

func DeckServer(body string, response int) (server *httptest.Server, client *http.Client, endpoints Endpoints) {
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(response)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, body)
	}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}
	client = &http.Client{Transport: transport}

	endpoints = TestEndpoints{server.URL}

	return
}

func TestCardString(t *testing.T) {
	var cardTests = []struct {
		suit, value, expected string
	}{
		{"SPADES", "10", "10 of Spades"},
		{"DIAMONDS", "JACK", "Jack of Diamonds"},
		{"CLUBS", "1", "1 of Clubs"},
		{"HEARTS", "KING", "King of Hearts"},
	}

	for _, c := range cardTests {
		card := CardRecord{
			Value: c.value,
			Suit:  c.suit,
		}
		got := card.String()
		if got != c.expected {
			t.Errorf(
				"For suit %s and value %s got: %s expected: %s",
				c.suit, c.value, got, c.expected,
			)
		}
	}
}

func TestDeckString(t *testing.T) {
	emptyDeck := DeckRecord{
		Remaining: 52,
		Cards:     nil,
	}
	fullDeck := DeckRecord{
		Remaining: 0,
		Cards: []CardRecord{
			{"SPADES", "10"},
			{"DIAMONDS", "JACK"},
			{"CLUBS", "1"},
			{"HEARTS", "KING"},
		},
	}

	emptyDeckExpected := "Remaining: 52\nCards: 0"
	emptyDeckGot := emptyDeck.String()

	fullDeckExpected := "Remaining: 0\n" +
		"Cards: 4\n" +
		"\t0: 10 of Spades\n" +
		"\t1: Jack of Diamonds\n" +
		"\t2: 1 of Clubs\n" +
		"\t3: King of Hearts"
	fullDeckGot := fullDeck.String()

	if emptyDeckExpected != emptyDeckGot {
		t.Errorf("For emptyDeck\ngot:\n%v\nexpected:\n%v",
			emptyDeckGot, emptyDeckExpected)
	}
	if fullDeckExpected != fullDeckGot {
		t.Errorf("For fullDeck\ngot:\n%v\nexpected:\n%v",
			fullDeckGot, fullDeckExpected)
	}
}

func TestGetDeck(t *testing.T) {
	response := int(200)
	body :=
		`{
				"success": true,
				"cards": [
						{
								"image": "http://deckofcardsapi.com/static/img/KH.png",
								"value": "KING",
								"suit": "HEARTS",
								"code": "KH"
						},
						{
								"image": "http://deckofcardsapi.com/static/img/8C.png",
								"value": "8",
								"suit": "CLUBS",
								"code": "8C"
						}
				],
				"deck_id":"3p40paa87x90",
				"remaining": 50
			}`
	server, client, endpoints := DeckServer(body, response)
	defer server.Close()

	deck, err := GetDeck(DeckOpts{
		Shuffle:   true,
		Cards:     2,
		Endpoints: endpoints,
	}, client)

	if err != nil {
		t.Errorf("Expected err = nil. Got err = %s", err)
	}

	if len(deck.Cards) != 2 {
		t.Errorf("Expected 2 cards. Got %d", len(deck.Cards))
	}
}

func TestGetDeckBadBody(t *testing.T) {
	response := int(200)
	body :=
		`{
				"success": true,
				"cards": [
						{
								"image": "http://deckofcardsapi.com/static/img/KH.png",
								"value": "KING",
								"sui`

	server, client, endpoints := DeckServer(body, response)
	defer server.Close()

	deck, err := GetDeck(DeckOpts{
		Shuffle:   true,
		Cards:     2,
		Endpoints: endpoints,
	}, client)

	if err == nil {
		t.Errorf("Expected err != nil. Got err = %s", err)
	}

	if len(deck.Cards) != 0 {
		t.Errorf("Expected 0 cards. Got %d", len(deck.Cards))
	}
}

func TestGetDeck404(t *testing.T) {
	response := int(404)
	body := ""

	server, client, endpoints := DeckServer(body, response)
	defer server.Close()

	deck, err := GetDeck(DeckOpts{
		Shuffle:   true,
		Cards:     2,
		Endpoints: endpoints,
	}, client)

	if err == nil {
		t.Errorf("Expected err != nil. Got err = %s", err)
	}

	if len(deck.Cards) != 0 {
		t.Errorf("Expected 0 cards. Got %d", len(deck.Cards))
	}
}

func TestGetDeckBadRequest(t *testing.T) {
	response := int(404)
	body := ""

	server, client, endpoints := DeckServer(body, response)
	server.Close()

	deck, err := GetDeck(DeckOpts{
		Shuffle:   true,
		Cards:     2,
		Endpoints: endpoints,
	}, client)

	if err == nil {
		t.Errorf("Expected err != nil. Got err = %s", err)
	}

	if len(deck.Cards) != 0 {
		t.Errorf("Expected 0 cards. Got %d", len(deck.Cards))
	}
}
