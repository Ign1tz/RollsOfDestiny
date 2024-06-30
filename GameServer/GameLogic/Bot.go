package GameLogic

import (
	"RollsOfDestiny/GameServer/Database"
	"RollsOfDestiny/GameServer/Types"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
)

func BotTurn(gameInfo Types.Resp) {

	err := PickColumn(gameInfo.Gameid, gameInfo.ColumnKey)

	if err != nil {
		panic(err)
	}

	gamefield, err := Database.GetPlayfield(gameInfo.Gameid)

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
				pickedValidColumn = true
			}
		} else if columnNumber == 1 {
			columnErr = gamefield.ActivePlayer.Grid.Middle.Add(roll)
			if columnErr == nil {
				enemy.Grid.Middle.Remove(roll)
				pickedValidColumn = true
			}
		} else if columnNumber == 2 {
			columnErr = gamefield.ActivePlayer.Grid.Middle.Add(roll)
			if columnErr == nil {
				enemy.Grid.Right.Remove(roll)
				pickedValidColumn = true
			}
		}
	}
	gamefield.ActivePlayer = enemy
}

func BotStartGame(queueEntry Types.BotResp) {
	fmt.Println("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	gridId1 := uuid.New().String()
	hostGrid := Types.Grid{
		Left: Types.Column{
			First:     0,
			Second:    0,
			Third:     0,
			IsFull:    false,
			GridId:    gridId1,
			Placement: 0,
		},
		Middle: Types.Column{
			First:     0,
			Second:    0,
			Third:     0,
			IsFull:    false,
			GridId:    gridId1,
			Placement: 1,
		},
		Right: Types.Column{
			First:     0,
			Second:    0,
			Third:     0,
			IsFull:    false,
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
			IsFull:    false,
			GridId:    gridId2,
			Placement: 0,
		},
		Middle: Types.Column{
			First:     0,
			Second:    0,
			Third:     0,
			IsFull:    false,
			GridId:    gridId2,
			Placement: 1,
		},
		Right: Types.Column{
			First:     0,
			Second:    0,
			Third:     0,
			IsFull:    false,
			GridId:    gridId2,
			Placement: 2,
		},
		GridId: gridId2,
	}
	fmt.Println("After creating grid")

	host := Types.Player{
		Username:              "Host",
		UserID:                queueEntry.Userid,
		Mana:                  0,
		Deck:                  Types.Deck{},
		Die:                   Types.Die{PossibleThrows: []int{1, 2, 3, 4, 5, 6}},
		WebsocketConnectionID: "",
		Grid:                  hostGrid,
	}
	guest := Types.Player{
		Username:              "Bot",
		UserID:                uuid.New().String(),
		Mana:                  0,
		Deck:                  Types.Deck{},
		Die:                   Types.Die{PossibleThrows: []int{1, 2, 3, 4, 5, 6}},
		WebsocketConnectionID: "",
		Grid:                  guestGrid,
	}

	fmt.Println("After creating player")

	playfield := Types.Playfield{
		Host:         host,
		Guest:        guest,
		HostGrid:     hostGrid,
		GuestGrid:    guestGrid,
		GameID:       "bot: " + uuid.New().String(),
		ActivePlayer: host,
	}

	fmt.Println("After creating stuff")
	err := Database.InsertWholeGame(playfield)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
