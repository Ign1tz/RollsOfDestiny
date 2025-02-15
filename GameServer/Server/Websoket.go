package Server

import (
	AccouuntDatabase "RollsOfDestiny/AccountServer/Database"
	Database "RollsOfDestiny/GameServer/Database"
	"RollsOfDestiny/GameServer/GameLogic"
	"RollsOfDestiny/GameServer/Types"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"strings"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn, c2 *chan map[string]string) {
	conn.WriteMessage(1, []byte(`{"info": "connected", "message": {"connected": "true"}}`))
	connectionID := strings.Split(conn.RemoteAddr().String(), ":")[len(strings.Split(conn.RemoteAddr().String(), ":"))-1]
	//fmt.Printf("test")
	for {
		fmt.Println("reader start")
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("incoming message", string(p))

		var message Types.WebsocketMessage

		err = json.Unmarshal(p, &message)
		if err != nil {
			log.Println(err)
			return
		}
		if message.Type == "botPickColumn" {
			ended := GameLogic.BotTurn(
				Types.Resp{Gameid: message.GameId, ColumnKey: message.MessageBody})

			playfield, err := Database.GetPlayfield(message.GameId)

			if err != nil {
				log.Println(err)
				return
			}
			var hostMsg = make(map[string]string)
			if ended {
				hostWon := playfield.Host.Grid.Value() > playfield.Guest.Grid.Value()
				guestWon := playfield.Host.Grid.Value() < playfield.Guest.Grid.Value()
				tie := playfield.Host.Grid.Value() == playfield.Guest.Grid.Value()

				var hostWonMessage string

				if tie {
					hostWonMessage = "Its a Tie!"
				} else if guestWon {
					hostWonMessage = "You Lost!"
				} else if hostWon {
					hostWonMessage = "You Won!"
				}

				hostMsg["id"] = playfield.Host.WebsocketConnectionID
				playfieldMessage := `{"gameid": "` + playfield.GameID + `", "YourInfo":` + playfield.Host.ToJson(true) + `, "EnemyInfo": ` + playfield.Guest.ToJson(false) + `, "ActivePlayer": {"active": ` + strconv.FormatBool(false) + `, "roll": "` + playfield.LastRoll + `"}}`
				newMessage := `{"yourScore": ` + strconv.Itoa(playfield.Host.Grid.Value()) + `, "enemyScore": ` + strconv.Itoa(playfield.Guest.Grid.Value()) + `, "youWon": "` + hostWonMessage + `"}`
				infoMessage := `{"info": "gameEnded", "message": {"gameInfo": ` + playfieldMessage + `, "endResults": ` + newMessage + `}}`
				hostMsg["message"] = infoMessage
				*c2 <- hostMsg
				Database.DeleteGame(playfield.Host.Grid.GridId)
				Database.DeleteGame(playfield.Guest.Grid.GridId)
			} else {
				hostMsg["id"] = playfield.Host.WebsocketConnectionID
				newMessage := `{"gameid": "` + playfield.GameID + `", "YourInfo":` + playfield.Host.ToJson(true) + `, "EnemyInfo": ` + playfield.Guest.ToJson(false) + `, "ActivePlayer": {"active": ` + strconv.FormatBool(true) + `, "roll": "` + playfield.LastRoll + `"}}`
				infoMessage := `{"info": "gameInfo", "message": {"gameInfo": ` + newMessage + `}}`
				hostMsg["message"] = infoMessage
				*c2 <- hostMsg
			}
		} else {
			log.Println("before message Categorization")
			msg, msg2 := categorizeMessage(message, connectionID)
			*c2 <- msg
			if msg2 != nil {
				*c2 <- msg2
			}
		}

		log.Println("reader end:", string(p))
	}
}

func newmessage(gameid string) (map[string]string, map[string]string) {
	playfield, err := Database.GetPlayfield(gameid)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	hostIsActive := playfield.Host.UserID == playfield.ActivePlayer.UserID

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

func categorizeMessage(message Types.WebsocketMessage, connectionId string) (map[string]string, map[string]string) {

	var msg = make(map[string]string)
	if message.Type == "id" {
		msg["id"] = connectionId
		msg["message"] = `{"info": "id", "message": {"id": "` + connectionId + `"}}`
		return msg, nil
	}
	playfield, err := Database.GetPlayfield(message.GameId)
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	if message.Type == "surrender" {
		if strings.Contains(playfield.GameID, "bot: ") {
			Database.DeleteGame(playfield.Host.Grid.GridId)
			Database.DeleteGame(playfield.Guest.Grid.GridId)
			return nil, nil
		}
		var surenderer Types.Player
		if playfield.Host.WebsocketConnectionID == connectionId {
			playfield.Host.Grid.Clear()
			surenderer = playfield.Host
		} else {
			playfield.Guest.Grid.Clear()
			surenderer = playfield.Guest
		}
		return handleGameEnded(playfield, true, surenderer)
	}
	if playfield.ActivePlayer.UserID != message.Userid {
		log.Println("not active player", err)
		return nil, nil
	}

	position, err := Database.GetPosition(message.GameId)
	if err != nil {
		log.Println("position", err)
		return nil, nil
	}
	log.Println("messageType", message.Type)
	switch message.Type {
	case "PickColumn":
		if position.CurrentStep == "afterRoll" {
			msg1, msg2 := handlePickedColumn(message)
			position.CurrentStep = "afterColumnPick"
			err = Database.UpdatePosition(position)
			return msg1, msg2
		} else {
			return nil, nil
		}
	case "FriendPickColumn":
		if position.CurrentStep == "afterRoll" {
			msg1, msg2 := handlePickedColumn(message)
			position.CurrentStep = "afterColumnPick"
			err = Database.UpdatePosition(position)
			return msg1, msg2
		} else {
			return nil, nil
		}
	case "playCard":
		if position.CurrentStep == "afterRoll" {
			return GameLogic.HandleCards(message, position)
		} else {
			return nil, nil
		}
	case "endTurn":
		if position.CurrentStep == "afterColumnPick" {
			msg1, msg2 := handleEndTurn(message)
			position.CurrentStep = "start"
			err = Database.UpdatePosition(position)
			return msg1, msg2

		} else {
			return nil, nil
		}
	case "rolled":
		position.CurrentStep = "afterRoll"
		err := Database.UpdatePosition(position)
		if err != nil {
			log.Println(err)
			return nil, nil
		}
		return nil, nil
	}
	return msg, nil
}

func handleEndTurn(message Types.WebsocketMessage) (map[string]string, map[string]string) {
	playfield, err := Database.GetPlayfield(message.GameId)
	if err != nil {
		panic(err)
	}

	var hostIsActive bool
	if playfield.ActivePlayer.UserID == playfield.Host.UserID {
		hostIsActive = false
	} else {
		hostIsActive = true
	}

	gameEnded := playfield.ActivePlayer.Grid.IsFull()
	playfield.ActivePlayer = playfield.EnemyPlayer()

	playfield.LastRoll = playfield.Host.Die.Throw()

	Database.UpdateActivePlayerGames(playfield)
	Database.UpdateLastRollGames(playfield)

	var hostMsg = make(map[string]string)
	var guestMsg = make(map[string]string)
	if gameEnded {
		hostMsg, guestMsg = handleGameEnded(playfield, false, Types.Player{})
	} else {
		hostMsg["id"] = playfield.Host.WebsocketConnectionID
		newMessage := `{"gameid": "` + playfield.GameID + `", "YourInfo":` + playfield.Host.ToJson(true) + `, "EnemyInfo": ` + playfield.Guest.ToJson(false) + `, "ActivePlayer": {"active": ` + strconv.FormatBool(hostIsActive) + `, "roll": "` + playfield.LastRoll + `"}}`
		infoMessage := `{"info": "gameInfo", "message": {"gameInfo": ` + newMessage + `}}`
		hostMsg["message"] = infoMessage

		guestMsg["id"] = playfield.Guest.WebsocketConnectionID
		newMessage = `{"gameid": "` + playfield.GameID + `", "YourInfo": ` + playfield.Guest.ToJson(true) + `, "EnemyInfo":` + playfield.Host.ToJson(false) + `, "ActivePlayer": {"active": ` + strconv.FormatBool(!hostIsActive) + `, "roll": "` + playfield.LastRoll + `"}}`
		infoMessage = `{"info": "gameInfo", "message": {"gameInfo": ` + newMessage + `}}`
		guestMsg["message"] = infoMessage
	}
	return hostMsg, guestMsg
}

func handlePickedColumn(message Types.WebsocketMessage) (map[string]string, map[string]string) {
	playfield, err := Database.GetPlayfield(message.GameId)
	if err != nil {
		panic(err)
	}
	position, err := Database.GetPosition(message.GameId)
	if err != nil {
		panic(err)
	}

	enemy := playfield.EnemyPlayer()
	columnInt, _ := strconv.Atoi(playfield.LastRoll)
	numberOfRemoved := 0
	switch message.MessageBody {
	case "0":
		err := playfield.ActivePlayer.Grid.Left.Add(columnInt)
		if err != nil {
			log.Println(err)
			position.CurrentStep = "afterColumnPick"
			return newmessage(message.GameId)
		} else {
			numberOfRemoved = enemy.Grid.Left.Remove(columnInt)
			err = Database.UpdateColumn(playfield.ActivePlayer.Grid.Left)
			if err != nil {
				return nil, nil
			}
			err = Database.UpdateColumn(enemy.Grid.Left)
			if err != nil {
				return nil, nil
			}
		}
	case "1":
		err := playfield.ActivePlayer.Grid.Middle.Add(columnInt)
		if err != nil {
			position.CurrentStep = "afterColumnPick"
			log.Println(err)
			return newmessage(message.GameId)
		} else {
			numberOfRemoved = enemy.Grid.Middle.Remove(columnInt)
			err = Database.UpdateColumn(playfield.ActivePlayer.Grid.Middle)
			if err != nil {
				return nil, nil
			}
			err = Database.UpdateColumn(enemy.Grid.Middle)
			if err != nil {
				return nil, nil
			}
		}
	case "2":
		err := playfield.ActivePlayer.Grid.Right.Add(columnInt)
		if err != nil {
			log.Println(err)
			position.CurrentStep = "afterColumnPick"
			return newmessage(message.GameId)
		} else {
			numberOfRemoved = enemy.Grid.Right.Remove(columnInt)
			err = Database.UpdateColumn(playfield.ActivePlayer.Grid.Right)
			if err != nil {
				return nil, nil
			}
			err = Database.UpdateColumn(enemy.Grid.Right)
			if err != nil {
				return nil, nil
			}
		}
	default:
		return nil, nil
	}

	var hostIsActive bool
	if playfield.ActivePlayer.UserID == playfield.Host.UserID {
		playfield.Host = playfield.ActivePlayer
		playfield.Guest = enemy
		addMana := 1
		if position.GuestInfo == "doubleMana" {
			addMana = 2
			numberOfRemoved *= 2
			position.GuestInfo = ""
			err := Database.UpdatePosition(position)
			if err != nil {
				log.Println("updatePosition", err)
			}
		}
		playfield.Guest.Mana = min(max(playfield.Guest.Mana+addMana+numberOfRemoved, 0), 10)
		hostIsActive = false
		if playfield.Guest.Deck.DeckID != "" {
			handCards := 0
			for cardIndex := range playfield.Guest.Deck.Cards {
				if playfield.Guest.Deck.Cards[cardIndex].InHand {
					handCards++
				}
			}
			if handCards != 4 {
				cards, err := Database.GetCardsByDeckID(playfield.Guest.Deck.DeckID)
				if err != nil {
					log.Println(err)
					return nil, nil
				}
				cards = RandShuffle(cards)
				index := 0
				for i := 0; i < 4-handCards; i++ {
					if len(cards) > index {
						if !cards[index].InHand {
							cards[index].InHand = true
							err := Database.UpdateCard(cards[index])
							if err != nil {
								log.Println(err)
								return nil, nil
							}
						} else {
							i--
						}
						index++
					}
				}
				playfield.Guest.Deck.Cards = cards
			}
		}
	} else {
		playfield.Guest = playfield.ActivePlayer
		playfield.Host = enemy
		addMana := 1
		if position.HostInfo == "doubleMana" {
			addMana = 2
			numberOfRemoved *= 2
			position.HostInfo = ""

			err := Database.UpdatePosition(position)
			if err != nil {
				log.Println("updatePosition", err)
			}
		}
		playfield.Host.Mana = min(max(playfield.Host.Mana+addMana+numberOfRemoved, 0), 10)
		hostIsActive = true
		if playfield.Host.Deck.DeckID != "" {
			handCards := 0
			for cardIndex := range playfield.Host.Deck.Cards {
				if playfield.Host.Deck.Cards[cardIndex].InHand {
					handCards++
				}
			}
			if handCards != 4 {
				cards, err := Database.GetCardsByDeckID(playfield.Host.Deck.DeckID)
				if err != nil {
					log.Println(err)
					return nil, nil
				}
				cards = RandShuffle(cards)
				index := 0
				for i := 0; i < 4-handCards; i++ {
					if len(cards) > index {
						if !cards[index].InHand {
							cards[index].InHand = true
							err := Database.UpdateCard(cards[index])
							if err != nil {
								log.Println(err)
								return nil, nil
							}
						} else {
							i--
						}
						index++
					}
				}
				playfield.Host.Deck.Cards = cards
			}
		}

	}
	gameEnded := playfield.ActivePlayer.Grid.IsFull()
	playfield.ActivePlayer = playfield.EnemyPlayer()
	playfield.LastRoll = playfield.Host.Die.Throw()

	err = Database.UpdatePlayerMana(playfield.ActivePlayer)
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	err = Database.UpdateActivePlayerGames(playfield)
	if err != nil {
		log.Println(err)
		return nil, nil

	}
	err = Database.UpdateLastRollGames(playfield)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	var hostMsg = make(map[string]string)
	var guestMsg = make(map[string]string)
	if gameEnded {
		hostMsg, guestMsg = handleGameEnded(playfield, false, Types.Player{})
	} else {
		hostMsg["id"] = playfield.Host.WebsocketConnectionID
		newMessage := `{"gameid": "` + playfield.GameID + `", "YourInfo":` + playfield.Host.ToJson(true) + `, "EnemyInfo": ` + playfield.Guest.ToJson(false) + `, "ActivePlayer": {"active": ` + strconv.FormatBool(hostIsActive) + `, "roll": "` + playfield.LastRoll + `"}}`
		infoMessage := `{"info": "gameInfo", "message": {"gameInfo": ` + newMessage + `}}`
		hostMsg["message"] = infoMessage

		guestMsg["id"] = playfield.Guest.WebsocketConnectionID
		newMessage = `{"gameid": "` + playfield.GameID + `", "YourInfo": ` + playfield.Guest.ToJson(true) + `, "EnemyInfo":` + playfield.Host.ToJson(false) + `, "ActivePlayer": {"active": ` + strconv.FormatBool(!hostIsActive) + `, "roll": "` + playfield.LastRoll + `"}}`
		infoMessage = `{"info": "gameInfo", "message": {"gameInfo": ` + newMessage + `}}`
		guestMsg["message"] = infoMessage
	}
	return hostMsg, guestMsg
}

func handleGameEnded(playfield Types.Playfield, surrendered bool, surrenderer Types.Player) (map[string]string, map[string]string) {
	var hostMsg = make(map[string]string)
	var guestMsg = make(map[string]string)

	hostWon := playfield.Host.Grid.Value() > playfield.Guest.Grid.Value()
	guestWon := playfield.Host.Grid.Value() < playfield.Guest.Grid.Value()
	tie := playfield.Host.Grid.Value() == playfield.Guest.Grid.Value()

	if surrendered {
		hostWon = surrenderer.UserID == playfield.Guest.UserID
		guestWon = surrenderer.UserID == playfield.Host.UserID
		tie = false
	}

	var hostWonMessage string
	var guestWonMessage string

	if tie {
		hostWonMessage = "Its a Tie!"
		guestWonMessage = "Its a Tie!"
	} else if guestWon {
		guestWonMessage = "You Won!"
		hostWonMessage = "You Lost!"
		AccouuntDatabase.UpdateRating(playfield.Host.UserID, -10)
		AccouuntDatabase.UpdateRating(playfield.Guest.UserID, 10)
	} else if hostWon {
		hostWonMessage = "You Won!"
		guestWonMessage = "You Lost!"
		AccouuntDatabase.UpdateRating(playfield.Host.UserID, 10)
		AccouuntDatabase.UpdateRating(playfield.Guest.UserID, -10)
	}

	hostMsg["id"] = playfield.Host.WebsocketConnectionID

	playfieldMessage := `{"gameid": "` + playfield.GameID + `", "YourInfo":` + playfield.Host.ToJson(true) + `, "EnemyInfo": ` + playfield.Guest.ToJson(false) + `, "ActivePlayer": {"active": ` + strconv.FormatBool(false) + `, "roll": "` + playfield.LastRoll + `"}}`
	newMessage := `{"yourScore": ` + strconv.Itoa(playfield.Host.Grid.Value()) + `, "enemyScore": ` + strconv.Itoa(playfield.Guest.Grid.Value()) + `, "youWon": "` + hostWonMessage + `"}`
	infoMessage := `{"info": "gameEnded", "message": {"gameInfo": ` + playfieldMessage + `, "endResults": ` + newMessage + `}}`
	hostMsg["message"] = infoMessage

	if playfield.Guest.WebsocketConnectionID == "" {
		Database.DeleteGame(playfield.Host.Grid.GridId)
		Database.DeleteGame(playfield.Guest.Grid.GridId)
		return hostMsg, nil
	}

	guestMsg["id"] = playfield.Guest.WebsocketConnectionID
	playfieldMessage = `{"gameid": "` + playfield.GameID + `", "YourInfo": ` + playfield.Guest.ToJson(true) + `, "EnemyInfo":` + playfield.Host.ToJson(false) + `, "ActivePlayer": {"active": ` + strconv.FormatBool(false) + `, "roll": "` + playfield.LastRoll + `"}}`
	newMessage = `{"yourScore": ` + strconv.Itoa(playfield.Guest.Grid.Value()) + `, "enemyScore": ` + strconv.Itoa(playfield.Host.Grid.Value()) + `, "youWon": "` + guestWonMessage + `"}`
	infoMessage = `{"info": "gameEnded", "message": {"gameInfo": ` + playfieldMessage + `, "endResults": ` + newMessage + `}}`
	guestMsg["message"] = infoMessage

	Database.DeleteGame(playfield.Host.Grid.GridId)
	Database.DeleteGame(playfield.Guest.Grid.GridId)
	return hostMsg, guestMsg
}
