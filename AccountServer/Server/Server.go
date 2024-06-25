package Server

import (
	"RollsOfDestiny/AccountServer/SignUpLogic"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var c = make(chan *websocket.Conn, 5) //5 is an arbitrary buffer size
var c2 = make(chan string, 5)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

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
		SignUpLogic.SignUpNewAccount(t)
		w.WriteHeader(http.StatusOK)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
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

		var t SignUpLogic.LoginInfo

		fmt.Println(string(body))

		err = json.Unmarshal(body, &t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("username", t.Username)
		fmt.Println("password", t.Password)
		if SignUpLogic.LoginToAccount(t) {
			tokenString, err := createToken(t.Username)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Errorf("No username found")

				return
			}

			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, tokenString)
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid credentials")
	}
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/login", login)
}

func Server() {
	fmt.Println("starting")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":9090", nil))
}
