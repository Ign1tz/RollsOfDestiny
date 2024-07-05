package CardLogic

import (
	"RollsOfDestiny/AccountServer/Database"
	"RollsOfDestiny/AccountServer/Types"
	"os"
	"strconv"
)

func HandleNewCard(userid string) ([]Types.Card, []Types.Card, error) {
	cards, err := Database.GetAllCardsByUserId(userid)
	if err != nil {
		return nil, nil, err
	}
	user, err := Database.GetAccountByUserID(userid)
	if err != nil {
		return nil, nil, err
	}
	numberOfCards, _ := strconv.Atoi(os.Getenv("NUMBER_OF_DIFFERENT_CARDS"))
	var oldCards = make([]Types.Card, numberOfCards)
	var newCards = make([]Types.Card, numberOfCards)
	index := 0
	indexTwo := 0
	for cardIndex := range cards {
		if cards[cardIndex].Count == 0 && cards[cardIndex].Threshold <= user.Rating {
			newCards[index] = cards[cardIndex]
			index++
		} else if cards[cardIndex].Count == 1 {
			oldCards[indexTwo] = cards[cardIndex]
			indexTwo++
		}
	}

	return newCards, oldCards, nil
}
