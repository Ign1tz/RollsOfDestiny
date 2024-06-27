package Server

import (
	"RollsOfDestiny/AccountServer/Database"
	"RollsOfDestiny/AccountServer/SignUpLogic"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

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
			fmt.Println("token", tokenString)
			w.WriteHeader(http.StatusOK)
			test := `{"token": "` + tokenString + `"}`
			fmt.Fprint(w, test)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid credentials")
	}
}

func isLoggedIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "*") // You can add more headers here if needed
		w.Header().Set("Access-Control-Allow-Methods", "*")
		return
	}
	fmt.Println("isloggedin", r.Method)

	if checkToken(w, r) {
		fmt.Println("worked")
		w.WriteHeader(http.StatusOK)
	}
}

func accountInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "*") // You can add more headers here if needed
		w.Header().Set("Access-Control-Allow-Methods", "*")
		return
	}
	fmt.Println("aaaaaaaa", w.Header().Get("Authorization"))
	if checkToken(w, r) {
		myUrl, _ := url.Parse(r.URL.String())
		params, _ := url.ParseQuery(myUrl.RawQuery)
		fmt.Println(r.URL.String())
		account, err := Database.GetAccountByUsername(params.Get("username"))

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		accInfo := fmt.Sprintf("{\"username\": \"%s\", \"email\": \"%s\", \"profilePicture\": \"%s\", \"rating\": \"%s\"}", account.Username, account.Email, account.ProfilePicture, strconv.Itoa(account.Rating))
		fmt.Fprint(w, accInfo)
	}
}

func refresh(w http.ResponseWriter, r *http.Request) {
	fmt.Println("refresh")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/login", login)
	http.HandleFunc("/isLoggedIn", isLoggedIn)
	http.HandleFunc("/refresh", refresh)
	http.HandleFunc("/userInfo", accountInfo)

}

func Server() {
	fmt.Println("starting")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":9090", nil))
}
