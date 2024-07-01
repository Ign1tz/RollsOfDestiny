package Database

import (
	"RollsOfDestiny/GameServer/Types"
)

func GetOldestEntry() (Types.QueueInfo, error) {
	dbPlayer := Database.QueryRow("Select * from queue where placement = (SELECT MIN(placement) FROM queue)")
	var player Types.QueueInfo
	var placement int
	err := dbPlayer.Scan(&player.UserId, &player.WebsocketConnectionId, &placement, &player.Username)
	if err != nil {
		player = Types.QueueInfo{}
	}
	return player, err
}

func AddToQueueDatabase(player Types.QueueInfo) error {
	_, err := Database.Exec("INSERT INTO queue (userid, websocketconnectionid, username) VALUES ($1, $2, $3)", player.UserId, player.WebsocketConnectionId, player.Username)
	return err
}

func DeleteFromQueue(player Types.QueueInfo) error {
	_, err := Database.Exec("Delete from queue where userid = $1", player.UserId)
	return err
}
