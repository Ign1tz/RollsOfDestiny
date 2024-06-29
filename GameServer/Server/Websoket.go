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
			return
		}
		msg := categorizeMessage(message, connectionID)
		*c2 <- msg

		log.Println("reader end:", string(p))
	}
}

func categorizeMessage(message websocketMessage, connectionId string) map[string]string {
	var msg = make(map[string]string)
	switch message.Type {
	case "id":
		msg["id"] = connectionId
		msg["message"] = `{"info": "id", "message": {"id": "` + connectionId + `"}}`
	case "pickColumn":
		handlePickedColumn(message)
	}
	return msg
}

func handlePickedColumn(message websocketMessage) {
	playfield, err := Database.GetPlayfield(message.GameId)
	if err != nil {
		panic(err)
	}
	enemy := playfield.EnemyPlayer()
	switch message.MessageBody {
	case "0":
		columnInt, _ := strconv.Atoi(playfield.LastRoll)
		err := playfield.ActivePlayer.Grid.Left.Add(columnInt)
		if err != nil {
			panic(err)
		}
		enemy.Grid.Left.Remove(columnInt)
	case "1":
		columnInt, _ := strconv.Atoi(playfield.LastRoll)
		err := playfield.ActivePlayer.Grid.Middle.Add(columnInt)
		if err != nil {
			panic(err)
		}
		enemy.Grid.Middle.Remove(columnInt)
	case "2":
		columnInt, _ := strconv.Atoi(playfield.LastRoll)
		err := playfield.ActivePlayer.Grid.Right.Add(columnInt)
		if err != nil {
			panic(err)
		}
		enemy.Grid.Right.Remove(columnInt)
	default:
		return
	}
}
