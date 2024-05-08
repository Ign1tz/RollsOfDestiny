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

func initDatabase() *sql.DB {
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

func insertPlayer(player Types.Player) error {
	_, err := Database.Exec("INSERT INTO players VALUES ($UserID, $Username, $Mana)",
		player.UserID, player.Username, player.Mana)
	return err
}

func insertDeck(deck Types.Deck) error {
	_, err := Database.Exec("INSERT INTO decks VALUES ($DeckID, $PlayerID)", deck.DeckID, deck.UserID)
	return err
}

func insertCard(card Types.Card) error {
	_, err := Database.Exec(
		"INSERT INTO cards VALUES ($Name, $Cost, $Effect, $Picture, $CardID, $DeckID, $Played, $Inhand)",
		card.Name, card.Cost, card.Effect, card.Picture, card.CardID, card.DeckID, card.Played, card.InHand)
	return err
}

func insertColumn(column Types.Column) error {
	_, err := Database.Exec("INSERT INTO columns VALUES ($GridId, $First, $Second, $Third)",
		column.GridId, column.First, column.Second, column.Third)
	return err
}

func insertGrid(grid Types.Grid) error {
	_, err := Database.Exec("INSERT INTO grids VALUES ($GridId)",
		grid.GridId)
	return err
}

func insertGame(playfield Types.Playfield) error {
	_, err := Database.Exec("INSERT INTO games VALUES ($GameId, $Host, $Guest, $ActivePlayer, $HostGrid, $GuestGrid)",
		playfield.GameID, playfield.Host.UserID, playfield.Guest.UserID, playfield.ActivePlayer,
		playfield.HostGrid.GridId, playfield.GuestGrid.GridId)
	return err
}

func undefinedDelete(table string, key string, value string) error {
	_, err := Database.Exec("Delete From $Table Where $Key = $UserId", table, key, value)
	return err
}

func update(query string) error {
	_, err := Database.Exec(query)
	return err
}

/*func getGameState(gameId string) (Types.Playfield, error) {
	gameSQL, err := Database.Query("select * from games where gameid = $GameId", gameId)
	gameSQL.Scan(Types.Playfield{})
	hostSQL, err := Database.Query("select * from players where userid = $UserId", gameSQL)

	return , err
}*/

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
