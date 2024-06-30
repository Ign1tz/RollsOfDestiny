package Database

func UpdateUsername(username string, newUsername string) error {
	_, err := Database.Exec("Update accounts set username = $1 where username = $2 ", newUsername, username)
	return err
}

func UpdatePassword(userid string, password string) error {
	_, err := Database.Exec("Update accounts set password = $1 where userid = $2 ", password, userid)
	return err
}
