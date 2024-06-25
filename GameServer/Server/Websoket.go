package Server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"strings"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type msg struct {
	Session string `json:"session"`
}

func reader(conn *websocket.Conn, c2 chan string) {
	conn.WriteMessage(1, []byte("connected"))
	//fmt.Printf("test")
	for {
		//fmt.Printf(conn.RemoteAddr())
		//
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("test3")

		log.Println(string(p))
		if strings.Contains(string(p), "session") {

			var msg msg
			err = json.Unmarshal(p, &msg)

			log.Println(string(msg.Session))
			c2 <- string(msg.Session)

		} else {
			c2 <- string(conn.RemoteAddr().String())
		}

		//conn.WriteMessage(messageType, []byte("testasdfas"))

		/*if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}*/
		//fmt.Printf("test2")
	}
}
