package Server

import (
	"RollsOfDestiny/GameServer/Database"
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

type websocketMessage struct {
	Type        string `json:"type"`
	MessageBody string `json:"messageBody"`
	GameId      string `json:"gameId"`
}

func reader(conn *websocket.Conn, c2 *chan map[string]string) {
	conn.WriteMessage(1, []byte(`{"info": "connected", "message": {"connected": "true"}}`))
	connectionID := strings.Split(conn.RemoteAddr().String(), ":")[len(strings.Split(conn.RemoteAddr().String(), ":"))-1]
	//fmt.Printf("test")
	for {
		fmt.Println("reader start")
		//fmt.Printf(conn.RemoteAddr())
		//
		_, p, err := conn.ReadMessage()
		fmt.Println("message recived")
		if err != nil {
			log.Println(err)
			return
		}
		/*
			if string(p) == "id" {
				log.Println(string(p))
				conn.WriteMessage(1, []byte("id:"+strings.Split(conn.RemoteAddr().String(), ":")[len(strings.Split(conn.RemoteAddr().String(), ":"))-1]))
			}

			var msg = make(map[string]string)
			msg["id"] = connectionID
			msg["message"] = string(conn.RemoteAddr().String())*/

		log.Println(string(p))

		var message websocketMessage

		err = json.Unmarshal(p, &message)
		if err != nil {
			log.Println(err)
			log.Println(string(p))
			return
		}
		log.Println(message.Type)
		if message.Type == "botPickColumn" {
			ended := GameLogic.BotTurn(
				Types.Resp{Gameid: message.GameId, ColumnKey: message.MessageBody})

			playfield, err := Database.GetPlayfield(message.GameId)

			if err != nil {
				log.Println(err)
				return
			}
			var hostMsg = make(map[string]string)
			log.Println(ended)
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
				log.Println(playfield.Host.WebsocketConnectionID)
				hostMsg["id"] = playfield.Host.WebsocketConnectionID
				newMessage := `{"gameid": "` + playfield.GameID + `", "YourInfo":` + playfield.Host.ToJson(true) + `, "EnemyInfo": ` + playfield.Guest.ToJson(false) + `, "ActivePlayer": {"active": ` + strconv.FormatBool(true) + `, "roll": "` + playfield.LastRoll + `"}}`
				infoMessage := `{"info": "gameInfo", "message": {"gameInfo": ` + newMessage + `}}`
				hostMsg["message"] = infoMessage
				*c2 <- hostMsg
			}
		} else {

			msg, msg2 := categorizeMessage(message, connectionID)
			*c2 <- msg
			if msg2 != nil {
				*c2 <- msg2
			}
		}

		log.Println("reader end:", string(p))
	}
}

func categorizeMessage(message websocketMessage, connectionId string) (map[string]string, map[string]string) {
	var msg = make(map[string]string)
	switch message.Type {
	case "id":
		msg["id"] = connectionId
		msg["message"] = `{"info": "id", "message": {"id": "` + connectionId + `"}}`
	case "PickColumn":
		return handlePickedColumn(message)
	case "surrender":
		playfield, err := Database.GetPlayfield(message.GameId)
		if err != nil {
			log.Println(err)
			return nil, nil
		}
		return handleGameEnded(playfield)
	}
	return msg, nil
}

func handlePickedColumn(message websocketMessage) (map[string]string, map[string]string) {
	playfield, err := Database.GetPlayfield(message.GameId)
	if err != nil {
		panic(err)
	}
	enemy := playfield.EnemyPlayer()
	columnInt, _ := strconv.Atoi(playfield.LastRoll)
	switch message.MessageBody {
	case "0":
		err := playfield.ActivePlayer.Grid.Left.Add(columnInt)
		if err != nil {
			panic(err)
		}
		enemy.Grid.Left.Remove(columnInt)
		Database.UpdateColumn(playfield.ActivePlayer.Grid.Left)
		Database.UpdateColumn(enemy.Grid.Left)
	case "1":
		err := playfield.ActivePlayer.Grid.Middle.Add(columnInt)
		if err != nil {
			panic(err)
		}
		enemy.Grid.Middle.Remove(columnInt)
		fmt.Println("websocket placement", playfield.ActivePlayer.Grid.Middle.Placement)
		Database.UpdateColumn(playfield.ActivePlayer.Grid.Middle)
		Database.UpdateColumn(enemy.Grid.Middle)
	case "2":
		err := playfield.ActivePlayer.Grid.Right.Add(columnInt)
		if err != nil {
			panic(err)
		}
		enemy.Grid.Right.Remove(columnInt)
		Database.UpdateColumn(playfield.ActivePlayer.Grid.Right)
		Database.UpdateColumn(enemy.Grid.Right)
	default:
		return nil, nil
	}
	var hostIsActive bool
	if playfield.ActivePlayer.UserID == playfield.Host.UserID {
		log.Println("HOST WAS ACTIVE")
		playfield.Host = playfield.ActivePlayer
		playfield.Guest = enemy
		hostIsActive = false
	} else {
		log.Println("GUEST WAS ACTIVE")
		playfield.Guest = playfield.ActivePlayer
		playfield.Host = enemy
		hostIsActive = true
	}
	gameEnded := playfield.ActivePlayer.Grid.IsFull()
	playfield.ActivePlayer = playfield.EnemyPlayer()

	playfield.LastRoll = playfield.Host.Die.Throw()

	Database.UpdateActivePlayerGames(playfield)
	Database.UpdateLastRollGames(playfield)

	fmt.Println(playfield.Host.Grid.Left.First)

	var hostMsg = make(map[string]string)
	var guestMsg = make(map[string]string)
	if gameEnded {
		fmt.Println("game ended")
		hostMsg, guestMsg = handleGameEnded(playfield)
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

func handleGameEnded(playfield Types.Playfield) (map[string]string, map[string]string) {
	var hostMsg = make(map[string]string)
	var guestMsg = make(map[string]string)

	hostWon := playfield.Host.Grid.Value() > playfield.Guest.Grid.Value()
	guestWon := playfield.Host.Grid.Value() < playfield.Guest.Grid.Value()
	tie := playfield.Host.Grid.Value() == playfield.Guest.Grid.Value()

	var hostWonMessage string
	var guestWonMessage string

	if tie {
		hostWonMessage = "Its a Tie!"
		guestWonMessage = "Its a Tie!"
	} else if guestWon {
		guestWonMessage = "You Won!"
		hostWonMessage = "You Lost!"
	} else if hostWon {
		hostWonMessage = "You Won!"
		guestWonMessage = "You Lost!"
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
