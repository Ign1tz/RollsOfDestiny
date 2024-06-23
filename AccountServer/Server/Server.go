package Server

import (
	"RollsOfDestiny/AccountServer/SignUpLogic"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
)

var c = make(chan *websocket.Conn, 5) //5 is an arbitrary buffer size
var c2 = make(chan string, 5)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homag Page")
}

func signUp(w http.ResponseWriter, r *http.Request) {
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

		var t SignUpLogic.SignUpInfo

		fmt.Println(string(body))

		err = json.Unmarshal(body, &t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("username", t.Username)
		fmt.Println("password", t.Password)
		fmt.Println("confirm password", t.ConfirmPassword)
		fmt.Println("email", t.Email)
		w.WriteHeader(http.StatusOK)
	}
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/signup", signUp)
}

func Server() {
	fmt.Println("starting")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":9090", nil))
}
