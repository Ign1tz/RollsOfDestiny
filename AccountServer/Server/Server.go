package Server

import (
	"RollsOfDestiny/AccountServer/AccountLogic"
	"RollsOfDestiny/AccountServer/Database"
	"RollsOfDestiny/AccountServer/SignUpLogic"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homag Page")
}

func signUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
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

		var t SignUpLogic.SignUpInfo

		err = json.Unmarshal(body, &t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		SignUpLogic.SignUpNewAccount(t)
		w.WriteHeader(http.StatusOK)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "*") // You can add more headers here if needed
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

		var t SignUpLogic.LoginInfo

		err = json.Unmarshal(body, &t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if SignUpLogic.LoginToAccount(t) {
			account, err := Database.GetAccountByUsername(t.Username)
			if err != nil {
				log.Println(err)
				return
			}
			tokenString, err := createToken(account.UserID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Errorf("No username found")
				return
			}
			w.WriteHeader(http.StatusOK)
			fmt.Println("logged in")
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
	_, valid := checkToken(w, r)
	if valid {
		w.WriteHeader(http.StatusOK)
	}
}

func accountInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "*") // You can add more headers here if needed
		w.Header().Set("Access-Control-Allow-Methods", "*")
		return
	}
	userid, valid := checkToken(w, r)
	if valid {
		account, err := Database.GetAccountByUserID(userid)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		accInfo := fmt.Sprintf("{\"username\": \"%s\", \"email\": \"%s\", \"profilePicture\": \"%s\", \"rating\": \"%s\", \"userid\": \"%s\"}", account.Username, account.Email, account.ProfilePicture, strconv.Itoa(account.Rating), account.UserID)
		fmt.Fprint(w, accInfo)
	}
}

func changeUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "*") // You can add more headers here if needed
		w.Header().Set("Access-Control-Allow-Methods", "*")
		return
	}

	if r.Method == "POST" {
		userid, valid := checkToken(w, r)
		if valid {
			// Read the raw body
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer r.Body.Close()

			fmt.Printf("Raw body: %s\n", body)

			var t AccountLogic.NewUsernameMessage

			err = json.Unmarshal(body, &t)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			account, err := Database.GetAccountByUserID(userid)
			if err != nil {
				log.Println(err)
				return
			}

			t.OldUsername = account.Username
			AccountLogic.ChangeUsername(t)
			w.WriteHeader(http.StatusOK)
		}
	}
}

func changePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "*") // You can add more headers here if needed
		w.Header().Set("Access-Control-Allow-Methods", "*")
		return
	}

	if r.Method == "POST" {
		userid, valid := checkToken(w, r)
		if valid {
			// Read the raw body
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer r.Body.Close()

			fmt.Printf("Raw body: %s\n", body)

			var t AccountLogic.NewPasswordMessage

			err = json.Unmarshal(body, &t)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			AccountLogic.ChangePasswprd(t, userid)
			w.WriteHeader(http.StatusOK)
		}
	}
}

func refresh(w http.ResponseWriter, r *http.Request) {
	fmt.Println("refresh")
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "*") // You can add more headers here if needed
		w.Header().Set("Access-Control-Allow-Methods", "*")
		return
	}
	if r.Method == "POST" {
		userid, valid := checkToken(w, r)
		if valid {
			Database.DeleteAccount(userid)
		}
	}
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/login", login)
	http.HandleFunc("/isLoggedIn", isLoggedIn)
	http.HandleFunc("/refresh", refresh)
	http.HandleFunc("/userInfo", accountInfo)
	http.HandleFunc("/changeUsername", changeUsername)
	http.HandleFunc("/changePassword", changePassword)
	http.HandleFunc("/deleteAccount", deleteAccount)
}

func Server() {
	fmt.Println("starting")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":9090", nil))
}
