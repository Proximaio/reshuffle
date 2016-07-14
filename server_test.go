package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type testEndpoints struct {
	url string
}

func (t testEndpoints) NewDeck(cards uint) string {
	return t.url
}

func deckServer(body string, response int) (server *httptest.Server, client *http.Client, endpoints Endpoints) {
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

func TestCarddString(t *testing.T) {
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

	srv := SetUp(Server{
		Endpoints: endpoints,
		Client:    client,
	})

	req, _ := http.NewRequest("GET", "/deck/10", nil)
	resp := httptest.NewRecorder()
	srv.Router.ServeHTTP(resp, req)
}

func TestCardBaddd(t *testing.T) {
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

	srv := SetUp(Server{
		Endpoints: endpoints,
		Client:    client,
	})

	req, _ := http.NewRequest("GET", "/deck/10A", nil)
	resp := httptest.NewRecorder()
	srv.Router.ServeHTTP(resp, req)
}

func TestCardHorro(t *testing.T) {
	response := int(200)
	body :=
		`{
				"success": true,
				"cards": [
						{
								"image": "http://deckofcardsapi.com/static/img/KH.png",
								"value": "KING",
								"suit": "HEARTS",
								"co`
	server, client, endpoints := DeckServer(body, response)
	defer server.Close()

	srv := SetUp(Server{
		Endpoints: endpoints,
		Client:    client,
	})

	req, _ := http.NewRequest("GET", "/deck/10", nil)
	resp := httptest.NewRecorder()
	srv.Router.ServeHTTP(resp, req)
}
