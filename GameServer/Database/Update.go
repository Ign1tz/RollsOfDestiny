package Database

import "RollsOfDestiny/GameServer/Types"

func UpdatePlayer(player Types.Player) error {
	_, err := Database.Exec("Update players set mana = $1 where userid = $2", player.Mana, player.UserID)
	return err
}

func UpdateColumn(column Types.Column) error {
	_, err := Database.Exec("Update columns set first = $1, second = $2, third = $3 where gridid = $4 and placement = $5", column.First, column.Second, column.Third, column.GridId, column.Placement)
	return err
}

func UpdateCard(card Types.Card) error {
	_, err := Database.Exec("Update cards set played = $1, inhand = $2 where cardid = $3 and name = $4 and deckid = $5", card.Played, card.InHand, card.CardID, card.Name, card.DeckID)
	return err
}

func UpdateActivePlayerGames(playfiled Types.Playfield) error {
	_, err := Database.Exec("Update games set activeplayer = $1 where gameid = $2", playfiled.ActivePlayer.UserID, playfiled.GameID)
	return err
}

func UpdateLastRollGames(playfiled Types.Playfield) error {
	_, err := Database.Exec("Update games set lastRoll = $1 where gameid = $2", playfiled.LastRoll, playfiled.GameID)
	return err
}
