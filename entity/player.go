package entity

import "fmt"

// Action can be performed by a user
type Action int

const (
	// Attack defener with card
	Attack Action = iota

	// FirstAttack starting a attack
	FirstAttack

	// TakeCards from the middle
	TakeCards

	// Defend attack with card
	Defend

	// Pass to attack defender
	Pass
)

// Actions is a list of actions a user can perform
type Actions []Action

// Contains checks if list contains a action
func (a Actions) Contains(ac Action) bool {
	for _, v := range a {
		if v == ac {
			return true
		}
	}
	return false
}

// User is the identification string for the user
type User string

// Player is a user with cards
type Player struct {
	user     User
	hand     Hand
	finished bool

	// Actions are the fings the user can do
	actions Actions
}

// IsPlaying checks if player is still playing
func (p Player) IsPlaying() bool {
	return !p.finished
}

// CanDo checks if player can perform action
func (p Player) CanDo(a Action) bool {
	return p.actions.Contains(a)
}

// AddAction adds action to possible actions
func (p *Player) AddAction(a Action) {
	if p.CanDo(a) {
		panic("user allready can do this action")
	}
	p.actions = append(p.actions, a)
}

// RemoveAction removes action from possible actions
func (p *Player) RemoveAction(action Action) {
	for i, a := range p.actions {
		if a == action {
			p.actions = append(p.actions[:i], p.actions[i+1:]...)
			return
		}
	}
	// Card not found something must be wrong
	panic(fmt.Errorf("Action %v is not possible %v", action, p.actions))
}
