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
	conn.WriteMessage(1, []byte("connected"))
	//fmt.Printf("test")
	for {
		//fmt.Printf(conn.RemoteAddr())
		//
		fmt.Println("test3")
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("test3")
		if string(p) == "id" {
			log.Println(string(p))
			conn.WriteMessage(1, []byte("id:"+conn.RemoteAddr().String()))
		}

		fmt.Println("test3")

		var msg = make(map[string]string)
		msg["id"] = strings.Split(conn.RemoteAddr().String(), ":")[len(strings.Split(conn.RemoteAddr().String(), ":"))-1]
		msg["message"] = string(conn.RemoteAddr().String())

		*c2 <- msg

		log.Println("reader end:", string(p))

		//conn.WriteMessage(messageType, []byte("testasdfas"))

		/*if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}*/
		//fmt.Printf("test2")
	}
}

func categorizeMessage(message websocketMessage, connectionId string) {
	var msg = make(map[string]string)
	switch message.Type {
	case "id":
		msg["id"] = connectionId
		msg["message"] = `{"id": "` + connectionId + `"}`
	case "pickColumn":
		handlePickedColumn(message)
	}
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
