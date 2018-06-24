package durak

import (
	"errors"

	"github.com/KeKsBoTer/durak/entity"
)

// Manager manages game
type Manager struct {
	game  entity.Game
	users []entity.User
}

// NewGame creates new game from users with default rules
func NewGame(users []entity.User) (manager *Manager) {
	game, err := entity.NewGame(users, entity.AllCards)
	if err != nil {
		panic(err)
	}
	manager.game = *game

	// starting player can attack
	game.ActivePlayer().AddAction(entity.Attack)
	return
}

// DoAction performs action with card
func (m *Manager) DoAction(user entity.User, action entity.Action, cards ...entity.Card) error {
	player := m.game.Player(user)
	if !player.CanDo(action) {
		return errors.New("you are not allowed to to this")
	}
	switch action {
	case entity.Attack:
		attackCard := cards[0]

		// no cheating for now
		if !m.game.IsNumberPresentInMiddle(attackCard) {
			return errors.New("no cheating")
		}
		fallthrough

	case entity.FirstAttack:
		attackCard := cards[0]
		// move card from player hand to middle
		err := m.game.Attack(user, attackCard)
		if err != nil {
			return err
		}
		// Allow defender to defend
		defender := m.game.Defender()
		if !defender.CanDo(entity.Defend) {
			defender.AddAction(entity.Defend)
		}
		return nil
	case entity.Defend:
		attackCard, defendCard := cards[0], cards[1]
		if !m.game.CanDefend(attackCard) {
			return errors.New("card is not present in middle or allready defended")
		}
		// no cheating for now
		if !defendCard.Trumps(attackCard, m.game.GetTrump()) {
			return errors.New("no cheating")
		}
		if err := m.game.Defend(user, attackCard, defendCard); err != nil {
			return err
		}

		if m.game.IsEverythingDefended() {
			// Every card is defended, no more defending possible
			m.game.Defender().RemoveAction(entity.Defend)

			// attackers can pass if they want to to no further attacks
			for _, p := range m.getAttackers() {
				p.AddAction(entity.Pass)
			}
		}

		return nil
	case entity.Pass:
		return nil
	case entity.TakeCards:
		// give cards from the middle to user
		middle := m.game.ClearMiddle()
		m.game.GiveCards(user, middle)
		return nil
	}
	return errors.New("undefined action")
}

// for now the attackers are the people next to the defender
func (m Manager) getAttackers() []*entity.Player {
	activePlayer := m.game.ActivePlayer()
	players := []*entity.Player{activePlayer}
	after := m.game.PlayerAfterDefender()
	if activePlayer != after {
		players = append(players, after)
	}
	return players
}
