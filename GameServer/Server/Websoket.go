package Server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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
		c2 <- string(conn.RemoteAddr().String())

		log.Println(string(p))

		//conn.WriteMessage(messageType, []byte("testasdfas"))

		/*if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}*/
		//fmt.Printf("test2")
	}
}
