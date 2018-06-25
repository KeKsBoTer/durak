package entity

// Action can be performed by a user
type Action int

const (
	// Attack defener with card
	Attack Action = iota

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
	Hand     Hand
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

// AddActions adds action to possible actions
func (p *Player) AddActions(actions ...Action) {
	for _, a := range actions {
		if !p.CanDo(a) {
			p.actions = append(p.actions, a)
		}
	}
}

// RemoveActions removes action from possible actions
func (p *Player) RemoveActions(actions ...Action) {
	for i := 0; i < len(p.actions); i++ {
		if Actions(actions).Contains(p.actions[i]) {
			p.actions = append(p.actions[:i], p.actions[i+1:]...)
			i--
		}
	}
	// Card not found something must be wrong
	// panic(fmt.Errorf("Action %v is not possible %v", action, p.actions))
}

// Finished marks player as finished
func (p *Player) Finished() {
	p.finished = true
}
