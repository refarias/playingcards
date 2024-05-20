package playingcards

type Card struct {
	Suit  CardType
	Value CardValue
}

func GetDeckOfCards() []Card {
	var deck = make([]Card, 0)
	for _, suit := range Suits {
		for _, value := range Values {
			deck = append(deck, Card{Suit: suit, Value: value})
		}
	}
	deck = append(deck, Card{Suit: "", Value: JOKER})
	deck = append(deck, Card{Suit: "", Value: JOKER})

	return deck
}
