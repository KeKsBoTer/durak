package durak

import (
	"fmt"
	"sort"
	"strings"
)

// Hand is the cards a player holds
type Hand []*Card

// SortBySuit sorts the cards by card Suit
func (h *Hand) SortBySuit() {
	sort.SliceStable(h, func(i, j int) bool {
		return (*h)[i].suit < (*h)[j].suit
	})
}

// SortByNumber sorts the cards by card number
func (h *Hand) SortByNumber() {
	sort.SliceStable(h, func(i, j int) bool {
		return (*h)[i].number < (*h)[j].number
	})
}

// Empty returns no cards are left in hand
func (h Hand) Empty() bool {
	return h.Size() == 0
}

// Size returns the amount of cards in the hand
func (h Hand) Size() int {
	return len(h)
}

// AddCards add cards to the hand
func (h *Hand) AddCards(cards []*Card) {
	*h = append(*h, cards...)
}

// RemoveCard removes cards to the hand
func (h *Hand) RemoveCard(card *Card) {
	for i, c := range *h {
		if c == card {
			*h = append((*h)[:i], (*h)[i+1:]...)
			return
		}
	}
	// Card not found something must be wrong
	panic(fmt.Errorf("Card %v is not in hand %v", *card, *h))
}

func (h Hand) String() string {
	str := make([]string, len(h))
	for i, c := range h {
		str[i] = c.String()
	}
	return "[" + strings.Join(str, ",") + "]"
}
