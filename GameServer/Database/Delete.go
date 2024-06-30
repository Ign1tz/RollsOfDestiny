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
