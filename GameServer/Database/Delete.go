package Database

func DeleteGame(gridid string) error {
	_, err := Database.Exec("DELETE FROM grids WHERE gridid = $1", gridid)
	return err
}
