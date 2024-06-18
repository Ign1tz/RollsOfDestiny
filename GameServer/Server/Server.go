package Server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var c = make(chan *websocket.Conn, 5) //5 is an arbitrary buffer size
var c2 = make(chan string, 5)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homag Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	c <- ws
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Successfully Connected")

	reader(ws, c2)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func Server() {
	fmt.Printf("starting")
	setupRoutes()
	go func() {
		var somekindofstorrage = map[string]*websocket.Conn{}
		for {
			select {
			case newC := <-c:
				somekindofstorrage[newC.RemoteAddr().String()] = newC
			case msg := <-c2:
				somekindofstorrage[msg].WriteMessage(1, []byte("if i see this it maybe works"+msg))
			}
		}
	}()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
