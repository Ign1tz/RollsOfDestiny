package Types

import (
	"errors"
	"strconv"
)

type Player struct {
	Username              string
	UserID                string
	Mana                  int
	Deck                  Deck
	Die                   Die
	Grid                  Grid
	WebsocketConnectionID string
}

type PlayerLogic interface {
	AddMana()
	RemoveMana()
	ToJson()
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

func (p Player) ToJson(extraInfo bool) string {

	var extra string

	if extraInfo {
		extra = `, "userId": "` + p.UserID + `"`
	} else {
		extra = ""
	}

	message := `{ "WebsocketId": "` + p.WebsocketConnectionID + `", "Username": "` + p.Username + `", "LeftColumn": { "First": "` + strconv.Itoa(p.Grid.Left.First) + `", "Second": "` + strconv.Itoa(p.Grid.Left.Second) + `", "Third": "` + strconv.Itoa(p.Grid.Left.Third) + `", "IsFull": ` + strconv.FormatBool(p.Grid.Left.IsFull) + `}, "MiddleColumn": { "First": "` + strconv.Itoa(p.Grid.Middle.First) + `", "Second": "` + strconv.Itoa(p.Grid.Middle.Second) + `", "Third": "` + strconv.Itoa(p.Grid.Middle.Third) + `", "IsFull": ` + strconv.FormatBool(p.Grid.Middle.IsFull) + `}, "RightColumn": { "First": "` + strconv.Itoa(p.Grid.Right.First) + `", "Second": "` + strconv.Itoa(p.Grid.Right.Second) + `", "Third": "` + strconv.Itoa(p.Grid.Right.Third) + `", "IsFull": ` + strconv.FormatBool(p.Grid.Right.IsFull) + `}` + extra + `}`

	return message
}
