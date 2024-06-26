package Server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
)

var c = make(chan *websocket.Conn, 5) //5 is an arbitrary buffer size
var c2 = make(chan string, 5)

type BotResp struct {
	Userid string `json:"Userid"`
}

func startBot(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		fmt.Println("OPTIONS request")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // You can add more headers here if needed
		w.Header().Set("Access-Control-Allow-Methods", "*")
		return
	}

	if r.Method == "POST" {
		fmt.Println("POST request")

		// Read the raw body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		fmt.Printf("Raw body: %s\n", body)

		var t BotResp

		fmt.Println(string(body))

		err = json.Unmarshal(body, &t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	}
}

func playBot(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		fmt.Println("OPTIONS request")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // You can add more headers here if needed
		w.Header().Set("Access-Control-Allow-Methods", "*")
		return
	}

	if r.Method == "POST" {
		fmt.Println("POST request")

		// Read the raw body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		fmt.Printf("Raw body: %s\n", body)

		var t resp

		fmt.Println(string(body))

		err = json.Unmarshal(body, &t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homag Page")
}

type resp struct {
	Gameid    string `json:"gameid"`
	ColumnKey string `json:"columnKey"`
}

func pickColumn(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		fmt.Println("OPTIONS request")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // You can add more headers here if needed
		w.Header().Set("Access-Control-Allow-Methods", "*")
		return
	}

	if r.Method == "POST" {
		fmt.Println("POST request")

		// Read the raw body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		fmt.Printf("Raw body: %s\n", body)

		var t resp

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
		log.Printf("Received gameid: %s\n", t.Gameid)
		log.Printf("Received column key: %s\n", t.ColumnKey)
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
	log.Println("Client Successfully Connected")

	reader(ws, c2)
}

func setupRoutes() {
	http.HandleFunc("/ws", wsEndpoint)
	http.HandleFunc("/picKColumn", pickColumn)
}

func Server() {

	fmt.Printf("starting")
	fmt.Println("starting")
	setupRoutes()
	go func() {
		fmt.Println("this thing")
		var somekindofstorrage = map[string]*websocket.Conn{}
		for {
			select {
			case newC := <-c:
				somekindofstorrage[newC.RemoteAddr().String()] = newC
			case msg := <-c2:
				somekindofstorrage[msg].WriteMessage(1, []byte(msg))
			}
		}
	}()
}
