package GameLogic

import (
	"RollsOfDestiny/GameServer/Database"
	"RollsOfDestiny/GameServer/Types"
	"fmt"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"strconv"
)

func BotTurn(gameInfo Types.Resp) bool {

	log.Println(gameInfo.Gameid, gameInfo.ColumnKey)
	ended, err := PickColumn(gameInfo.Gameid, gameInfo.ColumnKey)

	if err != nil {
		panic(err)
	}
	gamefield, err := Database.GetPlayfield(gameInfo.Gameid)
	if ended {
		gamefield.ActivePlayer = gamefield.EnemyPlayer()
		Database.UpdateActivePlayerGames(gamefield)
		return true
	}

	if err != nil {
		panic(err)
	}

	currentRoll := gamefield.ActivePlayer.Die.Throw()

	pickedValidColumn := false

	enemy := gamefield.EnemyPlayer()
	var columnErr error
	roll, _ := strconv.Atoi(currentRoll)
	for !pickedValidColumn {
		columnNumber := rand.Intn(3)
		if columnNumber == 0 {
			columnErr = gamefield.ActivePlayer.Grid.Left.Add(roll)
			if columnErr == nil {
				enemy.Grid.Left.Remove(roll)
				Database.UpdateColumn(gamefield.ActivePlayer.Grid.Left)
				Database.UpdateColumn(enemy.Grid.Left)
				pickedValidColumn = true
			}
		} else if columnNumber == 1 {
			columnErr = gamefield.ActivePlayer.Grid.Middle.Add(roll)
			if columnErr == nil {
				enemy.Grid.Middle.Remove(roll)
				Database.UpdateColumn(gamefield.ActivePlayer.Grid.Middle)
				Database.UpdateColumn(enemy.Grid.Middle)
				pickedValidColumn = true
			}
		} else if columnNumber == 2 {
			columnErr = gamefield.ActivePlayer.Grid.Right.Add(roll)
			if columnErr == nil {
				enemy.Grid.Right.Remove(roll)
				Database.UpdateColumn(gamefield.ActivePlayer.Grid.Right)
				Database.UpdateColumn(enemy.Grid.Right)
				pickedValidColumn = true
			}
		}
	}
	player := gamefield.ActivePlayer
	gamefield.ActivePlayer = enemy
	gamefield.LastRoll = gamefield.ActivePlayer.Die.Throw()
	Database.UpdateLastRollGames(gamefield)
	Database.UpdateActivePlayerGames(gamefield)
	return player.Grid.IsFull()
}

func BotStartGame(queueEntry Types.BotResp, c2 *chan map[string]string) {

	_, err := Database.GetPlayer(queueEntry.Userid)

	if err == nil {
		log.Println("relog")
		playfield, err := Database.GetPlayfieldByUserid(queueEntry.Userid)
		if err != nil {
			log.Println("noPlayfield", err)
			return
		}
		err = Database.UpdatePlayerWebsocketID(queueEntry.Userid, queueEntry.WebsocketConnectionId)
		if err != nil {
			log.Println("websocketUpdate", err)
			return
		}
		var msg = make(map[string]string)
		msg["id"] = queueEntry.WebsocketConnectionId
		newMessage := `{"gameid": "` + playfield.GameID + `", "YourInfo":` + playfield.Host.ToJson(true) + `, "EnemyInfo": ` + playfield.Guest.ToJson(false) + `, "ActivePlayer": {"active": true, "roll": "` + playfield.LastRoll + `"}}`
		infoMessage := `{"info": "gameInfo", "message": {"gameInfo": ` + newMessage + `}, "gameId": "` + playfield.GameID + `"}`
		msg["message"] = infoMessage

		*c2 <- msg
		return
	}

	gridId1 := uuid.New().String()
	hostGrid := Types.Grid{
		Left: Types.Column{
			First:     0,
			Second:    0,
			Third:     0,
			GridId:    gridId1,
			Placement: 0,
		},
		Middle: Types.Column{
			First:     0,
			Second:    0,
			Third:     0,
			GridId:    gridId1,
			Placement: 1,
		},
		Right: Types.Column{
			First:     0,
			Second:    0,
			Third:     0,
			GridId:    gridId1,
			Placement: 2,
		},
		GridId: gridId1,
	}
	gridId2 := uuid.New().String()

	guestGrid := Types.Grid{
		Left: Types.Column{
			First:     0,
			Second:    0,
			Third:     0,
			GridId:    gridId2,
			Placement: 0,
		},
		Middle: Types.Column{
			First:     0,
			Second:    0,
			Third:     0,
			GridId:    gridId2,
			Placement: 1,
		},
		Right: Types.Column{
			First:     0,
			Second:    0,
			Third:     0,
			GridId:    gridId2,
			Placement: 2,
		},
		GridId: gridId2,
	}

	host := Types.Player{
		Username:              "Host",
		UserID:                queueEntry.Userid,
		Mana:                  0,
		Deck:                  Types.Deck{},
		Die:                   Types.Die{PossibleThrows: []int{1, 2, 3, 4, 5, 6}},
		WebsocketConnectionID: queueEntry.WebsocketConnectionId,
		Grid:                  hostGrid,
	}
	guest := Types.Player{
		Username:              "Bot",
		UserID:                uuid.New().String(),
		Mana:                  0,
		Deck:                  Types.Deck{},
		Die:                   Types.Die{PossibleThrows: []int{1, 2, 3, 4, 5, 6}},
		WebsocketConnectionID: uuid.New().String(),
		Grid:                  guestGrid,
	}

	playfield := Types.Playfield{
		Host:         host,
		Guest:        guest,
		HostGrid:     hostGrid,
		GuestGrid:    guestGrid,
		GameID:       "bot: " + uuid.New().String(),
		ActivePlayer: host,
		LastRoll:     Types.Die{PossibleThrows: []int{1, 2, 3, 4, 5, 6}}.Throw(),
	}

	err = Database.InsertWholeGame(playfield)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var msg = make(map[string]string)
	msg["id"] = playfield.Host.WebsocketConnectionID
	newMessage := `{"gameid": "` + playfield.GameID + `", "YourInfo":` + playfield.Host.ToJson(true) + `, "EnemyInfo": ` + playfield.Guest.ToJson(false) + `, "ActivePlayer": {"active": true, "roll": "` + playfield.LastRoll + `"}}`
	infoMessage := `{"info": "gameInfo", "message": {"gameInfo": ` + newMessage + `}, "gameId": "` + playfield.GameID + `"}`
	msg["message"] = infoMessage

	*c2 <- msg
}
