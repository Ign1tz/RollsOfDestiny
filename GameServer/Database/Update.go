package Database

import (
	"RollsOfDestiny/GameServer/Types"
	"log"
)

func UpdatePlayerMana(player Types.Player) error {
	_, err := Database.Exec("Update players set mana = $1 where userid = $2", player.Mana, player.UserID)
	return err
}

func UpdatePlayerWebsocketID(userid string, websocketid string) error {
	_, err := Database.Exec("Update players set websocketconnectionid = $1 where userid = $2", websocketid, userid)
	return err
}

func UpdateColumn(column Types.Column) error {
	log.Println(column.First, column.Second, column.Third, column.GridId, column.Placement)
	_, err := Database.Exec("Update columns set first = $1, second = $2, third = $3 where gridid = $4 and placement = $5", column.First, column.Second, column.Third, column.GridId, column.Placement)
	return err
}

func UpdateCard(card Types.Card) error {
	_, err := Database.Exec("Update cards set played = $1, inhand = $2 where cardid = $3 and name = $4 and deckid = $5", card.Played, card.InHand, card.CardID, card.Name, card.DeckID)
	return err
}

func UpdateActivePlayerGames(playfield Types.Playfield) error {
	_, err := Database.Exec("Update games set activeplayer = $1 where gameid = $2", playfield.ActivePlayer.UserID, playfield.GameID)
	return err
}

func UpdateLastRollGames(playfield Types.Playfield) error {
	_, err := Database.Exec("Update games set lastRoll = $1 where gameid = $2", playfield.LastRoll, playfield.GameID)
	return err
}

func UpdateQueueEntry(userId string, websocketId string) error {
	_, err := Database.Exec("Update queue set websocketconnectionid = $1 where userid = $2", websocketId, userId)
	return err
}

func UpdateFriendQueueEntry(userId string, websocketId string) error {
	_, err := Database.Exec("Update friendqueue set websocketconnectionid = $1 where userid = $2", websocketId, userId)
	return err
}

func UpdatePosition(position Types.Position) error {
	_, err := Database.Exec("Update position set currentstep = $1, hostinfo = $2, guestinfo = $3 where gameid = $4", position.CurrentStep, position.HostInfo, position.GuestInfo, position.Gameid)
	return err
}
