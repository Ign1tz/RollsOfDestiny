package GameLogic

import (
	"RollsOfDestiny/GameServer/Database"
	"math/rand"
)

func BotTurn(gameId string) {
	gamefield, err := Database.GetPlayfield(gameId)

	if err != nil {
		panic(err)
	}

	currentRoll := gamefield.ActivePlayer.Die.Throw()

	pickedValidColumn := false

	enemy := gamefield.EnemyPlayer()
	var columnErr error

	for !pickedValidColumn {
		columnNumber := rand.Intn(3)
		if columnNumber == 0 {
			columnErr = gamefield.ActivePlayer.Grid.Left.Add(currentRoll)
			if columnErr == nil {
				enemy.Grid.Left.Remove(currentRoll)
				pickedValidColumn = true
			}
		} else if columnNumber == 1 {
			columnErr = gamefield.ActivePlayer.Grid.Middle.Add(currentRoll)
			if columnErr == nil {
				enemy.Grid.Middle.Remove(currentRoll)
				pickedValidColumn = true
			}
		} else if columnNumber == 2 {
			columnErr = gamefield.ActivePlayer.Grid.Middle.Add(currentRoll)
			if columnErr == nil {
				enemy.Grid.Right.Remove(currentRoll)
				pickedValidColumn = true
			}
		}
	}

}
