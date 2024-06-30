package Database

func DeleteQueue() {
	Database.Exec("TRUNCATE TABLE queue")
}
func DeleteGames() {
	Database.Exec("TRUNCATE TABLE grids")
}

func DeleteFromQueue(websocketId string) {
	Database.Exec("Delete from queue where websocketconnectionid = $1", websocketId)
}

func DeleteGame(gridid string) error {
	_, err := Database.Exec("DELETE FROM grids WHERE gridid = $1", gridid)
	return err
}
