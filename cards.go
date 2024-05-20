package playingcards

type Card struct {
	suit  Suit
	value CardValue
}

func GetDeckOfCards() []Card {
	var deck = make([]Card, 0)
	for _, suit := range Suits {
		for _, value := range Values {
			deck = append(deck, Card{suit: suit, value: value})
		}
	}
	deck = append(deck, Card{suit: "", value: JOKER})
	deck = append(deck, Card{suit: "", value: JOKER})

	return deck
}
