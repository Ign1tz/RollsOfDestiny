package Database

func UpdateUsername(username string, newUsername string) error {
	_, err := Database.Exec("Update accounts set username = $1 where username = $2 ", newUsername, username)
	return err
}

func UpdatePassword(userid string, password string) error {
	_, err := Database.Exec("Update accounts set password = $1 where userid = $2 ", password, userid)
	return err
}

func UpdateCardDeckId(userid string, deckid string) error {
	_, err := Database.Exec("Update accountcards set deckids = concat(Select deckids from accountcards where userid = $1, ', ', $2) where userid = $1", userid, deckid)
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

func ChangeActiveDeck(deckid string) error {
	_, err := Database.Exec("Update accountdecks set active = '0' where deckid = $1 and active = '1' ", deckid)
	if err != nil {
		return err
	}
	_, err = Database.Exec("Update accountdecks set active = '1' where deckid = $1 ", deckid)
	return err
}

func RemoveCardDeckIdByDeckId(userid string, deckid string) error {
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
	if err != nil {
		return err
	}

	_, err = Database.Exec("DELETE FROM accountdecks Where deckid = $1", deckid)
	return err
}
