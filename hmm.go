package main

import (
	"RollsOfDestiny/AccountServer/Database"
	Server2 "RollsOfDestiny/AccountServer/Server"
	Database2 "RollsOfDestiny/GameServer/Database"
	"RollsOfDestiny/GameServer/Server"
	"fmt"
	"github.com/joho/godotenv"
	"regexp"
)

func main() {
	emailRegex, _ := regexp.Compile(`^\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,3}$`)
	fmt.Println(emailRegex.MatchString("tester@test.tes"))

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

	go Server.Server()
	Server2.Server()
}
