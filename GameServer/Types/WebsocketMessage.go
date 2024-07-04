package Types

type WebsocketMessage struct {
	Type        string `json:"type"`
	MessageBody string `json:"messageBody"`
	GameId      string `json:"gameId"`
	Userid      string `json:"userid"`
}
