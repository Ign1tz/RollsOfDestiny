package main

import (
	"RollsOfDestiny/AccountServer/Database"
	"RollsOfDestiny/AccountServer/Encryption"
	Server2 "RollsOfDestiny/AccountServer/Server"
	Database2 "RollsOfDestiny/GameServer/Database"
	"RollsOfDestiny/GameServer/Server"
	"fmt"
)

func main() {
	password := "secret"
	hash, _ := Encryption.HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := Encryption.CheckPasswordHash("secret", hash)
	fmt.Println("Match:   ", match)
	Database.InitDatabase()
	Database2.InitDatabase()
	/*account := SignUpLogic.SignUpInfo{
		Username:        "test",
		Email:           "test",
		Password:        "testtest",
		ConfirmPassword: "testtest",
	}
	fmt.Println(account.CheckUsername())
	fmt.Println(account.CheckEmail())
	fmt.Println(account.ComparePassword())*/

	go Server.Server()
	Server2.Server()
}
