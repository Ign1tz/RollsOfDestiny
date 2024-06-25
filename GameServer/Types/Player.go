package Types

import (
	"errors"
)

type Player struct {
	Username string
	UserID   string
	Mana     int
	Deck     Deck
	Die      Die
	Grid     Grid
}

type PlayerLogic interface {
	AddMana()
	RemoveMana()
}

func (p *Player) AddMana(additionalMana int) {
	p.Mana = min((p.Mana + additionalMana), 10)
}

func (p *Player) RemoveMana(spent int) error {
	if p.Mana < spent {
		return errors.New("You don't have enough mana")
	} else {
		p.Mana -= spent
	}
	return nil
}
