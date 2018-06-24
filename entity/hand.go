package entity

import (
	"fmt"
	"strings"
)

// Hand is the cards a player holds
type Hand []*Card

// Empty returns no cards are left in hand
func (h Hand) Empty() bool {
	return h.Size() == 0
}

// Size returns the amount of cards in the hand
func (h Hand) Size() int {
	return len(h)
}

// addCards add cards to the hand
func (h *Hand) addCards(cards []*Card) {
	*h = append(*h, cards...)
}

// dropCard finds card in hand and removes it from it
func (h *Hand) dropCard(card Card) *Card {
	for _, c := range *h {
		if *c == card {
			h.removeCard(c)
			return c
		}
	}
	return nil
}

// RemoveCard removes cards to the hand
func (h *Hand) removeCard(card *Card) {
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
