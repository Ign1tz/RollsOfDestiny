package Types

type NewDeckMessage struct {
	Name string `json:"name"`
}

type AddCard struct {
	Name   string `json:"name"`
	Deckid string `json:"deckid"`
}

type AknowledgeCard struct {
	Name string `json:"name"`
}
