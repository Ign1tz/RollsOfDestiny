package main

import "RollsOfDestiny/GameServer/Server"

func main() {
	Server.Server()
	//GameLogic.GameLoop()
	/*Database.InitDatabase()
	//player := Types.Player{UserID: "5678", Username: "tester2", Mana: 10}
	//err := Database.InsertPlayer(player)
	/*err := Database.InsertGame(Types.Game{GameID: "game1234", HostId: "1234", GuestId: "5678", ActivePlayer: "1234", HostGrid: "grid1234", GuestGrid: "grid1234"})
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
