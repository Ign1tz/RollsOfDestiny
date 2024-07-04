package Database

import (
	"RollsOfDestiny/AccountServer/Types"
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

func GetCardsByDeckID(deckID string, name string) ([]Types.Card, error) {
	dbCards, err := Database.Query("Select * from accountcards where deckids like '%' || $1 || '%'", deckID)
	if err != nil {
		return []Types.Card{}, err
	}
	var cards = make([]Types.Card, 100)
	id := 0
	for dbCards.Next() {
		if err := dbCards.Scan(&cards[id].UserID, &cards[id].Name, &cards[id].Effect, &cards[id].DeckID, &cards[id].Count, &cards[id].Cost, &cards[id].Image); err != nil {
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
