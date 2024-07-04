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
			ProfilePicture: defaultImage(),
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
		Count:  1,
		Cost:   4,
		Image:  "/static/media/roll_again.21331c0ee525eb47281c.png",
	}
	doubleMana := Types.Card{
		UserID: account.UserID,
		Name:   "Double Mana",
		Effect: "doubleMana",
		DeckID: "",
		Count:  1,
		Cost:   3,
		Image:  "/static/media/double_mana.7c47c6670f52b76c8fa6.png",
	}
	destroyColumn := Types.Card{
		UserID: account.UserID,
		Name:   "Destroy Column",
		Effect: "destroyColumn",
		DeckID: "",
		Count:  1,
		Cost:   7,
		Image:  "/static/media/destroy_column.23caf4dcff16d50757e3.png",
	}
	flipClockwise := Types.Card{
		UserID: account.UserID,
		Name:   "Flip Clockwise",
		Effect: "flipClockwise",
		DeckID: "",
		Count:  1,
		Cost:   6,
		Image:  "/static/media/rotate_grid.6a18f6243e59b2edf045.png",
	}
	cards := []Types.Card{rollAgain, doubleMana, destroyColumn, flipClockwise}
	return cards
}
