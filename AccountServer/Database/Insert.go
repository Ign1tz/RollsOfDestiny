package Database

import (
	"RollsOfDestiny/AccountServer/Types"
	"fmt"
)

func InsertAccount(account Types.Account) error {
	_, err := Database.Exec("INSERT INTO accounts (userid, username, password, email, profilepicture, rating) VALUES ($1, $2, $3, $4, $5, $6)",
		account.UserID, account.Username, account.Password, account.Email, account.ProfilePicture, account.Rating)
	fmt.Println("aaaaaaaaaaaaaa", err)
	return err
}

func InsertDeck(deck Types.Deck) error {
	_, err := Database.Exec("INSERT INTO accountdecks (userid, deckid, name) Values ($1, $2, $3)", deck.UserID, deck.DeckID, deck.Name)
	return err
}

func InsertCard(card Types.Card) error {
	_, err := Database.Exec("INSERT INTO accountcards (userid, name, effect, deckid, count) Values ($1, $2, $3, $4, $5)",
		card.UserID, card.Name, card.Effect, card.DeckID, card.Count)
	return err
}
