package GameLogic

import (
	"RollsOfDestiny/GameServer/Database"
	"RollsOfDestiny/GameServer/Types"
	"log"
	"math/rand/v2"
	"strconv"
)

func HandleCards(message Types.WebsocketMessage, position Types.Position) (map[string]string, map[string]string) {
	card, err := Database.GetCardById(message.MessageBody)

	if err != nil {
		log.Println(err)
		return nil, nil
	}

	if !card.InHand || card.Played {
		return nil, nil
	}

	playfield, err := Database.GetPlayfield(message.GameId)

	if err != nil {
		log.Println("noPlayfield", err)
		return nil, nil
	}

	var hostIsActive bool
	if playfield.ActivePlayer.UserID == playfield.Host.UserID {
		hostIsActive = true
	} else {
		hostIsActive = false
	}

	enemy := playfield.EnemyPlayer()
	switch card.Effect {
	case "rollAgain":
		if position.CurrentStep == "afterRoll" {
			if playfield.ActivePlayer.Mana >= 4 {
				playfield.LastRoll = playfield.ActivePlayer.Die.Throw()
				Database.UpdateLastRollGames(playfield)
				playfield.ActivePlayer.Mana = playfield.ActivePlayer.Mana - 4
			} else {
				return nil, nil
			}
		} else {

		}
	case "doubleMana":
		if playfield.ActivePlayer.Mana >= 3 {
			if hostIsActive {
				position.HostInfo = "doubleMana"
			} else {
				position.GuestInfo = "doubleMana"
			}
			Database.UpdatePosition(position)
			playfield.ActivePlayer.Mana = playfield.ActivePlayer.Mana - 3
		} else {
			return nil, nil
		}
	case "flipClockwise":

		if playfield.ActivePlayer.Mana >= 6 {
			playfield.ActivePlayer.Grid.FlipClocwise()
			enemy.Grid.PrettyPrint()
			enemy.Grid = playfield.ActivePlayer.Grid.CheckGridForOverlap(enemy.Grid)

			Database.UpdateColumn(playfield.ActivePlayer.Grid.Left)
			Database.UpdateColumn(playfield.ActivePlayer.Grid.Middle)
			Database.UpdateColumn(playfield.ActivePlayer.Grid.Right)
			Database.UpdateColumn(enemy.Grid.Left)
			Database.UpdateColumn(enemy.Grid.Middle)
			Database.UpdateColumn(enemy.Grid.Right)
			playfield.ActivePlayer.Mana = playfield.ActivePlayer.Mana - 6
		} else {
			return nil, nil
		}
	case "destroyColumn":
		if playfield.ActivePlayer.Mana >= 7 {
			log.Println("destroyColumn")
			column := rand.IntN(3)
			switch column {
			case 0:
				enemy.Grid.Left.Clear()
				err := Database.UpdateColumn(enemy.Grid.Left)
				if err != nil {
					log.Println(err)
					return nil, nil
				}
			case 1:
				enemy.Grid.Middle.Clear()
				err := Database.UpdateColumn(enemy.Grid.Middle)
				if err != nil {
					log.Println(err)
					return nil, nil
				}
			case 2:
				enemy.Grid.Right.Clear()
				err := Database.UpdateColumn(enemy.Grid.Right)
				if err != nil {
					log.Println(err)
					return nil, nil
				}
			}
			playfield.ActivePlayer.Mana = playfield.ActivePlayer.Mana - 7
		} else {
			return nil, nil
		}
	}

	Database.UpdatePlayerMana(playfield.ActivePlayer)

	card.Played = true
	card.InHand = false
	Database.UpdateCard(card)
	if enemy.UserID == playfield.Host.UserID {
		playfield.Host = enemy

	} else {
		playfield.Guest = enemy
	}
	for cardIndex := range playfield.ActivePlayer.Deck.Cards {
		if playfield.ActivePlayer.Deck.Cards[cardIndex].CardID == card.CardID {
			playfield.ActivePlayer.Deck.Cards[cardIndex] = card
		}
	}

	if playfield.ActivePlayer.UserID == playfield.Host.UserID {
		playfield.Host = playfield.ActivePlayer
		hostIsActive = true
	} else {
		log.Println("GUEST WAS ACTIVE")
		playfield.Guest = playfield.ActivePlayer
		hostIsActive = false
	}

	var hostMsg = make(map[string]string)
	var guestMsg = make(map[string]string)
	hostMsg["id"] = playfield.Host.WebsocketConnectionID
	newMessage := `{"gameid": "` + playfield.GameID + `", "YourInfo":` + playfield.Host.ToJson(true) + `, "EnemyInfo": ` + playfield.Guest.ToJson(false) + `, "ActivePlayer": {"active": ` + strconv.FormatBool(hostIsActive) + `, "roll": "` + playfield.LastRoll + `"}}`
	infoMessage := `{"info": "gameInfo", "message": {"gameInfo": ` + newMessage + `}}`
	hostMsg["message"] = infoMessage

	guestMsg["id"] = playfield.Guest.WebsocketConnectionID
	newMessage = `{"gameid": "` + playfield.GameID + `", "YourInfo": ` + playfield.Guest.ToJson(true) + `, "EnemyInfo":` + playfield.Host.ToJson(false) + `, "ActivePlayer": {"active": ` + strconv.FormatBool(!hostIsActive) + `, "roll": "` + playfield.LastRoll + `"}}`
	infoMessage = `{"info": "gameInfo", "message": {"gameInfo": ` + newMessage + `}}`
	guestMsg["message"] = infoMessage

	return hostMsg, guestMsg
}
