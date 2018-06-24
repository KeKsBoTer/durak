package entity

import (
	"fmt"
)

// AllCards are all possible Cards
var AllCards []Card

func init() {
	// Calc all available cards
	AllCards = make([]Card, len(AllSuits)*len(AllNumbers))
	i := 0
	for _, c := range AllSuits {
		for _, n := range AllNumbers {
			AllCards[i] = Card{suit: c, number: n}
			i++
		}
	}
}

// CardSuit is one of: diamonds (♦), clubs (♣), hearts (♥) and spades (♠)
type CardSuit int

const (
	// Heart ♥
	Heart CardSuit = iota
	// Diamond ♦
	Diamond
	// Club ♣
	Club
	// Spade ♠
	Spade
)

// AllSuits are all available card Suits
var AllSuits = []CardSuit{Heart, Diamond, Club, Spade}

// SuitSymbols are the number of tha suit names
var SuitSymbols = []string{"♥", "♦", "♣", "♠"}

func (c CardSuit) String() string {
	return SuitSymbols[c]
}

// CardNumber is the number value of the card (6-10,Boy,Queen,King,Ass)
type CardNumber int

// AllNumbers are all available card numbers
var AllNumbers = []CardNumber{0, 1, 2, 3, 4, 5, 6, 7, 8}

var numberNames = []string{"6", "7", "8", "9", "10", "B", "Q", "K", "A"}

// Valid returns if card number is valid
func (c CardNumber) Valid() bool {
	return c >= 0 && c <= 8
}

func (c CardNumber) String() string {
	return numberNames[c]
}

// Card is a game card
type Card struct {
	suit   CardSuit
	number CardNumber
}

// Trumps returns if the card trumps another given card
func (c Card) Trumps(c2 Card, trump Card) bool {
	if c.suit == c2.suit {
		return c.number-c2.number > 0
	}
	return c.suit == trump.suit && c2.suit != trump.suit
}

func (c Card) String() string {
	return fmt.Sprintf("{%v%v}", c.suit, c.number)
}
