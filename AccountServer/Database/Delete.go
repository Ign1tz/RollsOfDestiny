package Database

func DeleteAccount(userID string) error {
	_, err := Database.Exec("DELETE FROM accounts Where userid = $1", userID)
	return err
}

func DeleteFriend(userid string, friendid string) error {
	_, err := Database.Exec("DELETE FROM accountfriends Where userid = $1 and friendid = $2", userid, friendid)
	return err
}
