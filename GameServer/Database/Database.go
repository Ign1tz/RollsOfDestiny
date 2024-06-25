package Database

import (
	"database/sql"
	"fmt"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
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

func DatabaseTest() {
	serviceURI := os.Getenv("DATABASE_URL")

	conn, _ := url.Parse(serviceURI)
	conn.RawQuery = "sslmode=verify-ca;sslrootcert=ca.pem"

	var psqlInfo = fmt.Sprintf("user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"avnadmin", "AVNS_LLYDbJlnEIHL8CxAH1z", "defaultdb")
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT version()")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var result string
		err = rows.Scan(&result)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Version: %s\n", result)
	}
}
