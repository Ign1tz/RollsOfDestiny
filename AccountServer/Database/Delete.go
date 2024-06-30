package Database

func DeleteAccount(userID string) error {
	_, err := Database.Exec("DELETE FROM accounts Where userid = $1", userID)
	return err
}
