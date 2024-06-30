package Database

func DeleteQueue() {
	Database.Exec("DELETE From queue")
}
func DeleteGames() {
	Database.Exec("DELETE From grids")
}

func DeleteFromQueueWebsocket(websocketId string) {
	Database.Exec("Delete from queue where websocketconnectionid = $1", websocketId)
}

func DeleteGame(gridid string) error {
	_, err := Database.Exec("DELETE FROM grids WHERE gridid = $1", gridid)
	return err
}
