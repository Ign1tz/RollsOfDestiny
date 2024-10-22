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

func UpdateCardCount(userid string, name string) error {
	_, err := Database.Exec("Update accountcards set count = 1 where userid = $1 and name = $2 ", userid, name)
	return err
}

func UpdateCardDeckId(cardInfo Types.AddCard, userid string) error {
	_, err := Database.Exec(`
    UPDATE accountcards 
    SET deckids = 
        CASE 
            WHEN deckids IS NULL THEN $3 
            ELSE deckids || ', ' || $3 
        END
    WHERE userid = $1 
      AND name = $2`,
		userid, cardInfo.Name, cardInfo.Deckid)
	return err
}

func RemoveCardDeckId(userid string, cardInfo Types.AddCard) error {
	_, err := Database.Exec(`UPDATE accountcards
	SET deckids = 
	REPLACE(
		deckids,
		$2,
		''
	)
	WHERE userid = $1 AND name = $3`, userid, ", "+cardInfo.Deckid, cardInfo.Name)
	if err != nil {
		return err
	}
	return err
}

func ChangeActiveDeck(deckid string, userid string) error {
	_, err := Database.Exec("Update accountdecks set active = case when deckid = $1 then not active else 'false' end where userid = $2", deckid, userid)
	if err != nil {
		log.Println(err)
		return err
	}
	/*_, err = Database.Exec("Update accountdecks set active = 'true' where deckid = $1 ", deckid)*/
	return err
}

func getCardByUseridAndName(userid, name string) (string, error) {
	deckids := Database.QueryRow("Select deckids from accountcards where userid = $1 and name = $2", userid, name)

	var deckidsString string
	if err := deckids.Scan(&deckidsString); err != nil {
		return "", err
	}
	return deckidsString, nil
}

func RemoveCardDeckIdByDeckId(userid string, deckid string) error {

	_, err := Database.Exec("Update accountcards set deckids = replace(deckids, $2, '') where userid = $1", userid, ", "+deckid)
	if err != nil {
		return err
	}
	_, err = Database.Exec("DELETE FROM accountdecks Where deckid = $1", deckid)
	return err
}
