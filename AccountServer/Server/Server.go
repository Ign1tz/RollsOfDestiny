package Server

import (
	"RollsOfDestiny/AccountServer/AccountLogic"
	"RollsOfDestiny/AccountServer/CardLogic"
	"RollsOfDestiny/AccountServer/Database"
	"RollsOfDestiny/AccountServer/DeckLogic"
	"RollsOfDestiny/AccountServer/SignUpLogic"
	"RollsOfDestiny/AccountServer/Types"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homag Page")
}

func signUp(w http.ResponseWriter, r *http.Request) {
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
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Println(err)
				return
			}
		}(r.Body)

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
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Println(err)
				return
			}
		}(r.Body)

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
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					log.Println(err)
					return
				}
			}(r.Body)

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
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					log.Println(err)
					return
				}
			}(r.Body)

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

func changeProfilePicture(w http.ResponseWriter, r *http.Request) {
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
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					log.Println(err)
					return
				}
			}(r.Body)

			var t AccountLogic.NewProfilePicture

			err = json.Unmarshal(body, &t)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				log.Println(err)
				return
			}
			err = Database.UpdateProfilePicture(userid, t.ProfilePicture)
			log.Println(err)
			if err != nil {
				log.Println(err)
				return
			}
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
		log.Println("deleteAccount", valid)
		if valid {
			err := Database.DeleteAccount(userid)
			if err != nil {
				return
			}
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
		var array string
		if len(friendsString) > 2 {
			array = friendsString[2:]
		} else {
			array = ""
		}
		friendInfo := fmt.Sprintf("{\"friends\": [%s]}", array)
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
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					log.Println(err)
					return
				}
			}(r.Body)

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
				log.Println(err)
				return
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					log.Println(err)
					return
				}
			}(r.Body)

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

			err = Database.DeleteFriend(userid, friend.UserID)
			if err != nil {
				log.Println(err)
				return
			}
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
	userid, valid := checkToken(w, r)
	if valid {
		myUrl, _ := url.Parse(r.URL.String())
		params, _ := url.ParseQuery(myUrl.RawQuery)
		account := params.Get("username")

		possibleAccounts, err := Database.GetAccountByPartUsername(account, userid)

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
		var array string
		if len(accountString) > 2 {
			array = accountString[2:]
		} else {
			array = ""
		}
		friendInfo := fmt.Sprintf("{\"friends\": [%s]}", array)
		fmt.Fprint(w, friendInfo)
	}
}

func getDecks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*") // You can add more headers here if needed
	w.Header().Set("Access-Control-Allow-Methods", "*")
	if r.Method == "OPTIONS" {

		return
	}
	userid, valid := checkToken(w, r)
	if valid {

		decks, err := Database.GetDecksByUserID(userid)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var deckString string
		var cardString string
		//todo: check for empty
		for deckIndex := range decks {
			if decks[deckIndex].UserID != "" {
				cardString = DeckLogic.GetCardsOfDeckAsJsonString(decks[deckIndex].DeckID)
				deckString = fmt.Sprintf(`%s, {"name": "%s", "deckid": "%s", "active": %s, "cards": [%s]}`, deckString, decks[deckIndex].Name, decks[deckIndex].DeckID, strconv.FormatBool(decks[deckIndex].Active), cardString)
			}
		}
		var array string
		if len(deckString) > 2 {
			array = deckString[2:]
		} else {
			array = ""
		}
		friendInfo := fmt.Sprintf("{\"decks\": [%s]}", array)
		fmt.Fprint(w, friendInfo)
	}
}

func getDeck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*") // You can add more headers here if needed
	w.Header().Set("Access-Control-Allow-Methods", "*")
	if r.Method == "OPTIONS" {

		return
	}
	_, valid := checkToken(w, r)
	if valid {

		myUrl, _ := url.Parse(r.URL.String())
		params, _ := url.ParseQuery(myUrl.RawQuery)
		deckid := params.Get("deckid")

		log.Println(deckid)

		deck, err := Database.GetDeckByDeckId(deckid)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var deckString string
		var cardString string
		if deck.UserID != "" {
			cardString = DeckLogic.GetCardsOfDeckAsJsonString(deck.DeckID)
			deckString = fmt.Sprintf(`%s{"name": "%s", "deckid": "%s", "active": %s, "cards": [%s]}`, deckString, deck.Name, deck.DeckID, strconv.FormatBool(deck.Active), cardString)
		}
		deckInfo := fmt.Sprintf("{\"deck\": %s}", deckString)
		log.Println(deckInfo)
		fmt.Fprint(w, deckInfo)
	}
}

func createDeck(w http.ResponseWriter, r *http.Request) {
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
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					log.Println(err)
					return
				}
			}(r.Body)

			fmt.Printf("Raw body: %s\n", body)

			var t Types.NewDeckMessage

			err = json.Unmarshal(body, &t)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			log.Println(t.Name)
			replacedName := strings.Replace(t.Name, " ", "", 0)
			log.Println(replacedName)
			if replacedName != "" {
				DeckLogic.CreateNewDeck(t.Name, userid)
			}
			w.WriteHeader(http.StatusOK)
		}
	}
}

func addCardToDeck(w http.ResponseWriter, r *http.Request) {
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
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					log.Println(err)
					return
				}
			}(r.Body)

			fmt.Printf("Raw body: %s\n", body)

			var t Types.AddCard

			err = json.Unmarshal(body, &t)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			DeckLogic.AddCardToDeck(t, userid)
			w.WriteHeader(http.StatusOK)
		}
	}
}

func removeCardFromDeck(w http.ResponseWriter, r *http.Request) {
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
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					log.Println(err)
					return
				}
			}(r.Body)

			fmt.Printf("Raw body: %s\n", body)

			var t Types.AddCard

			err = json.Unmarshal(body, &t)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			DeckLogic.RemoveCardFromDeck(t, userid)
			log.Println("remove")
			w.WriteHeader(http.StatusOK)
		}
	}
}

func setDeckActive(w http.ResponseWriter, r *http.Request) {
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
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					log.Println(err)
					return
				}
			}(r.Body)

			fmt.Printf("Raw body: %s\n", body)

			var t Types.AddCard

			err = json.Unmarshal(body, &t)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			DeckLogic.ChangeActiveDeck(t, userid)
			w.WriteHeader(http.StatusOK)
		}
	}
}

func removeDeck(w http.ResponseWriter, r *http.Request) {
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
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					log.Println(err)
					return
				}
			}(r.Body)

			fmt.Printf("Raw body: %s\n", body)

			var t Types.AddCard

			err = json.Unmarshal(body, &t)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			DeckLogic.RemoveDeck(t.Deckid, userid)
			w.WriteHeader(http.StatusOK)
		}
	}
}

func getTopTen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "*") // You can add more headers here if needed
		w.Header().Set("Access-Control-Allow-Methods", "*")
		return
	}
	_, valid := checkToken(w, r)
	if valid {

		topTenPlayers, err := Database.GetTopTenPlayers()

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var deckString string
		for deckIndex := range topTenPlayers {
			if topTenPlayers[deckIndex].UserID != "" {
				deckString = fmt.Sprintf(`%s, {"username": "%s", "rating": %s, "profilePicture": "%s"}`, deckString, topTenPlayers[deckIndex].Username, strconv.Itoa(topTenPlayers[deckIndex].Rating), topTenPlayers[deckIndex].ProfilePicture)
			}
		}
		var array string
		if len(deckString) > 2 {
			array = deckString[2:]
		} else {
			array = ""
		}
		friendInfo := fmt.Sprintf("{\"topTenPlayers\": [%s]}", array)
		fmt.Fprint(w, friendInfo)
	}
}

func getYourCards(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "*") // You can add more headers here if needed
		w.Header().Set("Access-Control-Allow-Methods", "*")
		return
	}
	userid, valid := checkToken(w, r)
	if valid {

		cards, err := Database.GetCardsByUserId(userid)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var cardString string
		for cardIndex := range cards {
			if cards[cardIndex].UserID != "" {
				cardString = fmt.Sprintf("%s, \"%s\"", cardString, cards[cardIndex].Name)
			}
		}

		var array string
		if len(cardString) > 2 {
			array = cardString[2:]
		} else {
			array = ""
		}

		friendInfo := fmt.Sprintf("{\"cards\": [%s]}", array)
		fmt.Fprint(w, friendInfo)
	}
}

func getNewCards(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*") // You can add more headers here if needed
	w.Header().Set("Access-Control-Allow-Methods", "*")
	if r.Method == "OPTIONS" {
		return
	}
	userid, valid := checkToken(w, r)
	if valid {

		newCards, oldCards, err := CardLogic.HandleNewCard(userid)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var oldCardString string
		for cardIndex := range oldCards {
			if oldCards[cardIndex].UserID != "" {
				oldCardString = fmt.Sprintf("%s, \"%s\"", oldCardString, oldCards[cardIndex].Name)
			}
		}

		var newCardString string
		for cardIndex := range newCards {
			if newCards[cardIndex].UserID != "" {
				newCardString = fmt.Sprintf("%s, \"%s\"", newCardString, newCards[cardIndex].Name)
			}
		}

		var oldArray string
		if len(oldCardString) > 2 {
			oldArray = oldCardString[2:]
		} else {
			oldArray = ""
		}

		var newArray string
		if len(newCardString) > 2 {
			newArray = newCardString[2:]
		} else {
			newArray = ""
		}

		message := fmt.Sprintf("{\"oldCards\": [%s], \"newCards\": [%s]}", oldArray, newArray)
		log.Println("newCards", message)
		fmt.Fprint(w, message)
	}
}

func aknowledgeCard(w http.ResponseWriter, r *http.Request) {
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
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					log.Println(err)
					return
				}
			}(r.Body)

			fmt.Printf("Raw body: %s\n", body)

			var t Types.AknowledgeCard

			err = json.Unmarshal(body, &t)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			Database.UpdateCardCount(userid, t.Name)
			w.WriteHeader(http.StatusOK)
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
	http.HandleFunc("/changeProfilePicture", changeProfilePicture)
	http.HandleFunc("/changeUsername", changeUsername)
	http.HandleFunc("/changePassword", changePassword)
	http.HandleFunc("/deleteAccount", deleteAccount)
	http.HandleFunc("/getFriends", getFriends)
	http.HandleFunc("/getAccounts", getAccounts)
	http.HandleFunc("/getTopTen", getTopTen)
	http.HandleFunc("/addFriend", newFriend)
	http.HandleFunc("/removeFriend", deleteFriend)
	http.HandleFunc("/getDeck", getDeck)
	http.HandleFunc("/getDecks", getDecks)
	http.HandleFunc("/createDeck", createDeck)
	http.HandleFunc("/addCardToDeck", addCardToDeck)
	http.HandleFunc("/removeCardFromDeck", removeCardFromDeck)
	http.HandleFunc("/setActiveDeck", setDeckActive)
	http.HandleFunc("/removeDeck", removeDeck)
	http.HandleFunc("/getYourCards", getYourCards)
	http.HandleFunc("/getNewCards", getNewCards)
	http.HandleFunc("/aknowledgeNewCard", aknowledgeCard)
}

func Server() {
	setupRoutes()
	log.Fatal(http.ListenAndServe(":9090", nil))
}
