package Server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"strings"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn, c2 *chan map[string]string) {
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

		if string(p) == "id" {
			conn.WriteMessage(1, []byte("id:"+conn.RemoteAddr().String()))
		}

		fmt.Printf("test3")

		var msg = make(map[string]string)
		msg["id"] = strings.Split(conn.RemoteAddr().String(), ":")[len(strings.Split(conn.RemoteAddr().String(), ":"))-1]
		msg["message"] = string(conn.RemoteAddr().String())

		*c2 <- msg

		log.Println(string(p))

		//conn.WriteMessage(messageType, []byte("testasdfas"))

		/*if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}*/
		//fmt.Printf("test2")
	}
}
