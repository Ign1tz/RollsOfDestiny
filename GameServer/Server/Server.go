package Server

import (
	"RollsOfDestiny/GameServer/Database"
	"RollsOfDestiny/GameServer/GameLogic"
	"RollsOfDestiny/GameServer/Types"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var c = make(chan *websocket.Conn, 50) //5 is an arbitrary buffer size
var c2 = make(chan map[string]string, 50)

func startBot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
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

		var t Types.BotResp

		err = json.Unmarshal(body, &t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		GameLogic.BotStartGame(t, &c2)
	}
}

func playBot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
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

		var t Types.Resp

		err = json.Unmarshal(body, &t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		GameLogic.BotTurn(t)
	}
}

func queueForGameWithFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
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

		var t Types.QueueInfoFriend

		err = json.Unmarshal(body, &t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		AddToFriendQueue(t, &c2)

		w.WriteHeader(http.StatusOK)
	}
}

func queueForGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
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

		err = json.Unmarshal(body, &t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		AddToQueue(t, &c2)

		w.WriteHeader(http.StatusOK)
	}
}

func removeFromQueue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
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

		err = json.Unmarshal(body, &t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		Database.DeleteFromQueueWebsocket(t.WebsocketConnectionId)

		w.WriteHeader(http.StatusOK)
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	c <- ws
	if err != nil {
		log.Println(err)
	}

	//log.Println("Client Successfully Connected")
	reader(ws, &c2)
	//log.Println("Websocket Closed")
	Database.DeleteFromQueueWebsocket(strings.Split(ws.RemoteAddr().String(), ":")[len(strings.Split(ws.RemoteAddr().String(), ":"))-1])
}

func setupRoutes() {
	http.HandleFunc("/queue", queueForGame)
	http.HandleFunc("/deleteQueue", removeFromQueue)
	http.HandleFunc("/queueFroGameWithFriend", queueForGameWithFriend)
	http.HandleFunc("/ws", wsEndpoint)
	http.HandleFunc("/startBot", startBot)
}

func Server() {
	setupRoutes()
	go func() {
		var somekindofstorrage = map[string]*websocket.Conn{}
		for {
			select {
			case newC := <-c:
				somekindofstorrage[strings.Split(newC.RemoteAddr().String(), ":")[len(strings.Split(newC.RemoteAddr().String(), ":"))-1]] = newC
			case msg := <-c2:
				if msg != nil && somekindofstorrage[msg["id"]] != nil {
					fmt.Println("s", msg["id"])
					fmt.Println("s", msg["message"])
					err := somekindofstorrage[msg["id"]].WriteMessage(1, []byte(msg["message"]))
					if err != nil {
						log.Println(err)
						_ = err
					}
				} else {
					log.Println("message or storrage is empty", msg["id"], somekindofstorrage[msg["id"]])
				}
			}
		}
	}()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
