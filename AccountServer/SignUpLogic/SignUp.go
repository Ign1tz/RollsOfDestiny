package SignUpLogic

import (
	"RollsOfDestiny/AccountServer/Database"
	"RollsOfDestiny/AccountServer/Encryption"
	"RollsOfDestiny/AccountServer/Types"
	"fmt"
	"github.com/google/uuid"
	"log"
)

func SignUpNewAccount(newInfo SignUpInfo) {
	validPassword := newInfo.ComparePassword()
	validUsername := newInfo.CheckUsername()
	validEmail := newInfo.CheckEmail()

	fmt.Println(validPassword, validUsername, validEmail)
	if validPassword && validUsername && validEmail {
		hashedPassword, _ := Encryption.HashPassword(newInfo.Password)
		newAccount := Types.Account{
			UserID:         uuid.New().String(),
			Username:       newInfo.Username,
			Password:       hashedPassword,
			Email:          newInfo.Email,
			ProfilePicture: "https://via.placeholder.com/100",
			Rating:         1000,
		}
		err := Database.InsertAccount(newAccount)
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}
}
