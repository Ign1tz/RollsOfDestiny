package Server

import (
	"RollsOfDestiny/GameServer/Database"
	"RollsOfDestiny/GameServer/Types"
	"fmt"
	"github.com/google/uuid"
)

func AddToQueue(queueEntry Types.QueueInfo, c2 *chan map[string]string) {
	player, _ := Database.GetOldestEntry()

	fmt.Println(player.UserId)
	if player.UserId == "" {
		fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		err := Database.AddToQueueDatabase(queueEntry)
		if err != nil {
			fmt.Println(err)
		}
	} else {
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
			UserID:                player.UserId,
			Mana:                  0,
			Deck:                  Types.Deck{},
			Die:                   Types.Die{PossibleThrows: []int{1, 2, 3, 4, 5, 6}},
			WebsocketConnectionID: player.WebsocketConnectionId,
			Gridid:                hostGrid.GridId,
		}
		guest := Types.Player{
			Username:              "GuestId",
			UserID:                queueEntry.UserId,
			Mana:                  0,
			Deck:                  Types.Deck{},
			Die:                   Types.Die{PossibleThrows: []int{1, 2, 3, 4, 5, 6}},
			WebsocketConnectionID: queueEntry.WebsocketConnectionId,
			Gridid:                guestGrid.GridId,
		}

		fmt.Println("After creating player")

		playfield := Types.Playfield{
			Host:         host,
			Guest:        guest,
			HostGrid:     hostGrid,
			GuestGrid:    guestGrid,
			GameID:       uuid.New().String(),
			ActivePlayer: host,
		}

		fmt.Println("After creating stuff")
		err := Database.InsertWholeGame(playfield)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		var msg = make(map[string]string)

		msg["id"] = host.WebsocketConnectionID
		message := `{"gameid": "` + playfield.GameID + `", "YourInfo": { "WebsocketId": "` + playfield.Host.WebsocketConnectionID + `", "Username": "` + playfield.Host.Username + `"}, "EnemyInfo": { "WebsocketId": "` + playfield.Guest.WebsocketConnectionID + `", "Username": "` + playfield.Guest.Username + `"}}`
		msg["message"] = message

		*c2 <- msg

		var msg2 = make(map[string]string)
		fmt.Println("first", msg["id"])
		fmt.Println("After first message")
		msg2["id"] = guest.WebsocketConnectionID
		message = `{"gameid": "` + playfield.GameID + `", "YourInfo": { "WebsocketId": "` + playfield.Guest.WebsocketConnectionID + `", "Username": "` + playfield.Guest.Username + `"}, "EnemyInfo": { "WebsocketId": "` + playfield.Host.WebsocketConnectionID + `", "Username": "` + playfield.Host.Username + `"}}`
		msg2["message"] = message

		*c2 <- msg2
		fmt.Println("second", msg2["id"])
		fmt.Println("After second message")
	}
}
