package GameLogic

import (
	"RollsOfDestiny/GameServer/Types"
	"github.com/google/uuid"
)

func CreateNewGame(host Types.Player, guest Types.Player) {
	var game Types.Game

	game.GameID = uuid.NewString()
	game.HostId = host.UserID
	game.GuestId = guest.UserID
	game.ActivePlayer = host.UserID
	game.HostGrid = host.Grid.GridId
	game.GuestGrid = host.Grid.GridId

}
