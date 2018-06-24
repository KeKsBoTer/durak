package entity

import "fmt"

// AttackStack holds the card who was attacked with and the one which fought off the attack
type AttackStack struct {
	attacker, defender *Card
}

func (a AttackStack) String() string {
	return fmt.Sprintf("{%v->%v}", *a.attacker, *a.defender)
}

// AttackPool pool represents the cards in the middle
type AttackPool []AttackStack

// ContainsNumber checks if any card in the pool has the given number
func (p AttackPool) ContainsNumber(n CardNumber) bool {
	for _, a := range p {
		if a.attacker.number == n || (a.defender != nil && a.defender.number == n) {
			return true
		}
	}
	return false
}

// ContainsAttackCard checks any attack has this card
func (p AttackPool) ContainsAttackCard(c Card) bool {
	for _, a := range p {
		if *a.attacker == c {
			return true
		}
	}
	return false
}

// IsUndefended checks if attack is undefended
func (p AttackPool) IsUndefended(c Card) bool {
	for _, a := range p {
		if *a.attacker == c {
			return a.defender == nil
		}
	}
	panic(fmt.Errorf("card %v is not present in middle", c))
}

// Attack creates new attack
func (p *AttackPool) attack(c *Card) error {
	/* TODO do this in game logic to allow easier rule changes
	if !p.ContainsNumber(c.number) {
		return errors.New("tried to attack with card not present in the attack pool")
	}
	*/
	*p = append(*p, AttackStack{attacker: c})
	return nil
}

// Defend defends of card in the middle
func (p *AttackPool) defend(def *Card, att Card) error {
	for i, a := range *p {
		if *a.attacker == att {
			(*p)[i].defender = def
			return nil
		}
	}
	return fmt.Errorf("Card %v was not found in the attack pool", att)
}

// Clear removes all cards from pool and returns them as list
func (p *AttackPool) clear() []*Card {
	cards := []*Card{}
	for _, a := range *p {
		cards = append(cards, a.attacker, a.defender)
	}
	*p = AttackPool{}
	return cards
}

// IsEverythingDefended checks if every attack in the middle was defended
func (p AttackPool) IsEverythingDefended() bool {
	for _, a := range p {
		if a.defender == nil {
			return false
		}
	}
	return true
}
