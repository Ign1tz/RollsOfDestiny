package Types

type NewDeckMessage struct {
	Name string
}

type AddCard struct {
	Name   string
	Deckid string
}
