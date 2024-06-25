package main

import (
	"RollsOfDestiny/AccountServer/Encryption"
	"fmt"
)

func main() {
	password := "secret"
	hash, _ := Encryption.HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := Encryption.CheckPasswordHash("secret", hash)
	fmt.Println("Match:   ", match)
	//Database.InitDatabase()
	/*account := SignUpLogic.SignUpInfo{
		Username:        "test",
		Email:           "test",
		Password:        "testtest",
		ConfirmPassword: "testtest",
	}
	fmt.Println(account.CheckUsername())
	fmt.Println(account.CheckEmail())
	fmt.Println(account.ComparePassword())*/

	//go Server.Server()
	//Server2.Server()
}
