package Server

import (
	"RollsOfDestiny/AccountServer/AccountLogic"
	"RollsOfDestiny/AccountServer/Database"
	"RollsOfDestiny/AccountServer/SignUpLogic"
	"RollsOfDestiny/AccountServer/Types"
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

func getFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "*") // You can add more headers here if needed
		w.Header().Set("Access-Control-Allow-Methods", "*")
		return
	}
	userid, valid := checkToken(w, r)
	if valid {
		friends, err := Database.GetFriendsByUserID(userid)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var friendsString string

		for friend := range friends {
			if friends[friend].UserID != "" {
				friendsString = fmt.Sprintf(`%s, {"username": "%s", "rating": "%s", "profilePicture": "%s"}`, friendsString, friends[friend].Friend.Username, strconv.Itoa(friends[friend].Friend.Rating), friends[friend].Friend.ProfilePicture)
			}
		}
		friendInfo := fmt.Sprintf("{\"friends\": [%s]}", friendsString[2:])
		log.Println(friendInfo)
		fmt.Fprint(w, friendInfo)
		return
	}
	w.WriteHeader(http.StatusForbidden)
}

func newFriend(w http.ResponseWriter, r *http.Request) {
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

			var t Types.FriendInfo

			err = json.Unmarshal(body, &t)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			friend, err := Database.GetAccountByUsername(t.FriendUsername)

			if err != nil {
				log.Println(err)
				return
			}

			Database.InsertNewFriend(userid, friend.UserID)
			w.WriteHeader(http.StatusOK)
		}
	}
}

func deleteFriend(w http.ResponseWriter, r *http.Request) {
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

			var t Types.FriendInfo

			err = json.Unmarshal(body, &t)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			friend, err := Database.GetAccountByUsername(t.FriendUsername)

			if err != nil {
				log.Println(err)
				return
			}

			Database.DeleteFriend(userid, friend.UserID)
			w.WriteHeader(http.StatusOK)
		}
	}
}

func getAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "*") // You can add more headers here if needed
		w.Header().Set("Access-Control-Allow-Methods", "*")
		return
	}
	_, valid := checkToken(w, r)
	if valid {
		myUrl, _ := url.Parse(r.URL.String())
		params, _ := url.ParseQuery(myUrl.RawQuery)
		account := params.Get("username")

		possibleAccounts, err := Database.GetAccountByPartUsername(account)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var accountString string
		for accountId := range possibleAccounts {
			if possibleAccounts[accountId].UserID != "" {
				accountString = fmt.Sprintf(`%s, {"username": "%s", "rating": "%s", "profilePicture": "%s"}`, accountString, possibleAccounts[accountId].Username, strconv.Itoa(possibleAccounts[accountId].Rating), possibleAccounts[accountId].ProfilePicture)
			}
		}
		friendInfo := fmt.Sprintf("{\"friends\": [%s]}", accountString[2:])
		log.Println(friendInfo)
		fmt.Fprint(w, friendInfo)
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
	http.HandleFunc("/getFriends", getFriends)
	http.HandleFunc("/getAccounts", getAccounts)
	http.HandleFunc("/addFriend", newFriend)
}

func Server() {
	fmt.Println("starting")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":9090", nil))
}
