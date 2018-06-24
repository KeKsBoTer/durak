package entity

import (
	"errors"
	"fmt"
	"math/rand"
)

// Game is the state of a running game
type Game struct {
	players    []Player
	stack      CardStack
	middlePool AttackPool
	trump      *Card
	// the player who is attacing
	activePlayer int // index of the active player
}

// NewGame generates a new game with a random stack and gives 6 cards to every user
func NewGame(users []User, cards []Card) (*Game, error) {
	game := new(Game)
	game.stack = *RandomStack(cards)
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

	game.activePlayer = rand.Intn(lenUsers)

	return game, nil
}

// GetTrump returns trump card
func (g Game) GetTrump() Card {
	return *g.trump
}

// ActivePlayer returns the player which is attacking
func (g Game) ActivePlayer() *Player {
	return &g.players[g.activePlayer]
}

// ActivePlayers returns the count of active playerds
func (g Game) ActivePlayers() (active int) {
	for _, p := range g.players {
		if p.IsPlaying() {
			active++
		}
	}
	return
}

// NextPlayer sets next player as active player
func (g *Game) NextPlayer() {
	g.activePlayer = g.nextPlayer(g.activePlayer)
}

// returns next playing player after given player index
// skippes players who are not playing any more
func (g Game) nextPlayer(start int) int {
	var next = start + 1
	players := len(g.players)
	if next >= players {
		next = 0
	}
	if g.players[next].IsPlaying() {
		return next
	}
	for ; !g.players[next].IsPlaying(); next++ {
	} // count up until reaching a active player
	return next
}

// Player returns player by user
func (g Game) Player(u User) *Player {
	for _, p := range g.players {
		if p.user == u {
			return &p
		}
	}
	panic(fmt.Errorf("user %v is not participating in the game", u))
}

// SetFinish marks the player/user as finished
func (g *Game) SetFinish(u User) {
	for i, p := range g.players {
		if p.user == u {
			g.players[i].finished = true
			return
		}
	}
	panic(fmt.Errorf("user %v is not participating in the game", u))
}

func (g Game) defender() int {
	return g.nextPlayer(g.activePlayer)
}

// IsAttacker checks if a user is currently attacking
/* TODO move to game rules
func (g Game) IsAttacker(u User) bool {
	activePlayer := g.players[g.activePlayer]
	if activePlayer.user == u {
		return true
	} else if active := g.ActivePlayers(); active == 2 {
		return false
	} else if active < 2 {
		panic(errors.New("cannot play with less than 2 players"))
	}
	defender := g.defender()
	afterdefender := g.nextPlayer(defender)
	return g.players[afterdefender].user == u
}
*/

// PlayerAfterDefender returns the player after the defender
func (g Game) PlayerAfterDefender() *Player {
	afterDefender := g.nextPlayer(g.defender())
	return &g.players[afterDefender]
}

// IsDefender checks if user is currently defending
func (g Game) IsDefender(u User) bool {
	defender := g.players[g.defender()]
	return defender.user == u
}

// Defender returns the defending user
func (g Game) Defender() *Player {
	return &g.players[g.defender()]
}

// IsNumberPresentInMiddle checks if the cards number is present in the middle
func (g *Game) IsNumberPresentInMiddle(card Card) bool {
	return g.middlePool.ContainsNumber(card.number)
}

// Attack adds card to middle
// checks if attacker holds card and removes it from his hand
func (g *Game) Attack(user User, c Card) error {
	card, err := g.dropCard(user, c)
	if err != nil {
		return err
	}
	g.middlePool.attack(card)
	return nil
}

func (g *Game) dropCard(user User, card Card) (*Card, error) {
	player := g.Player(user)
	c := player.hand.dropCard(card)
	if c == nil {
		return nil, errors.New("you do not have this card")
	}
	return c, nil
}

// Defend defends card in middle
// att is the card to defend
// def is the card to defend the attack with
func (g *Game) Defend(user User, att, def Card) error {
	card, err := g.dropCard(user, def)
	if err != nil {
		return err
	}
	return g.middlePool.defend(card, att)
}

// IsEverythingDefended checks if every cards in the middle is defended
func (g Game) IsEverythingDefended() bool {
	return g.middlePool.IsEverythingDefended()
}

// GiveCards adds the given cards to the players hand
func (g *Game) GiveCards(u User, cards []*Card) {
	g.Player(u).hand.addCards(cards)
}

// AttackPool returns all cards in the middle
func (g Game) AttackPool() []Card {
	cards := []Card{}
	for _, a := range g.middlePool {
		cards = append(cards, *a.attacker, *a.defender)
	}
	return cards
}

// ClearMiddle removes card form middle and returns them
func (g *Game) ClearMiddle() []*Card {
	return g.middlePool.clear()
}

// CanDefend checks if card is present in middle and not defended yet
func (g Game) CanDefend(att Card) bool {
	return g.middlePool.ContainsAttackCard(att) && g.middlePool.IsUndefended(att)
}
