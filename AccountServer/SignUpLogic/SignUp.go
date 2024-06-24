package SignUpLogic

import (
	"RollsOfDestiny/AccountServer/Database"
	"RollsOfDestiny/AccountServer/Types"
	"github.com/google/uuid"
	"log"
)

func SignUpNewAccount(newInfo SignUpInfo) {
	validPassword := newInfo.ComparePassword()
	validUsername := newInfo.CheckUsername()
	validEmail := newInfo.CheckEmail()

	if validPassword && validUsername && validEmail {
		newAccount := Types.Account{
			UserID:         uuid.New().String(),
			Username:       newInfo.Username,
			Password:       newInfo.Password,
			Email:          newInfo.Email,
			ProfilePicture: "default",
			Rating:         1000,
		}
		err := Database.InsertAccount(newAccount)
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}
}
