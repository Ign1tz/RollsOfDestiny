package Database

import (
	"RollsOfDestiny/GameServer/Types"
	_ "RollsOfDestiny/GameServer/Types"
	"database/sql"
	"fmt"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "47XC#fFMhy4$bmPa"
	dbname   = "postgres"
)

var Database *sql.DB

func InitDatabase() *sql.DB {
	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	Database = db
	return db
}

func InsertPlayer(player Types.Player) error {
	_, err := Database.Exec("INSERT INTO players (userid, username, mana) VALUES ($1, $2, $3)",
		player.UserID, player.Username, player.Mana)
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
	_, err := Database.Exec("INSERT INTO columns VALUES ($1, $2, $3, $4, $5)",
		column.GridId, column.First, column.Second, column.Third, column.Placement)
	return err
}

func InsertGrid(grid Types.Grid) error {
	_, err := Database.Exec("INSERT INTO grids VALUES ($1)",
		grid.GridId)
	return err
}

func InsertGame(game Types.Game) error {
	_, err := Database.Exec("INSERT INTO games VALUES ($1, $2, $3, $4, $5, $6)",
		game.GameID, game.HostId, game.Guest, game.ActivePlayer,
		game.HostGrid, game.GuestGrid)
	return err
}

func UndefinedDelete(table string, key string, value string) error {
	_, err := Database.Exec("Delete From $1 Where $2 = $3", table, key, value)
	return err
}

func Update(query string) error {
	_, err := Database.Exec(query)
	return err
}

/*func getGameState(gameId string) (Types.Playfield, error) {
	gameSQL, err := Database.Query("select * from games where gameid = $GameId", gameId)
	gameSQL.Scan(Types.Playfield{})
	hostSQL, err := Database.Query("select * from players where userid = $UserId", gameSQL)

	return , err
}*/

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
	if err := dbGame.Scan(&game.GameID, &game.HostId, &game.Guest, &game.ActivePlayer, &game.HostGrid, &game.GuestGrid); err != nil {
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
	var player Types.Player
	if err := dbPlayer.Scan(&player.UserID, &player.Username, &player.Mana); err != nil {
		return Types.Player{}, err
	}
	deck, err := GetDeckByPlayerId(player.UserID)
	if err != nil {
		return player, err
	}
	player.Deck = deck

	return player, nil
}

func GetPlayfield(gameId string) (Types.Playfield, error) {
	dbGame := Database.QueryRow("Select * from games where gameid = $1", gameId)
	var game Types.Game
	if err := dbGame.Scan(&game.GameID, &game.HostId, &game.Guest, &game.ActivePlayer, &game.HostGrid, &game.GuestGrid); err != nil {
		return Types.Playfield{}, err
	}
	var playfield Types.Playfield
	playfield.GameID = game.GameID
	playfield.ActivePlayer = game.ActivePlayer
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

func DatabaseTest() {
	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db.Exec("Insert into players Values ('testID', 'testName', 5)")
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Established a successful connection!")
}
