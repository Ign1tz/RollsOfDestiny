package main

import (
	"RollsOfDestiny/AccountServer/Database"
	Server2 "RollsOfDestiny/AccountServer/Server"
	Database2 "RollsOfDestiny/GameServer/Database"
	"RollsOfDestiny/GameServer/Server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	log.Println("Starting server")
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	Database.InitDatabase()
	Database2.InitDatabase()

	go Server.Server()
	Server2.Server()

}
