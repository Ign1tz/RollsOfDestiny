package GameLogic

import (
	"RollsOfDestiny/GameServer/Database"
	"RollsOfDestiny/GameServer/Types"
	"errors"
	"strconv"
)

func RollDie(gameId string) string {
	gamefield, err := Database.GetPlayfield(gameId)

	if err != nil {
		panic(err)
	}

	currentRoll := gamefield.ActivePlayer.Die.Throw()
	err = Database.UpdateLastRollGames(gamefield)
	if err != nil {
		panic(err)
	}
	return currentRoll
}

func PickColumn(gameId string, columnNumber string) error {
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

	if columnNumber == "0" {
		columnErr = gamefield.ActivePlayer.Grid.Left.Add(rolls)
		if columnErr == nil {
			enemy.Grid.Left.Remove(rolls)
			playerColumn = gamefield.ActivePlayer.Grid.Left
			enemyColumn = enemy.Grid.Left
		}
	} else if columnNumber == "1" {
		columnErr = gamefield.ActivePlayer.Grid.Middle.Add(rolls)
		if columnErr == nil {
			enemy.Grid.Middle.Remove(rolls)
			playerColumn = gamefield.ActivePlayer.Grid.Middle
			enemyColumn = enemy.Grid.Middle
		}
	} else if columnNumber == "2" {
		columnErr = gamefield.ActivePlayer.Grid.Right.Add(rolls)
		if columnErr == nil {
			enemy.Grid.Right.Remove(rolls)
			playerColumn = gamefield.ActivePlayer.Grid.Right
			enemyColumn = enemy.Grid.Right
		}
	} else {
		columnErr = errors.New("wrong column number")
	}
	gamefield.ActivePlayer = gamefield.EnemyPlayer()
	if columnErr == nil {
		err = Database.UpdateColumn(playerColumn)
		if err != nil {
			return err
		}
		err = Database.UpdateColumn(enemyColumn)
		if err != nil {
			return err
		}
		err = Database.UpdateActivePlayerGames(gamefield)
	}
	return columnErr
}
