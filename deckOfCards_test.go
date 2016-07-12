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
