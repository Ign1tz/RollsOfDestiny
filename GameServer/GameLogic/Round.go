package GameLogic

import (
	"RollsOfDestiny/GameServer/Database"
	"RollsOfDestiny/GameServer/Types"
)

func RollDie(gameId string) int {
	gamefield, err := Database.GetPlayfield(gameId)

	if err != nil {
		panic(err)
	}
	var activePlayer Types.Player
	if gamefield.Host.Username == gamefield.ActivePlayer {
		activePlayer = gamefield.Host
	} else {
		activePlayer = gamefield.Guest
	}
	currentRoll := activePlayer.Die.Throw()

	return currentRoll
}

func PickColumn(gameId string, columnNumber int) error {

	return nil
}
