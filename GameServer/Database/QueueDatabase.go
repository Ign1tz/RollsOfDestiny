package Database

import (
	"RollsOfDestiny/GameServer/Types"
	"fmt"
)

func GetOldestEntry() (Types.QueueInfo, error) {
	dbPlayer := Database.QueryRow("Select * from queue where placement = (SELECT MIN(placement) FROM queue)")
	var player Types.QueueInfo
	var placement int
	err := dbPlayer.Scan(&player.UserId, &player.WebsocketConnectionId, &placement)
	if err != nil {
		player = Types.QueueInfo{}
	}
	fmt.Println(player.UserId)
	Database.Exec("Delete from queue where userid = $1", player.UserId)
	return player, err
}

func AddToQueueDatabase(player Types.QueueInfo) error {
	_, err := Database.Exec("INSERT INTO queue (userid, websocketconnectionid) VALUES ($1, $2)", player.UserId, player.WebsocketConnectionId)
	return err
}
