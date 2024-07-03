package main

import (
	"RollsOfDestiny/AccountServer/Database"
	Server2 "RollsOfDestiny/AccountServer/Server"
	Database2 "RollsOfDestiny/GameServer/Database"
	"RollsOfDestiny/GameServer/Server"
	"RollsOfDestiny/GameServer/Types"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	Database.InitDatabase()
	Database2.InitDatabase()
	//Database2.DatabaseTest()
	/*account := SignUpLogic.SignUpInfo{
		Username:        "test",
		Email:           "test",
		Password:        "testtest",
		ConfirmPassword: "testtest",
	}
	fmt.Println(account.CheckUsername())
	fmt.Println(account.CheckEmail())
	fmt.Println(account.ComparePassword())*/

	grid := Types.Grid{
		Left: Types.Column{
			First:     1,
			Second:    2,
			Third:     3,
			GridId:    "",
			Placement: 0,
		},
		Middle: Types.Column{
			First:     4,
			Second:    5,
			Third:     6,
			GridId:    "",
			Placement: 0},
		Right: Types.Column{
			First:     7,
			Second:    8,
			Third:     9,
			GridId:    "",
			Placement: 0},
		GridId: "",
	}
	preGrid := grid
	preGrid.PrettyPrint()
	grid.FlipClocwise()
	preGrid = grid.CheckGridForOverlap(preGrid)
	grid.PrettyPrint()
	preGrid.PrettyPrint()

	go Server.Server()
	Server2.Server()

}
