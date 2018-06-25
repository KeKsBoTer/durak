package durak

import (
	"errors"
	"fmt"

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
	game.ActivePlayer().AddActions(entity.Attack)
	// defender needs to defend itself
	game.Defender().AddActions(entity.Defend)
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
		if !m.game.Middle.IsEmpty() && !m.game.Middle.ContainsCardNumber(attackCard) {
			return errors.New("no cheating")
		}
		// move card from player hand to middle
		err := m.game.Attack(user, attackCard)
		if err != nil {
			return err
		}
		// Allow defender to defend
		defender := m.game.Defender()
		defender.AddActions(entity.TakeCards)

	case entity.Defend:
		attackCard, defendCard := cards[0], cards[1]
		// check if there is a not defended attack with that card
		if !m.game.Middle.CanDefend(attackCard) {
			return errors.New("card is not present in middle or allready defended")
		}
		// no cheating for now
		if !defendCard.Trumps(attackCard, m.game.GetTrump()) {
			return errors.New("no cheating")
		}
		if err := m.game.Defend(user, attackCard, defendCard); err != nil {
			return err
		}
		if m.game.Middle.IsEverythingDefended() {
			player.RemoveActions(entity.TakeCards)
		}
		activePlayer := m.game.ActivePlayer()
		if !activePlayer.CanDo(entity.Pass) {
			// if the active player cannot pass, this means he allready passed and the other players can attack too
			for _, p := range m.getAttackers() {
				// every player that allready passed can attack now
				p.AddActions(entity.Attack, entity.Pass)
			}
		} else {
			activePlayer.AddActions(entity.Attack, entity.Pass)
		}

	case entity.Pass:
		// player can only pass once
		player.RemoveActions(entity.Pass, entity.Attack)

		if player == m.game.ActivePlayer() {
			// allow other attackers to attack
			for _, p := range m.getAttackers() {
				if p != player {
					p.AddActions(entity.Attack, entity.Pass)
				}
			}
		}
	case entity.TakeCards:
		// give cards from the middle to user
		middle := m.game.Middle.Clear()
		player.Hand.AddCards(middle)
		player.RemoveActions(entity.TakeCards)
	default:
		return fmt.Errorf("undefined action: %v", action)
	}

	// check if attack is defend off
	m.checkTriggers()

	return nil
}

// for now the attackers are the people next to the defender
func (m Manager) getAttackers() []*entity.Player {
	activePlayer := m.game.ActivePlayer()
	after := m.game.PlayerAfterDefender()
	if activePlayer != after {
		return []*entity.Player{activePlayer, after}
	}
	return []*entity.Player{activePlayer}
}

func (m *Manager) checkTriggers() {
	pileEmpty := m.game.Pile.Empty()
	if pileEmpty {
		// check for winners
		for _, p := range m.game.ActivePlayers() {
			if p.Hand.Empty() {
				p.Finished()
			}
		}
	}
	if m.game.Middle.IsEmpty() {
		// fill up cards in player hands
		if !pileEmpty {
			for _, p := range m.getAttackers() {
				missing := 6 - p.Hand.Size()
				p.Hand.AddCards(m.game.Pile.PopN(missing))
				if m.game.Pile.Empty() {
					break
				}
			}
			if !m.game.Pile.Empty() {
				defender := m.game.Defender()
				missing := 6 - defender.Hand.Size()
				defender.Hand.AddCards(m.game.Pile.PopN(missing))
			}
		}
		// player after defender is attacker
		m.game.NextPlayer()
		m.game.NextPlayer()

		m.newRound()
		return
	}

	if m.game.Middle.IsEverythingDefended() && m.allAttackersPassed() {
		// everything is defended and all attackers passed
		m.game.Middle.Clear()

		// defender is new attacker
		m.game.NextPlayer()
		m.newRound()
		return
	}
}

func (m *Manager) newRound() {
	// round is over, nobody can do anything any more
	m.game.ResetActions()

	// allow to attack and defend
	m.game.ActivePlayer().AddActions(entity.Attack)
	m.game.Defender().AddActions(entity.Defend)
}

func (m *Manager) allAttackersPassed() bool {
	for _, a := range m.getAttackers() {
		if a.CanDo(entity.Pass) {
			return false
		}
	}
	return true
}
