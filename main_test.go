package main

import "testing"

func TestCardString(t *testing.T) {
	card := Card{"10", "SPADES"}
	s := card.String()
	e := "[10 of Spades]"
	if s != e {
		t.Error(
			"For", card,
			"expected", e,
			"got", s,
		)
	}
}
