package main

import (
	"RollsOfDestiny/AccountServer/Database"
	"RollsOfDestiny/AccountServer/SignUpLogic"
	"fmt"
)

func main() {
	Database.InitDatabase()
	account := SignUpLogic.SignUpInfo{
		Username:        "test",
		Email:           "test",
		Password:        "testtest",
		ConfirmPassword: "testtest",
	}
	fmt.Println(account.CheckUsername())
	fmt.Println(account.CheckEmail())
	fmt.Println(account.ComparePassword())
	//go Server.Server()
	//Server2.Server()
	//GameLogic.GameLoop()
	/*Database.InitDatabase()
	//player := Types.Player{UserID: "5678", Username: "tester2", Mana: 10}
	//err := Database.InsertPlayer(player)
	err := Database.InsertGame(Types.Game{GameID: "game1234", HostId: "1234", Guest: "5678", ActivePlayer: "1234", HostGrid: "grid1234", GuestGrid: "grid1234"})
	if err != nil {
		panic(err)
	}

	newPlayer, err := Database.GetPlayfield("game1234")
	if err != nil {
		panic(err)
	}
	fmt.Println(newPlayer.Guest.Deck.Cards[0].Name)*/
	//Database.Database()
	//a()
}
