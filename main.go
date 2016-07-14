package main

import (
	"flag"
	"fmt"
)

func main() {
	cardFlag := flag.Int("cards", 0, "number of cards, greater than 0")
	serverFlag := flag.Bool("server", false, "run server")
	flag.Parse()

	if *cardFlag > 0 {
		deck, err := GetDeck(DeckOpts{
			Cards: 52,
		})
		if err == nil {
			fmt.Println(deck)
		} else {
			fmt.Println(err)
		}
	}

	if *serverFlag {
		SetUp(Server{}).Start()
	}
}
