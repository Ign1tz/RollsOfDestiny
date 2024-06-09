package Database

import "RollsOfDestiny/GameServer/Types"

func InsertPlayer(player Types.Player) error {
	_, err := Database.Exec("INSERT INTO players (userid, username, mana) VALUES ($1, $2, $3)",
		player.UserID, player.Username, player.Mana)
	return err
}

func InsertDeck(deck Types.Deck) error {
	_, err := Database.Exec("INSERT INTO decks VALUES ($1, $2)", deck.DeckID, deck.UserID)
	return err
}

func InsertCard(card Types.Card) error {
	_, err := Database.Exec(
		"INSERT INTO cards VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		card.Name, card.Cost, card.Effect, card.Picture, card.CardID, card.DeckID, card.Played, card.InHand)
	return err
}

func InsertColumn(column Types.Column) error {
	_, err := Database.Exec("INSERT INTO columns VALUES ($1, $2, $3, $4, $5)",
		column.GridId, column.First, column.Second, column.Third, column.Placement)
	return err
}

func InsertGrid(grid Types.Grid) error {
	_, err := Database.Exec("INSERT INTO grids VALUES ($1)",
		grid.GridId)
	return err
}

func InsertGame(game Types.Game) error {
	_, err := Database.Exec("INSERT INTO games VALUES ($1, $2, $3, $4, $5, $6)",
		game.GameID, game.HostId, game.Guest, game.ActivePlayer,
		game.HostGrid, game.GuestGrid)
	return err
}

func UndefinedDelete(table string, key string, value string) error {
	_, err := Database.Exec("Delete From $1 Where $2 = $3", table, key, value)
	return err
}
