package Database

import (
	"database/sql"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/url"
	"os"
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
	serviceURI := os.Getenv("DATABASE_URL")

	conn, _ := url.Parse(serviceURI)
	conn.RawQuery = "sslmode=verify-ca;sslrootcert=ca.pem"
	//var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)
	db, err := sql.Open("postgres", conn.String())
	if err != nil {
		panic(err)
	}
	Database = db
	return db
}
