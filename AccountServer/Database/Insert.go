package Database

import (
	"RollsOfDestiny/AccountServer/Types"
)

func InsertAccount(account Types.Account) error {
	_, err := Database.Exec("INSERT INTO accounts (userid, username, password, email, profilepicture, rating) VALUES ($1, $2, $3, $4, $5, $6)",
		account.UserID, account.Username, account.Password, account.Email, account.ProfilePicture, account.Rating)
	return err
}

func InsertDeck(deck Types.Deck) error {
	_, err := Database.Exec("INSERT INTO accountdecks (userid, deckid, name, active) Values ($1, $2, $3, $4)", deck.UserID, deck.DeckID, deck.Name, deck.Active)
	return err
}

func InsertCard(card Types.Card) error {
	_, err := Database.Exec("INSERT INTO accountcards (userid, name, effect, deckids, count, cost, image, threshold) Values ($1, $2, $3, $4, $5, $6, $7, $8)",
		card.UserID, card.Name, card.Effect, card.DeckID, card.Count, card.Cost, card.Image, card.Threshold)
	return err
}

func InsertNewFriend(userid string, friendid string) error {
	_, err := Database.Exec("INSERT INTO accountfriends (userid, friendid) Values ($1, $2)",
		userid, friendid)
	return err
}
