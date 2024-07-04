package Database

import (
	"RollsOfDestiny/AccountServer/Types"
	"log"
)

func UpdateUsername(username string, newUsername string) error {
	_, err := Database.Exec("Update accounts set username = $1 where username = $2 ", newUsername, username)
	return err
}

func UpdateProfilePicture(userid string, profilePicture string) error {
	_, err := Database.Exec("Update accounts set profilepicture = $1 where userid = $2 ", profilePicture, userid)
	return err
}

func UpdateRating(userid string, rating int) error {
	_, err := Database.Exec("Update accounts set rating = rating + $1 where userid = $2 ", rating, userid)
	return err
}

func UpdatePassword(userid string, password string) error {
	_, err := Database.Exec("Update accounts set password = $1 where userid = $2 ", password, userid)
	return err
}

func UpdateCardDeckId(cardInfo Types.AddCard, userid string) error {
	_, err := Database.Exec("Update accountcards set deckids = concat(Select deckids from accountcards where userid = $1 and name = $2, ', ', $3) where userid = $1 and name = $3", userid, cardInfo.Name, cardInfo.Deckid)
	return err
}

func RemoveCardDeckId(userid string, deckid string) error {
	_, err := Database.Exec("Update accountcards set deckids = replace(Select deckids from accountcards where userid = $1 and name = 'Roll Again', $2, '') where userid = $1 and name = 'Roll Again'", userid, deckid)
	if err != nil {
		return err
	}
	_, err = Database.Exec("Update accountcards set deckids = replace(Select deckids from accountcards where userid = $1 and name = 'Double Mana', $2, '') where userid = $1 and name = 'Double Mana'", userid, deckid)
	if err != nil {
		return err
	}
	_, err = Database.Exec("Update accountcards set deckids = replace(Select deckids from accountcards where userid = $1 and name = 'Destroy Column', $2, '') where userid = $1 and name = 'Destroy Column'", userid, deckid)
	if err != nil {
		return err
	}
	_, err = Database.Exec("Update accountcards set deckids = replace(Select deckids from accountcards where userid = $1 and name = 'Flip Clockwise', $2, '') where userid = $1 and name = 'Flip Clockwise'", userid, deckid)
	return err
}

func ChangeActiveDeck(deckid string, userid string) error {
	_, err := Database.Exec("Update accountdecks set active = 'false' where userid = $1 and active = 'true' ", userid)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = Database.Exec("Update accountdecks set active = 'true' where deckid = $1 ", deckid)
	return err
}

func getCardByUseridAndName(userid, name string) (string, error) {
	log.Println(userid, name)
	deckids := Database.QueryRow("Select deckids from accountcards where userid = $1 and name = $2", userid, name)

	var deckidsString string
	if err := deckids.Scan(&deckidsString); err != nil {
		return "", err
	}
	return deckidsString, nil
}

func RemoveCardDeckIdByDeckId(userid string, deckid string) error {

	deckIds, err := getCardByUseridAndName(userid, "Roll Again")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = Database.Exec("Update accountcards set deckids = replace($3, $2, '') where userid = $1 and name = 'Roll Again'", userid, deckid, deckIds)
	if err != nil {
		return err
	}

	deckIds, err = getCardByUseridAndName(userid, "Double Mana")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = Database.Exec("Update accountcards set deckids = replace($3, $2, '') where userid = $1 and name = 'Double Mana'", userid, deckid, deckIds)
	if err != nil {
		return err
	}

	deckIds, err = getCardByUseridAndName(userid, "Destroy Column")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = Database.Exec("Update accountcards set deckids = replace($3, $2, '') where userid = $1 and name = 'Destroy Column'", userid, deckid, deckIds)
	if err != nil {
		return err
	}

	deckIds, err = getCardByUseridAndName(userid, "Flip Clockwise")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = Database.Exec("Update accountcards set deckids = replace($3, $2, '') where userid = $1 and name = 'Flip Clockwise'", userid, deckid, deckIds)
	if err != nil {
		return err
	}

	_, err = Database.Exec("DELETE FROM accountdecks Where deckid = $1", deckid)
	return err
}
