package Database

import (
	"RollsOfDestiny/AccountServer/Types"
)

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
	var decks []Types.Deck
	id := 0
	for dbDecks.Next() {
		if err := dbDecks.Scan(&decks[id].UserID, &decks[id].DeckID, &decks[id].Name); err != nil {
			return []Types.Deck{}, err
		}
		id++
	}
	return decks, nil
}

func GetCardsByDeckID(deckID string) ([]Types.Card, error) {
	dbCards, err := Database.Query("Select * from accountcards where deckid like %$1%", deckID)
	if err != nil {
		return []Types.Card{}, err
	}
	var cards []Types.Card
	id := 0
	for dbCards.Next() {
		if err := dbCards.Scan(&cards[id].UserID, &cards[id].Name, &cards[id].Effect, &cards[id].DeckID, &cards[id].Count); err != nil {
			return []Types.Card{}, err
		}
		id++
	}
	return cards, err
}
