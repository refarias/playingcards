package playingcards

import (
	"testing"
)

const countOfCardValuesPerSuit int = 104

func TestGetDeckOfCards(t *testing.T) {
	suitCount := make(map[CardType]int)
	deckOfCards := GetDeckOfCards()
	// testing the number of cards
	if len(deckOfCards) != 54 {
		t.Errorf("")
	}

	//summing the count of values per suit
	for _, card := range deckOfCards {
		value, ok := suitCount[card.Suit]
		if ok {
			suitCount[card.Suit] = value + int(card.Value)
		} else {
			suitCount[card.Suit] = int(card.Value)
		}
	}

	// testing the total count of values per suit
	if suitCount[HEART] != countOfCardValuesPerSuit {
		t.Errorf("The expected sum of %q values was 104, but the actual result was %d.", HEART, suitCount[HEART])
	}
	if suitCount[DIAMOND] != countOfCardValuesPerSuit {
		t.Errorf("The expected sum of %q values was 104, but the actual result was %d.", DIAMOND, suitCount[DIAMOND])
	}
	if suitCount[CLUB] != countOfCardValuesPerSuit {
		t.Errorf("The expected sum of %q values was 104, but the actual result was %d.", CLUB, suitCount[CLUB])
	}
	if suitCount[SPADE] != countOfCardValuesPerSuit {
		t.Errorf("The expected sum of %q values was 104, but the actual result was %d.", SPADE, suitCount[SPADE])
	}

	// testing the values of Jokes
	if suitCount[""] != 30 {
		t.Errorf("The expected sum of 'jokes' values was 30, but the actual result was %d.", suitCount[""])
	}
}
