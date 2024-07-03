package Database

import (
	"RollsOfDestiny/GameServer/Types"
	"fmt"
	"log"
)

func InsertPlayer(player Types.Player) error {
	_, err := Database.Exec("INSERT INTO players (userid, username, mana, gridid, websocketconnectionid) VALUES ($1, $2, $3, $4, $5)",
		player.UserID, player.Username, player.Mana, player.Grid.GridId, player.WebsocketConnectionID)
	return err
}

func InsertDeck(deck Types.Deck) error {
	_, err := Database.Exec("INSERT INTO decks VALUES ($1, $2)", deck.DeckID, deck.UserID)
	return err
}

func InsertCard(card Types.Card) error {
	_, err := Database.Exec(
		"INSERT INTO cards VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		card.Name, card.Cost, card.Effect, card.Picture, card.CardID, card.DeckID, card.Played, card.InHand)
	return err
}

func InsertColumn(column Types.Column) error {
	fmt.Println(column.First, column.Second, column.Third, column.GridId, column.Placement)
	_, err := Database.Exec("INSERT INTO columns (gridid, first, second, third, placement) VALUES ($1, $2, $3, $4, $5)",
		column.GridId, column.First, column.Second, column.Third, column.Placement)
	return err
}

func InsertGrid(grid Types.Grid) error {
	_, err := Database.Exec("INSERT INTO grids VALUES ($1)",
		grid.GridId)
	return err
}

func InsertGame(game Types.Game) error {
	fmt.Println("insert roll", game.LastRoll)
	_, err := Database.Exec("INSERT INTO games (gameid, host, guest, activeplayer, hostgrid, guestgrid, lastroll) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		game.GameID, game.HostId, game.GuestId, game.ActivePlayer,
		game.HostGrid, game.GuestGrid, game.LastRoll)
	fmt.Println(err)
	return err
}

func UndefinedDelete(table string, key string, value string) error {
	_, err := Database.Exec("Delete From $1 Where $2 = $3", table, key, value)
	return err
}

func InsertWholeGame(playfield Types.Playfield) error {
	err := InsertGrid(playfield.HostGrid)
	if err != nil {
		return err
	}
	err = InsertGrid(playfield.GuestGrid)
	if err != nil {
		return err
	}

	err = InsertColumn(playfield.HostGrid.Left)
	if err != nil {
		return err
	}
	err = InsertColumn(playfield.HostGrid.Middle)
	if err != nil {
		return err
	}
	err = InsertColumn(playfield.HostGrid.Right)
	if err != nil {
		return err
	}
	err = InsertColumn(playfield.GuestGrid.Left)
	if err != nil {
		return err
	}
	err = InsertColumn(playfield.GuestGrid.Middle)
	if err != nil {
		return err
	}
	err = InsertColumn(playfield.GuestGrid.Right)
	if err != nil {
		return err
	}

	err = InsertPlayer(playfield.Host)
	if err != nil {
		return err
	}
	err = InsertPlayer(playfield.Guest)
	if err != nil {
		return err
	}

	log.Println(playfield.Host.Deck.DeckID)
	log.Println(playfield.Guest.Deck.DeckID)
	if playfield.Host.Deck.DeckID != "" {
		err = InsertDeck(playfield.Host.Deck)
		if err != nil {
			return err
		}
		for cardIndex := range playfield.Host.Deck.Cards {
			if playfield.Host.Deck.Cards[cardIndex].CardID != "" {
				err = InsertCard(playfield.Host.Deck.Cards[cardIndex])
				if err != nil {
					return err
				}
			}
		}
	}
	if playfield.Guest.Deck.DeckID != "" {
		err = InsertDeck(playfield.Guest.Deck)
		if err != nil {
			return err
		}
		for cardIndex := range playfield.Guest.Deck.Cards {
			if playfield.Guest.Deck.Cards[cardIndex].CardID != "" {
				err = InsertCard(playfield.Guest.Deck.Cards[cardIndex])
				if err != nil {
					return err
				}
			}
		}
	}

	game := Types.Game{
		HostId:       playfield.Host.UserID,
		GuestId:      playfield.Guest.UserID,
		HostGrid:     playfield.HostGrid.GridId,
		GuestGrid:    playfield.GuestGrid.GridId,
		GameID:       playfield.GameID,
		ActivePlayer: playfield.ActivePlayer.UserID,
		LastRoll:     playfield.LastRoll,
	}

	err = InsertGame(game)
	return err
}

func InsertPosition(position Types.Position) error {
	_, err := Database.Exec("INSERT INTO position VALUES ($1, $2, $3, $4)",
		position.Gameid, position.CurrentStep, position.HostInfo, position.GuestInfo)
	return err
}
