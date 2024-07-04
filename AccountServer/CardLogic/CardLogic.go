package CardLogic

import (
	"RollsOfDestiny/AccountServer/Database"
	"RollsOfDestiny/AccountServer/Types"
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
	var newCards = make([]Types.Card, 100)
	var oldCards = make([]Types.Card, 100)
	index := 0
	indexTwo := 0
	for cardIndex := range cards {
		if cards[cardIndex].Count == 0 && cards[cardIndex].Threshold <= user.Rating {
			newCards[index] = cards[cardIndex]
			index++
		} else {
			oldCards[indexTwo] = cards[cardIndex]
			indexTwo++
		}
	}

	return newCards, oldCards, nil
}
