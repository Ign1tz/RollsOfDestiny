package Types

type Card struct {
	CardID  string
	Name    string
	Cost    int
	Effect  string
	Picture string
	DeckID  string
	Played  bool
	InHand  bool
}
