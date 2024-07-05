package Database

import (
	"RollsOfDestiny/AccountServer/Types"
	"os"
	"strconv"
)

func GetAccountByUserID(userID string) (Types.Account, error) {
	dbAccount := Database.QueryRow("Select * from accounts where userid = $1", userID)
	var account Types.Account
	if err := dbAccount.Scan(&account.UserID, &account.Username, &account.Password, &account.Email, &account.ProfilePicture, &account.Rating); err != nil {
		return Types.Account{}, err
	}
	return account, nil
}

func GetAccountByUsername(username string) (Types.Account, error) {
	dbAccount := Database.QueryRow("Select * from accounts where username = $1", username)
	var account Types.Account
	if err := dbAccount.Scan(&account.UserID, &account.Username, &account.Password, &account.Email, &account.ProfilePicture, &account.Rating); err != nil {
		return Types.Account{}, err
	}
	return account, nil
}

func GetAccountByEmail(email string) (Types.Account, error) {
	dbAccount := Database.QueryRow("Select * from accounts where email = $1", email)
	var account Types.Account
	if err := dbAccount.Scan(&account.UserID, &account.Username, &account.Password, &account.Email, &account.ProfilePicture, &account.Rating); err != nil {
		return Types.Account{}, err
	}
	return account, nil
}

func GetDecksByUserID(userID string) ([]Types.Deck, error) {
	dbDecks, err := Database.Query("Select * from accountdecks where userid = $1", userID)
	if err != nil {
		return []Types.Deck{}, err
	}
	var decks = make([]Types.Deck, 100)
	id := 0
	for dbDecks.Next() {
		if err := dbDecks.Scan(&decks[id].UserID, &decks[id].DeckID, &decks[id].Name, &decks[id].Active); err != nil {
			return []Types.Deck{}, err
		}
		id++
	}
	return decks, nil
}

func GetDeckByDeckId(deckId string) (Types.Deck, error) {
	dbDecks := Database.QueryRow("Select * from accountdecks where deckid = $1", deckId)
	var deck = Types.Deck{}
	if err := dbDecks.Scan(&deck.UserID, &deck.DeckID, &deck.Name, &deck.Active); err != nil {
		return Types.Deck{}, err
	}
	return deck, nil
}

func GetCardsByDeckID(deckID string) ([]Types.Card, error) {
	dbCards, err := Database.Query("Select * from accountcards where deckids like '%' || $1 || '%'", deckID)
	if err != nil {
		return []Types.Card{}, err
	}
	numberOfCards, _ := strconv.Atoi(os.Getenv("NUMBER_OF_DIFFERENT_CARDS"))
	var cards = make([]Types.Card, numberOfCards)
	id := 0
	for dbCards.Next() {
		if err := dbCards.Scan(&cards[id].UserID, &cards[id].Name, &cards[id].Effect, &cards[id].DeckID, &cards[id].Count, &cards[id].Cost, &cards[id].Image, &cards[id].Threshold); err != nil {
			return []Types.Card{}, err
		}
		id++
	}
	return cards, err
}

func GetCardsByUserId(userid string) ([]Types.Card, error) {
	dbCards, err := Database.Query("Select * from accountcards where userid = $1 and not cost = 0", userid)
	if err != nil {
		return []Types.Card{}, err
	}
	numberOfCards, _ := strconv.Atoi(os.Getenv("NUMBER_OF_DIFFERENT_CARDS"))
	var cards = make([]Types.Card, numberOfCards)
	id := 0
	for dbCards.Next() {
		if err := dbCards.Scan(&cards[id].UserID, &cards[id].Name, &cards[id].Effect, &cards[id].DeckID, &cards[id].Count, &cards[id].Cost, &cards[id].Image, &cards[id].Threshold); err != nil {
			return []Types.Card{}, err
		}
		id++
	}
	return cards, err
}

func GetAllCardsByUserId(userid string) ([]Types.Card, error) {
	dbCards, err := Database.Query("Select * from accountcards where userid = $1", userid)
	if err != nil {
		return []Types.Card{}, err
	}
	numberOfCards, _ := strconv.Atoi(os.Getenv("NUMBER_OF_DIFFERENT_CARDS"))
	var cards = make([]Types.Card, numberOfCards)
	id := 0
	for dbCards.Next() {
		if err := dbCards.Scan(&cards[id].UserID, &cards[id].Name, &cards[id].Effect, &cards[id].DeckID, &cards[id].Count, &cards[id].Cost, &cards[id].Image, &cards[id].Threshold); err != nil {
			return []Types.Card{}, err
		}
		id++
	}
	return cards, err
}

func GetFriendsByUserID(userID string) ([]Types.Friend, error) {
	dbDecks, err := Database.Query("Select * from accountfriends where userid = $1", userID)
	if err != nil {
		return []Types.Friend{}, err
	}
	var friends = make([]Types.Friend, 100)
	id := 0
	for dbDecks.Next() {
		var userid string
		if err := dbDecks.Scan(&friends[id].UserID, &userid); err != nil {
			return []Types.Friend{}, err
		}
		account, err := GetAccountByUserID(userid)
		if err != nil {

		} else {
			friends[id].Friend = account
			id++
		}
	}
	return friends, nil
}

func GetAccountByPartUsername(usernamePart string, userid string) ([]Types.Account, error) {
	dbAccount, err := Database.Query("Select * from accounts where username ilike '%' || $1 || '%' and userid is distinct from $2 limit 10", usernamePart, userid)
	if err != nil {
		return []Types.Account{}, err
	}
	var accounts = make([]Types.Account, 10)
	id := 0
	for dbAccount.Next() {
		var account Types.Account
		if err := dbAccount.Scan(&account.UserID, &account.Username, &account.Password, &account.Email, &account.ProfilePicture, &account.Rating); err != nil {
			return []Types.Account{}, err
		}
		accounts[id] = account
		id++
	}

	return accounts, nil
}

func GetTopTenPlayers() ([]Types.Account, error) {
	dbAccount, err := Database.Query("Select * from accounts order by rating desc limit 10")
	if err != nil {
		return []Types.Account{}, err
	}
	var accounts = make([]Types.Account, 10)
	id := 0
	for dbAccount.Next() {
		var account Types.Account
		if err := dbAccount.Scan(&account.UserID, &account.Username, &account.Password, &account.Email, &account.ProfilePicture, &account.Rating); err != nil {
			return []Types.Account{}, err
		}
		accounts[id] = account
		id++
	}

	return accounts, nil
}
