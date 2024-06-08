package Database

import "RollsOfDestiny/GameServer/Types"

func GetDBPlayer(playerId string) (Types.Player, error) {
	dbPlayer := Database.QueryRow("Select * from players where userid = $1", playerId)
	var player Types.Player
	if err := dbPlayer.Scan(&player.UserID, &player.Username, &player.Mana); err != nil {
		return Types.Player{}, err
	}
	return player, nil
}

func GetGame(gameId string) (Types.Game, error) {
	dbGame := Database.QueryRow("Select * from games where gameid = $1", gameId)
	var game Types.Game
	if err := dbGame.Scan(&game.GameID, &game.HostId, &game.Guest, &game.ActivePlayer, &game.HostGrid, &game.GuestGrid, &game.LastRoll); err != nil {
		return Types.Game{}, err
	}
	return game, nil
}

func GetGrid(gridId string) (Types.Grid, error) {
	dbGrid := Database.QueryRow("Select * from grids where gridid = $1", gridId)
	var grid Types.Grid
	if err := dbGrid.Scan(&grid.GridId); err != nil {
		return Types.Grid{}, err
	}
	dbColumns, err := Database.Query("SELECT * from columns where gridid = $1", grid.GridId)
	if err != nil {
		return grid, err
	}
	for dbColumns.Next() {
		var col Types.Column
		var placement string
		if err := dbColumns.Scan(&col.GridId, &col.First, &col.Second, &col.Third, &placement); err != nil {
			return grid, err
		}
		if placement == "left" {
			grid.Left = col
		} else if placement == "right" {
			grid.Right = col
		} else if placement == "middle" {
			grid.Middle = col
		}
	}
	return grid, nil
}

func GetDeckByDeckId(deckId string) (Types.Deck, error) {
	dbDeck := Database.QueryRow("Select * from decks where deckid = $1", deckId)
	var deck Types.Deck
	if err := dbDeck.Scan(&deck.DeckID, &deck.UserID); err != nil {
		return Types.Deck{}, err
	}
	dbCards, err := Database.Query("select * from cards where deckid = $1", deck.DeckID)
	if err != nil {
		return deck, err
	}

	for dbCards.Next() {
		var card Types.Card

		if err := dbCards.Scan(&card.Name, &card.Cost, &card.Effect, &card.Picture, card.CardID, card.DeckID, card.Played, card.InHand); err != nil {
			return deck, err
		}
		deck.Cards = append(deck.Cards, card)
	}
	deck.Size = len(deck.Cards)
	return deck, nil
}

func GetDeckByPlayerId(playerId string) (Types.Deck, error) {
	dbDeck := Database.QueryRow("Select * from decks where playerid = $1", playerId)
	var deck Types.Deck
	if err := dbDeck.Scan(&deck.DeckID, &deck.UserID); err != nil {
		return Types.Deck{}, err
	}
	dbCards, err := Database.Query("select * from cards where deckid = $1", deck.DeckID)
	if err != nil {
		return deck, err
	}

	for dbCards.Next() {
		var card Types.Card

		if err := dbCards.Scan(&card.Name, &card.Cost, &card.Effect, &card.Picture, &card.CardID, &card.DeckID, &card.Played, &card.InHand); err != nil {
			return deck, err
		}
		deck.Cards = append(deck.Cards, card)
	}
	deck.Size = len(deck.Cards)
	return deck, nil
}

func GetPlayer(playerId string) (Types.Player, error) {
	dbPlayer := Database.QueryRow("Select * from players where userid = $1", playerId)
	var gridId string
	var player Types.Player
	if err := dbPlayer.Scan(&player.UserID, &player.Username, &player.Mana, &gridId); err != nil {
		return Types.Player{}, err
	}
	deck, err := GetDeckByPlayerId(player.UserID)
	if err != nil {
		return player, err
	}
	player.Deck = deck

	grid, err := GetGrid(gridId)
	if err != nil {
		return player, err
	}
	player.Grid = grid

	return player, nil
}

func GetPlayfield(gameId string) (Types.Playfield, error) {
	dbGame := Database.QueryRow("Select * from games where gameid = $1", gameId)
	var game Types.Game
	if err := dbGame.Scan(&game.GameID, &game.HostId, &game.Guest, &game.ActivePlayer, &game.HostGrid, &game.GuestGrid, &game.LastRoll); err != nil {
		return Types.Playfield{}, err
	}
	var playfield Types.Playfield
	playfield.GameID = game.GameID

	activePlayer, err := GetPlayer(game.ActivePlayer)
	if err != nil {
		return playfield, err
	}
	playfield.ActivePlayer = activePlayer
	hostPlayer, err := GetPlayer(game.HostId)
	if err != nil {
		return playfield, err
	}
	playfield.Host = hostPlayer
	guestPlayer, err := GetPlayer(game.Guest)
	if err != nil {
		return playfield, err
	}
	playfield.Guest = guestPlayer
	hostGrid, err := GetGrid(game.HostGrid)
	if err != nil {
		return playfield, err
	}
	playfield.HostGrid = hostGrid
	guestGrid, err := GetGrid(game.GuestGrid)
	if err != nil {
		return playfield, err
	}
	playfield.GuestGrid = guestGrid
	return playfield, nil
}
