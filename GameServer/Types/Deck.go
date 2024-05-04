package Types

import (
	"errors"
	"math/rand"
)

type Deck struct {
	Name   string
	UserID string
	Cards  []Card
	Size   int
}

type DeckLogic interface {
	Shuffle()
	Draw()
}

func (d *Deck) Shuffle() {
	for i := range d.Cards {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

func (d *Deck) Draw() (Card, error) {
	if len(d.Cards) <= 0 {
		return Card{Name: "Error"}, errors.New("Not enough cards to draw!")
	} else {
		drawnCard := d.Cards[0]
		d.Cards = d.Cards[1:]
		return drawnCard, nil
	}
}
