package Database

import (
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
