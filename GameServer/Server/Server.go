package Server

import (
	"RollsOfDestiny/GameServer/Types"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var c = make(chan *websocket.Conn, 5) //5 is an arbitrary buffer size
var c2 = make(chan map[string]string, 5)

func queueForGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println(r.Method)
	if r.Method == "OPTIONS" {
		fmt.Println("OPTIONS request")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // You can add more headers here if needed
		w.Header().Set("Access-Control-Allow-Methods", "*")
		return
	}
	if r.Method == "POST" {
		// Read the raw body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		fmt.Printf("Raw body: %s\n", body)

		var t Types.QueueInfo

		fmt.Println(string(body))

		err = json.Unmarshal(body, &t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		/*err = GameLogic.PickColumn(t.Gameid, t.ColumnKey)

		if err != nil {
			panic(err)
		}
		*/
		log.Printf("Received gameid: %s\n", t.UserId)
		log.Printf("Received column key: %s\n", t.WebsocketConnectionId)

		AddToQueue(t, &c2)

		w.WriteHeader(http.StatusOK)
	}
}

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

	reader(ws, &c2)
}

func setupRoutes() {
	fmt.Println("handle something")
	http.HandleFunc("/queue", queueForGame)
	http.HandleFunc("/ws", wsEndpoint)
}

func Server() {

	fmt.Println("starting")
	setupRoutes()
	go func() {
		var somekindofstorrage = map[string]*websocket.Conn{}
		for {
			select {
			case newC := <-c:
				somekindofstorrage[strings.Split(newC.RemoteAddr().String(), ":")[len(strings.Split(newC.RemoteAddr().String(), ":"))-1]] = newC
			case msg := <-c2:
				fmt.Println("s", msg["id"])
				fmt.Println("s", msg["message"])
				err := somekindofstorrage[msg["id"]].WriteMessage(1, []byte(msg["message"]))
				fmt.Println(err)
			}
		}
	}()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
