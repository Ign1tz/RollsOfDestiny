package main

import (
	"RollsOfDestiny/GameServer/Database"
	"fmt"
)

func main() {
	//GameLogic.GameLoop()
	Database.InitDatabase()
	//player := Types.Player{UserID: "5678", Username: "tester2", Mana: 10}
	//err := Database.InsertPlayer(player)
	/*err := Database.InsertGame(Types.Game{GameID: "game1234", Host: "1234", Guest: "5678", ActivePlayer: "1234", HostGrid: "grid1234", GuestGrid: "grid1234"})
	if err != nil {
		panic(err)
	}*/

	newPlayer, err := Database.GetGrid("grid1234")
	if err != nil {
		panic(err)
	}
	fmt.Println(newPlayer.Middle)
	fmt.Println(newPlayer.Left)
	fmt.Println(newPlayer.Right)
	//Database.Database()
	//a()
}
