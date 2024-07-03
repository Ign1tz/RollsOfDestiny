package Types

import (
	"errors"
	"fmt"
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
		extra = `, "userId": "` + p.UserID + `", "deck": ` + handleDeckToString(p.Deck)
	} else {
		extra = `, "enemyDeck": ` + handleDeckToStringEnemy(p.Deck)
	}
	message := `{ "WebsocketId": "` + p.WebsocketConnectionID + `", "Username": "` + p.Username + `", "Score": ` + strconv.Itoa(p.Grid.Value()) + `, "LeftColumn": { "First": "` + strconv.Itoa(p.Grid.Left.First) + `", "Second": "` + strconv.Itoa(p.Grid.Left.Second) + `", "Third": "` + strconv.Itoa(p.Grid.Left.Third) + `", "IsFull": ` + strconv.FormatBool(p.Grid.Left.IsFull()) + `}, "MiddleColumn": { "First": "` + strconv.Itoa(p.Grid.Middle.First) + `", "Second": "` + strconv.Itoa(p.Grid.Middle.Second) + `", "Third": "` + strconv.Itoa(p.Grid.Middle.Third) + `", "IsFull": ` + strconv.FormatBool(p.Grid.Middle.IsFull()) + `}, "RightColumn": { "First": "` + strconv.Itoa(p.Grid.Right.First) + `", "Second": "` + strconv.Itoa(p.Grid.Right.Second) + `", "Third": "` + strconv.Itoa(p.Grid.Right.Third) + `", "IsFull": ` + strconv.FormatBool(p.Grid.Right.IsFull()) + `}` + extra + `}`
	return message
}

func handleDeckToString(deck Deck) string {
	cardsLeft := 0
	cardsLeft = 0
	cardsInHand := ""
	for cardIndex := range deck.Cards {
		if deck.Cards[cardIndex].CardID != "" {
			if !deck.Cards[cardIndex].Played && !deck.Cards[cardIndex].InHand {
				cardsLeft += 1
			}
			if deck.Cards[cardIndex].InHand {
				cardsInHand = fmt.Sprintf(`%s, {"name": "%s", "cost": %s, "picture": "%s", "effect": "%s"}`, cardsInHand, deck.Cards[cardIndex].Name, strconv.Itoa(deck.Cards[cardIndex].Cost), deck.Cards[cardIndex].Picture, deck.Cards[cardIndex].Effect)
			}
		}
	}
	if cardsInHand != "" {
		cardsInHand = cardsInHand[2:]
	}
	infoMessage := `{"cardsLeft": ` + strconv.Itoa(cardsLeft) + `, "inHand": [` + cardsInHand + `]}`
	return infoMessage
}

func handleDeckToStringEnemy(deck Deck) string {
	cardsLeft := 0
	cardsInHand := 0
	for cardIndex := range deck.Cards {
		if !deck.Cards[cardIndex].Played && !deck.Cards[cardIndex].InHand {
			cardsLeft += 1
		}
		if deck.Cards[cardIndex].InHand {
			cardsInHand++
		}
	}

	infoMessage := `{"cardsLeft": ` + strconv.Itoa(cardsLeft) + `, "inHand": ` + strconv.Itoa(cardsInHand) + `}`
	return infoMessage
}
