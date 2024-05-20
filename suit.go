package playingcards

type CardType string

const (
	HEART   CardType = "heart"
	DIAMOND CardType = "diamond"
	CLUB    CardType = "club"
	SPADE   CardType = "spade"
)

var Suits = []CardType{
	HEART,
	DIAMOND,
	CLUB,
	SPADE,
}
