package DeckLogic

import (
	"RollsOfDestiny/AccountServer/Database"
	"RollsOfDestiny/AccountServer/Types"
	"fmt"
	"github.com/google/uuid"
	"log"
	"strconv"
)

func GetCardsOfDeckAsJsonString(deckid string) string {

	cards, err := Database.GetCardsByDeckID(deckid)

	if err != nil {
		log.Println(err)
		return ""
	}
	var cardString = ""
	for cardIndex := range cards {
		if cards[cardIndex].UserID != "" {
			cardString = fmt.Sprintf(`%s, {"name": "%s", "count": "%s", "image": "%s"}`, cardString, cards[cardIndex].Name, strconv.Itoa(cards[cardIndex].Count), cards[cardIndex].Image)
		}
	}
	if cardString != "" {
		cardString = cardString[2:]
	}
	return cardString
}

func CreateNewDeck(name string, userid string) {
	newDeck := Types.Deck{
		UserID: userid,
		DeckID: uuid.New().String(),
		Name:   name,
		Active: false,
	}
	err := Database.InsertDeck(newDeck)
	if err != nil {
		return
	}
}

func AddCardToDeck(cardInfos Types.AddCard, userid string) {
	err := Database.UpdateCardDeckId(cardInfos, userid)
	if err != nil {
		return
	}

}

func RemoveCardFromDeck(cardInfos Types.AddCard, userid string) {
	err := Database.RemoveCardDeckId(userid, cardInfos)
	if err != nil {
		return
	}
}

func ChangeActiveDeck(deckInfo Types.AddCard, userid string) {
	err := Database.ChangeActiveDeck(deckInfo.Deckid, userid)
	if err != nil {
		return
	}
}

func RemoveDeck(deckid string, userid string) {
	err := Database.RemoveCardDeckIdByDeckId(userid, deckid)
	if err != nil {
		log.Println(err)
		return
	}
}
