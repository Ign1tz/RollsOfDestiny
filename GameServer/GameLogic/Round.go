package GameLogic

import (
	"RollsOfDestiny/GameServer/Database"
	"RollsOfDestiny/GameServer/Types"
	"errors"
	"strconv"
)

func RollDie(gameId string) int {
	gamefield, err := Database.GetPlayfield(gameId)

	if err != nil {
		panic(err)
	}

	currentRoll := gamefield.ActivePlayer.Die.Throw()
	//todo: add database update to save roll
	err = Database.UpdateLastRollGames(gamefield)
	if err != nil {
		panic(err)
	}
	return currentRoll
}

func PickColumn(gameId string, columnNumber int) error {
	gamefield, err := Database.GetPlayfield(gameId)

	if err != nil {
		panic(err)
	}

	rolls, err := strconv.Atoi(gamefield.LastRoll)

	if err != nil {
		return err
	}
	var columnErr error
	enemy := gamefield.EnemyPlayer()
	var playerColumn Types.Column
	var enemyColumn Types.Column

	if columnNumber == 0 {
		columnErr = gamefield.ActivePlayer.Grid.Left.Add(rolls)
		if columnErr == nil {
			enemy.Grid.Left.Remove(rolls)
			playerColumn = gamefield.ActivePlayer.Grid.Left
			enemyColumn = gamefield.EnemyPlayer().Grid.Left
		}
	} else if columnNumber == 1 {
		columnErr = gamefield.ActivePlayer.Grid.Middle.Add(rolls)
		if columnErr == nil {
			enemy.Grid.Middle.Remove(rolls)
			playerColumn = gamefield.ActivePlayer.Grid.Middle
			enemyColumn = gamefield.EnemyPlayer().Grid.Middle
		}
	} else if columnNumber == 2 {
		columnErr = gamefield.ActivePlayer.Grid.Right.Add(rolls)
		if columnErr == nil {
			enemy.Grid.Right.Remove(rolls)
			playerColumn = gamefield.ActivePlayer.Grid.Right
			enemyColumn = gamefield.EnemyPlayer().Grid.Right
		}
	} else {
		columnErr = errors.New("wrong column number")
	}

	//todo: update database
	if columnErr == nil {
		Database.UpdateColumn(playerColumn)
		Database.UpdateColumn(enemyColumn)
	}
	return columnErr
}
