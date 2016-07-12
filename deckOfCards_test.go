package main

import "testing"

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
		card := CardRecord{c.value, c.suit}
		got := card.String()
		if got != c.expected {
			t.Errorf(
				"For suit %s and value %s got: %s expected: %s",
				c.suit, c.value, got, c.expected,
			)
		}
	}
}
