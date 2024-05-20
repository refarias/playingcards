package playingcards

type Suit string

const (
	HEART   Suit = "heart"
	DIAMOND Suit = "diamond"
	CLUB    Suit = "club"
	SPADE   Suit = "spade"
)

var Suits = []Suit{
	HEART,
	DIAMOND,
	CLUB,
	SPADE,
}
