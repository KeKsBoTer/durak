package entity

import (
	"math/rand"
	"strings"
)

// CardStack is the representation of a stack of cards where you can only take away the top card
type CardStack struct {
	root *stackNode
	size int
}

// Empty returns wether stack is empty
func (s CardStack) Empty() bool {
	return s.size == 0
}

// Pop removes top card and returns it
func (s *CardStack) Pop() *Card {
	if s.Empty() {
		return nil
	}
	card := s.root.card
	s.root = s.root.next
	s.size--
	return &card
}

// PopN returns the top n cards. If there are less cards than n a smaller list is returned
func (s *CardStack) PopN(n int) []*Card {
	var result []*Card
	for i := 0; i < n && !s.Empty(); i++ {
		result = append(result, s.Pop())
	}
	return result
}

// Peek returns card at top of stack
func (s CardStack) Peek() *Card {
	if s.Empty() {
		return nil
	}
	return &s.root.card
}

// Size returns the amount of cards in the stack
func (s CardStack) Size() (size int) {
	return s.size
}

// Push puts card to top of stack
func (s *CardStack) Push(c Card) {
	top := &stackNode{card: c}
	s.size++
	if s.Empty() {
		s.root = top
		return
	}
	top.next = s.root
	s.root = top
}

func (s CardStack) String() string {
	str := make([]string, s.size)
	node := s.root
	for i := range str {
		str[i] = node.card.String()
		node = node.next
	}
	return "[" + strings.Join(str, ",") + "]"
}

// A card/ node in the card stack
type stackNode struct {
	card Card
	next *stackNode
}

// RandomStack generates a random card stack
func RandomStack(cards []Card) *CardStack {
	// shuffling cards
	for i := len(cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
	stack := new(CardStack)
	for _, c := range cards {
		stack.Push(c)
	}
	return stack
}
