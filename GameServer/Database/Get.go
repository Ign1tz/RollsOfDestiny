package Database

import (
	"RollsOfDestiny/GameServer/Types"
	"fmt"
	"log"
)

func GetDBPlayer(playerId string) (Types.Player, error) { //ONLY USED FOR TESTING IF ALREADY IN GAME
	dbPlayer := Database.QueryRow("Select * from players where userid = $1", playerId)
	var player Types.Player
	var gridId string
	if err := dbPlayer.Scan(&player.UserID, &player.Username, &player.Mana, &gridId, &player.WebsocketConnectionID); err != nil {
		return Types.Player{}, err
	}
	return player, nil
}

func GetGame(gameId string) (Types.Game, error) {
	dbGame := Database.QueryRow("Select * from games where gameid = $1", gameId)
	var game Types.Game
	if err := dbGame.Scan(&game.GameID, &game.HostId, &game.GuestId, &game.ActivePlayer, &game.HostGrid, &game.GuestGrid, &game.LastRoll); err != nil {
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
		var placement int
		if err := dbColumns.Scan(&col.GridId, &placement, &col.First, &col.Second, &col.Third); err != nil {
			return grid, err
		}
		if placement == 0 {
			grid.Left = col
			grid.Left.Placement = placement
		} else if placement == 1 {
			grid.Middle = col
			grid.Middle.Placement = placement
		} else if placement == 2 {
			grid.Right = col
			grid.Right.Placement = placement
		}
	}
	return grid, nil
}

func GetDeckByDeckId(deckId string) (Types.Deck, error) {
	log.Println("Error querying cards", deckId)
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
	var player Types.Player
	var gridid string
	if err := dbPlayer.Scan(&player.UserID, &player.Username, &player.Mana, &gridid, &player.WebsocketConnectionID); err != nil {
		return Types.Player{}, err
	}
	/*deck, err := GetDeckByPlayerId(player.UserID)
	if err != nil {
		return player, err
	}
	player.Deck = deck*/

	grid, err := GetGrid(gridid)
	if err != nil {
		return player, err
	}
	player.Grid = grid

	return player, nil
}

func GetPlayfield(gameId string) (Types.Playfield, error) {
	dbGame := Database.QueryRow("Select * from games where gameid = $1", gameId)
	var game Types.Game
	err := dbGame.Scan(&game.GameID, &game.HostId, &game.GuestId, &game.ActivePlayer, &game.HostGrid, &game.GuestGrid, &game.LastRoll)
	fmt.Println(err)
	if err != nil {
		return Types.Playfield{}, err
	}
	var playfield Types.Playfield
	playfield.GameID = game.GameID
	playfield.LastRoll = game.LastRoll
	fmt.Println(game.ActivePlayer)
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
	guestPlayer, err := GetPlayer(game.GuestId)
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

	hostDeck, err := GetDeckByPlayerId(playfield.Host.UserID)
	if err != nil {
		hostDeck = Types.Deck{}
	}
	playfield.Host.Deck = hostDeck

	guestDeck, err := GetDeckByPlayerId(playfield.Guest.UserID)
	if err != nil {
		guestDeck = Types.Deck{}
	}
	playfield.Guest.Deck = guestDeck
	activeDeck, err := GetDeckByPlayerId(playfield.ActivePlayer.UserID)
	if err != nil {
		activeDeck = Types.Deck{}
	}
	playfield.ActivePlayer.Deck = activeDeck

	return playfield, nil
}

func GetPlayfieldByUserid(userid string) (Types.Playfield, error) {

	log.Println("test", userid)
	dbGame := Database.QueryRow("Select * from games where host = $1 or guest = $1", userid)
	var game Types.Game
	err := dbGame.Scan(&game.GameID, &game.HostId, &game.GuestId, &game.ActivePlayer, &game.HostGrid, &game.GuestGrid, &game.LastRoll)
	fmt.Println(err)
	if err != nil {
		return Types.Playfield{}, err
	}
	var playfield Types.Playfield
	playfield.GameID = game.GameID
	playfield.LastRoll = game.LastRoll
	fmt.Println(game.ActivePlayer)
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
	guestPlayer, err := GetPlayer(game.GuestId)
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

	hostDeck, err := GetDeckByPlayerId(playfield.Host.UserID)
	if err != nil {
		hostDeck = Types.Deck{}
	}
	playfield.Host.Deck = hostDeck

	guestDeck, err := GetDeckByPlayerId(playfield.Guest.UserID)
	if err != nil {
		guestDeck = Types.Deck{}
	}
	playfield.Guest.Deck = guestDeck
	activeDeck, err := GetDeckByPlayerId(playfield.ActivePlayer.UserID)
	if err != nil {
		activeDeck = Types.Deck{}
	}
	playfield.ActivePlayer.Deck = activeDeck

	return playfield, nil
}

func GetDeckByDeckIDFromAccount(userid string) (Types.Deck, error) {
	dbDeck := Database.QueryRow("Select * from accountdecks where userid = $1 and active = 'true'", userid)
	var tempDeck Types.Deck
	var active bool
	if err := dbDeck.Scan(&tempDeck.UserID, &tempDeck.DeckID, &tempDeck.Name, &active); err != nil {
		return Types.Deck{}, err
	}

	return tempDeck, nil
}

func GetCardsByDeckIDFromAccount(deckID string) ([]string, error) {
	dbCards, err := Database.Query("Select name from accountcards where deckids like '%' || $1 || '%'", deckID)
	if err != nil {
		return []string{}, err
	}
	var cardNames = make([]string, 100)
	id := 0
	for dbCards.Next() {
		if err := dbCards.Scan(&cardNames[id]); err != nil {
			return []string{}, err
		}
		id++
	}
	return cardNames, err
}

func GetCardsByDeckID(deckID string) ([]Types.Card, error) {
	dbCards, err := Database.Query("Select * from cards where deckid = $1 and played = 'false'", deckID)
	if err != nil {
		return []Types.Card{}, err
	}
	var cards = make([]Types.Card, 20)
	id := 0
	for dbCards.Next() {
		if err := dbCards.Scan(&cards[id].Name, &cards[id].Cost, &cards[id].Effect, &cards[id].Picture, &cards[id].CardID, &cards[id].DeckID, &cards[id].Played, &cards[id].InHand); err != nil {
			return []Types.Card{}, err
		}
		id++
	}
	return cards, err
}

func GetCardById(cardId string) (Types.Card, error) {
	dbDeck := Database.QueryRow("Select * from cards where cardid = $1", cardId)
	var tempCard Types.Card
	if err := dbDeck.Scan(&tempCard.Name, &tempCard.Cost, &tempCard.Effect, &tempCard.Picture, &tempCard.CardID, &tempCard.DeckID, &tempCard.Played, &tempCard.InHand); err != nil {
		return Types.Card{}, err
	}

	return tempCard, nil
}

func GetPosition(gameid string) (Types.Position, error) {
	dbDeck := Database.QueryRow("Select * from position where gameid = $1", gameid)
	var tempCard Types.Position
	if err := dbDeck.Scan(&tempCard.Gameid, &tempCard.CurrentStep, &tempCard.HostInfo, &tempCard.GuestInfo); err != nil {
		return Types.Position{}, err
	}

	return tempCard, nil
}
