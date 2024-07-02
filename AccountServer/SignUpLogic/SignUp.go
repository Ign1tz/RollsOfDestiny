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
		cards := createDefaultCards(newAccount)
		for card := range cards {
			err := Database.InsertCard(cards[card])
			if err != nil {
				log.Println(err)
				return
			}
		}
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}
}

func createDefaultCards(account Types.Account) []Types.Card {
	rollAgain := Types.Card{
		UserID: account.UserID,
		Name:   "Roll Again",
		Effect: "rollAgain",
		DeckID: "",
		Count:  0,
		Cost:   3,
	}
	doubleMana := Types.Card{
		UserID: account.UserID,
		Name:   "Double Mana",
		Effect: "doubleMana",
		DeckID: "",
		Count:  0,
		Cost:   2,
	}
	destroyColumn := Types.Card{
		UserID: account.UserID,
		Name:   "Destroy Column",
		Effect: "destroyColumn",
		DeckID: "",
		Count:  0,
		Cost:   5,
	}
	flipClockwise := Types.Card{
		UserID: account.UserID,
		Name:   "Flip Clockwise",
		Effect: "flipClockwise",
		DeckID: "",
		Count:  0,
		Cost:   4,
	}
	cards := []Types.Card{rollAgain, doubleMana, destroyColumn, flipClockwise}
	return cards
}
