package Server

import (
	"RollsOfDestiny/GameServer/Database"
	"RollsOfDestiny/GameServer/Types"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func AddToQueue(queueEntry Types.QueueInfo, c2 *chan map[string]string) {
	if alreadyInGame(queueEntry.UserId) {
		err := Database.UpdatePlayerWebsocketID(queueEntry.UserId, queueEntry.WebsocketConnectionId)
		if err != nil {
			panic(err)
			return
		}
		playfield, err := Database.GetPlayfieldByUserid(queueEntry.UserId)
		log.Println("deckid", playfield.Host.Deck.DeckID)
		if playfield.Host.UserID == queueEntry.UserId {
			active := playfield.ActivePlayer.UserID == playfield.Host.UserID
			var msg = make(map[string]string)
			msg["id"] = queueEntry.WebsocketConnectionId
			newMessage := `{"gameid": "` + playfield.GameID + `", "YourInfo":` + playfield.Host.ToJson(true) + `, "EnemyInfo": ` + playfield.Guest.ToJson(false) + `, "ActivePlayer": {"active": ` + strconv.FormatBool(active) + `, "roll": "` + playfield.LastRoll + `"}}`
			infoMessage := `{"info": "gameInfo", "message": {"gameInfo": ` + newMessage + `}, "gameId": "` + playfield.GameID + `"}`
			msg["message"] = infoMessage
			log.Println(infoMessage)
			*c2 <- msg
		} else {
			active := playfield.ActivePlayer.UserID == playfield.Guest.UserID
			var msg2 = make(map[string]string)
			msg2["id"] = queueEntry.WebsocketConnectionId
			newMessage := `{"gameid": "` + playfield.GameID + `", "YourInfo": ` + playfield.Guest.ToJson(true) + `, "EnemyInfo":` + playfield.Host.ToJson(false) + `, "ActivePlayer": {"active": ` + strconv.FormatBool(active) + `, "roll": "` + playfield.LastRoll + `"}}`
			infoMessage := `{"info": "gameInfo", "message": {"gameInfo": ` + newMessage + `}, "gameId": "` + playfield.GameID + `"}`
			msg2["message"] = infoMessage
			log.Println(infoMessage)

			*c2 <- msg2
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

			hostDeckInfo, err := Database.GetDeckByDeckIDFromAccount(player.UserId)
			if err != nil {
				log.Println("first Deck", err)
			}
			hostCards := []Types.Card{}
			if hostDeckInfo.DeckID != "" {
				hostCard, err := Database.GetCardsByDeckIDFromAccount(hostDeckInfo.DeckID)
				if err != nil {
					log.Println(err)
				}

				hostCards = createCards(hostCard, hostDeckInfo.DeckID)
			}

			hostDeck := Types.Deck{
				DeckID: hostDeckInfo.DeckID,
				Name:   hostDeckInfo.Name,
				UserID: player.UserId,
				Cards:  hostCards,
				Size:   20,
			}
			log.Println(player.UserId)
			guestDeckInfo, err := Database.GetDeckByDeckIDFromAccount(queueEntry.UserId)
			if err != nil {
				log.Println("second Deck", err)
			}

			guestCards := []Types.Card{}
			if guestDeckInfo.DeckID != "" {
				guestDeck, err := Database.GetCardsByDeckIDFromAccount(guestDeckInfo.DeckID)
				if err != nil {
					log.Println(err)
				}

				guestCards = createCards(guestDeck, guestDeckInfo.DeckID)
			}

			guestDeck := Types.Deck{
				DeckID: guestDeckInfo.DeckID,
				Name:   guestDeckInfo.Name,
				UserID: queueEntry.UserId,
				Cards:  guestCards,
				Size:   20,
			}

			host := Types.Player{
				Username:              player.Username,
				UserID:                player.UserId,
				Mana:                  2,
				Deck:                  hostDeck,
				Die:                   Types.Die{PossibleThrows: []int{1, 2, 3, 4, 5, 6}},
				WebsocketConnectionID: player.WebsocketConnectionId,
				Grid:                  hostGrid,
			}
			guest := Types.Player{
				Username:              queueEntry.Username,
				UserID:                queueEntry.UserId,
				Mana:                  2,
				Deck:                  guestDeck,
				Die:                   Types.Die{PossibleThrows: []int{1, 2, 3, 4, 5, 6}},
				WebsocketConnectionID: queueEntry.WebsocketConnectionId,
				Grid:                  guestGrid,
			}

			diceTrow := Types.Die{PossibleThrows: []int{1, 2, 3, 4, 5, 6}}.Throw()
			playfield := Types.Playfield{
				Host:         host,
				Guest:        guest,
				HostGrid:     hostGrid,
				GuestGrid:    guestGrid,
				GameID:       uuid.New().String(),
				ActivePlayer: host,
				LastRoll:     diceTrow,
			}

			err = Database.InsertWholeGame(playfield)
			if err != nil {
				log.Println(err)
			}
			position := Types.Position{
				Gameid:      playfield.GameID,
				CurrentStep: "started",
				HostInfo:    "",
				GuestInfo:   "",
			}
			err = Database.InsertPosition(position)
			if err != nil {
				Database.DeleteGame(hostGrid.GridId)
				Database.DeleteGame(guestGrid.GridId)
				log.Println(err)
			}
			Database.DeleteFromQueue(player)

			var msg = make(map[string]string)
			msg["id"] = playfield.Host.WebsocketConnectionID
			newMessage := `{"gameid": "` + playfield.GameID + `", "YourInfo":` + playfield.Host.ToJson(true) + `, "EnemyInfo": ` + playfield.Guest.ToJson(false) + `, "ActivePlayer": {"active": true, "roll": "` + playfield.LastRoll + `"}}`
			infoMessage := `{"info": "gameInfo", "message": {"gameInfo": ` + newMessage + `}, "gameId": "` + playfield.GameID + `"}`
			msg["message"] = infoMessage
			log.Println(infoMessage)
			*c2 <- msg

			var msg2 = make(map[string]string)
			msg2["id"] = playfield.Guest.WebsocketConnectionID
			newMessage = `{"gameid": "` + playfield.GameID + `", "YourInfo": ` + playfield.Guest.ToJson(true) + `, "EnemyInfo":` + playfield.Host.ToJson(false) + `, "ActivePlayer": {"active": false, "roll": "` + playfield.LastRoll + `"}}`
			infoMessage = `{"info": "gameInfo", "message": {"gameInfo": ` + newMessage + `}, "gameId": "` + playfield.GameID + `"}`
			msg2["message"] = infoMessage
			log.Println(infoMessage)

			*c2 <- msg2
		}
	}
}

func alreadyInGame(userID string) bool {
	_, err := Database.GetDBPlayer(userID)
	log.Println(err)
	return err == nil
}

func createCards(stringCards []string, deckid string) []Types.Card {
	var cards = make([]Types.Card, 20)
	log.Println(stringCards)
	cardCount := -1
	for index := range stringCards {
		switch stringCards[index] {
		case "Roll Again":
			for j := 0; j < 5; j++ {
				cardCount += 1
				log.Println("Card Count: ", cardCount)
				cards[cardCount] = Types.Card{
					CardID:  uuid.New().String(),
					Name:    stringCards[index],
					Cost:    4,
					Effect:  "rollAgain",
					Picture: "/static/media/roll_again.21331c0ee525eb47281c.png",
					DeckID:  deckid,
					Played:  false,
					InHand:  false,
				}
			}
		case "Double Mana":
			for j := 0; j < 5; j++ {
				cardCount += 1
				log.Println("Card Count: ", cardCount)
				cards[cardCount] = Types.Card{
					CardID:  uuid.New().String(),
					Name:    stringCards[index],
					Cost:    3,
					Effect:  "doubleMana",
					Picture: "/static/media/double_mana.7c47c6670f52b76c8fa6.png",
					DeckID:  deckid,
					Played:  false,
					InHand:  false,
				}
			}
		case "Destroy Column":
			for j := 0; j < 5; j++ {
				cardCount += 1
				cards[cardCount] = Types.Card{
					CardID:  uuid.New().String(),
					Name:    stringCards[index],
					Cost:    7,
					Effect:  "destroyColumn",
					Picture: "/static/media/destroy_column.23caf4dcff16d50757e3.png",
					DeckID:  deckid,
					Played:  false,
					InHand:  false,
				}
			}
		case "Flip Clockwise":
			for j := 0; j < 5; j++ {
				cardCount += 1
				cards[cardCount] = Types.Card{
					CardID:  uuid.New().String(),
					Name:    stringCards[index],
					Cost:    6,
					Effect:  "flipClockwise",
					Picture: "/static/media/rotate_grid.6a18f6243e59b2edf045.png",
					DeckID:  deckid,
					Played:  false,
					InHand:  false,
				}
			}
		}
	}

	cards = RandShuffle(cards)

	cards[0].InHand = true
	cards[1].InHand = true
	cards[2].InHand = true
	cards[3].InHand = true

	return cards
}

func RandShuffle(a []Types.Card) []Types.Card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}
