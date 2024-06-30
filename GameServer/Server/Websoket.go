package Server

import (
	"RollsOfDestiny/GameServer/Database"
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
		msg, msg2 := categorizeMessage(message, connectionID)
		*c2 <- msg
		if msg2 != nil {
			*c2 <- msg2
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
	case "pickColumn":
		return handlePickedColumn(message)
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
		playfield.Host = playfield.ActivePlayer
		hostIsActive = true
	} else {
		playfield.Guest = playfield.ActivePlayer
		hostIsActive = false
	}

	fmt.Println(playfield.Host.Grid.Left.First)
	var msg = make(map[string]string)
	msg["id"] = playfield.Host.WebsocketConnectionID
	newMessage := `{"gameid": "` + playfield.GameID + `", "YourInfo":` + playfield.Host.ToJson(true) + `, "EnemyInfo": ` + playfield.Guest.ToJson(false) + `, "ActivePlayer": {"active": "` + strconv.FormatBool(hostIsActive) + `", "roll": "` + playfield.LastRoll + `"}}`
	infoMessage := `{"info": "gameInfo", "message": {"gameInfo": ` + newMessage + `}}`
	msg["message"] = infoMessage

	var msg2 = make(map[string]string)
	msg2["id"] = playfield.Guest.WebsocketConnectionID
	newMessage = `{"gameid": "` + playfield.GameID + `", "YourInfo": ` + playfield.Guest.ToJson(true) + `, "EnemyInfo":` + playfield.Host.ToJson(false) + `, "ActivePlayer": {"active": "` + strconv.FormatBool(!hostIsActive) + `", "roll": "` + playfield.LastRoll + `"}}`
	infoMessage = `{"info": "gameInfo", "message": {"gameInfo": ` + newMessage + `}}`
	msg2["message"] = infoMessage

	return msg, msg2
}
