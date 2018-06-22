package durak

import (
	"fmt"
	"math/rand"
)

// User is the identification string for the user
type User string

// Player is a user with cards
type Player struct {
	user     User
	hand     Hand
	finished bool
}

// IsPlaying checks if player is still playing
func (p Player) IsPlaying() bool {
	return !p.finished
}

// Game is the state of a running game
type Game struct {
	players      []Player
	stack        CardStack
	middlePool   AttackPool
	trump        *Card
	activePlayer *Player // pointer to one of the players
}

// NewGame generates a new game with a random stack and gives 6 cards to every user
func NewGame(users []User) (*Game, error) {
	game := new(Game)
	game.stack = *RandomStack()
	game.trump = game.stack.Pop()
	lenUsers := len(users)

	// check if player count is ok
	if lenUsers == 0 {
		return nil, fmt.Errorf("cannot ceate game with zero players")
	} else if lenCards := len(AllCards); lenUsers > lenCards/6 {
		return nil, fmt.Errorf("to many players(%d) for %d cards", lenUsers, lenCards)
	}
	game.players = make([]Player, len(users))
	// hand out cards to players
	for i, u := range users {
		game.players[i] = Player{
			user:     u,
			finished: false,
			hand:     game.stack.PopN(6),
		}
	}

	// choose random player who starts
	game.activePlayer = &game.players[rand.Intn(lenUsers)]

	return game, nil
}

// GetTrump returns trump card
func (g Game) GetTrump() *Card {
	return g.trump
}

// Attack holds the card who was attacked with and the one which fought off the attack
type Attack struct {
	attacker, defender *Card
}

func (a Attack) String() string {
	return fmt.Sprintf("{%v->%v}", *a.attacker, *a.defender)
}

// AttackPool pool represents the cards in the middle
type AttackPool []Attack

// ContainsNumber checks if any card in the pool has the given number
func (p AttackPool) ContainsNumber(n CardNumber) bool {
	for _, a := range p {
		if a.attacker.number == n || (a.defender != nil && a.defender.number == n) {
			return true
		}
	}
	return false
}

// Attack creates new attack
func (p *AttackPool) Attack(c *Card) error {
	/* TODO do this in game logic to allow easier rule changes
	if !p.ContainsNumber(c.number) {
		return errors.New("tried to attack with card not present in the attack pool")
	}
	*/
	*p = append(*p, Attack{attacker: c})
	return nil
}

// Defend defends of card in the middle
func (p *AttackPool) Defend(def, att *Card) error {
	for i, a := range *p {
		if a.attacker == att {
			(*p)[i].defender = def
			return nil
		}
	}
	return fmt.Errorf("Card %v was not found in the attack pool", *att)
}

// Clear clears attack pool and removes all cards
func (p *AttackPool) Clear() {
	*p = AttackPool{}
}
