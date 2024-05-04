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
}

type PlayerLogic interface {
	AddMana()
	RemoveMana()
	ChangeRating()
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

/*func (p *Player) ChangeRating(ratingChange int) {
	p.Rating = max((p.Rating + ratingChange), 0)
}*/
