package Server

import (
	"RollsOfDestiny/GameServer/Database"
	"RollsOfDestiny/GameServer/Types"
	"github.com/google/uuid"
	"log"
)

func AddToQueue(queueEntry Types.QueueInfo, c2 *chan map[string]string) {
	if alreadyInGame(queueEntry.UserId) {
		err := Database.UpdatePlayerWebsocketID(queueEntry.UserId, queueEntry.WebsocketConnectionId)
		if err != nil {
			panic(err)
			return
		}
	} else {
		player, _ := Database.GetOldestEntry()

		if player.UserId == "" {
			err := Database.AddToQueueDatabase(queueEntry)
			if err != nil {
			}
		} else if player.UserId == queueEntry.UserId {
			err := Database.UpdateQueueEntry(queueEntry.UserId, queueEntry.WebsocketConnectionId)
			if err != nil {
				log.Println(err)
			}
		} else {
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

			playfield := Types.Playfield{
				Host:         host,
				Guest:        guest,
				HostGrid:     hostGrid,
				GuestGrid:    guestGrid,
				GameID:       uuid.New().String(),
				ActivePlayer: host,
			}

			err := Database.InsertWholeGame(playfield)
			if err != nil {
				Database.DeleteGame(hostGrid.GridId)
				Database.DeleteGame(guestGrid.GridId)
				log.Println(err)
			}
			Database.DeleteFromQueue(player)

			var msg = make(map[string]string)

			msg["id"] = host.WebsocketConnectionID
			message := `{"gameid": "` + playfield.GameID + `", "YourInfo": { "WebsocketId": "` + playfield.Host.WebsocketConnectionID + `", "Username": "` + playfield.Host.Username + `"}, "EnemyInfo": { "WebsocketId": "` + playfield.Guest.WebsocketConnectionID + `", "Username": "` + playfield.Guest.Username + `"}}`
			msg["message"] = message

			*c2 <- msg

			var msg2 = make(map[string]string)
			msg2["id"] = guest.WebsocketConnectionID
			message = `{"gameid": "` + playfield.GameID + `", "YourInfo": { "WebsocketId": "` + playfield.Guest.WebsocketConnectionID + `", "Username": "` + playfield.Guest.Username + `"}, "EnemyInfo": { "WebsocketId": "` + playfield.Host.WebsocketConnectionID + `", "Username": "` + playfield.Host.Username + `"}}`
			msg2["message"] = message

			*c2 <- msg2
		}
	}
}

func alreadyInGame(userID string) bool {
	_, err := Database.GetDBPlayer(userID)
	return err == nil
}
