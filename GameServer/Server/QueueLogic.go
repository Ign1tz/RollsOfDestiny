package Server

import (
	"RollsOfDestiny/GameServer/Database"
	"RollsOfDestiny/GameServer/Types"
	"fmt"
	"github.com/google/uuid"
)

func AddToQueue(queueEntry Types.QueueInfo, c2 *chan map[string]string) {
	fmt.Println("Adding to queue")
	if alreadyInGame(queueEntry.UserId) {
		fmt.Println("Already in game")
		err := Database.UpdatePlayerWebsocketID(queueEntry.UserId, queueEntry.WebsocketConnectionId)
		if err != nil {
			panic(err)
			return
		}
	} else {
		fmt.Println("Not in game")
		player, _ := Database.GetOldestEntry()

		fmt.Println("userid:", player.UserId)

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
				Grid:                  hostGrid,
			}
			guest := Types.Player{
				Username:              "Guest",
				UserID:                queueEntry.UserId,
				Mana:                  0,
				Deck:                  Types.Deck{},
				Die:                   Types.Die{PossibleThrows: []int{1, 2, 3, 4, 5, 6}},
				WebsocketConnectionID: queueEntry.WebsocketConnectionId,
				Grid:                  guestGrid,
			}

			fmt.Println("After creating player")
			diceTrow := Types.Die{PossibleThrows: []int{1, 2, 3, 4, 5, 6}}.Throw()
			fmt.Println("diceTh", diceTrow)
			playfield := Types.Playfield{
				Host:         host,
				Guest:        guest,
				HostGrid:     hostGrid,
				GuestGrid:    guestGrid,
				GameID:       uuid.New().String(),
				ActivePlayer: host,
				LastRoll:     diceTrow,
			}

			fmt.Println("After creating stuff")
			err := Database.InsertWholeGame(playfield)
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

			var msg2 = make(map[string]string)
			msg2["id"] = playfield.Guest.WebsocketConnectionID
			newMessage = `{"gameid": "` + playfield.GameID + `", "YourInfo": ` + playfield.Guest.ToJson(true) + `, "EnemyInfo":` + playfield.Host.ToJson(false) + `, "ActivePlayer": {"active": false, "roll": "` + playfield.LastRoll + `"}}`
			infoMessage = `{"info": "gameInfo", "message": {"gameInfo": ` + newMessage + `}, "gameId": "` + playfield.GameID + `"}`
			msg2["message"] = infoMessage

			*c2 <- msg2
			fmt.Println("second", msg2["id"])
			fmt.Println("After second message")
		}
	}
}

func alreadyInGame(userID string) bool {
	_, err := Database.GetDBPlayer(userID)
	return err == nil
}
