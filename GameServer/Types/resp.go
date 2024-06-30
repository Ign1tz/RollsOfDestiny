package Types

type Resp struct {
	Gameid    string `json:"gameid"`
	ColumnKey string `json:"columnKey"`
}

type BotResp struct {
	Userid                string `json:"Userid"`
	WebsocketConnectionId string `json:"websocketConnectionId"`
}
