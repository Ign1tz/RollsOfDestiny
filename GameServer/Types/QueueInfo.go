package Types

type QueueInfo struct {
	UserId                string `json:"userid"`
	WebsocketConnectionId string `json:"websocketconnectionid"`
	Username              string `json:"username"`
}

type QueueInfoFriend struct {
	UserId                string `json:"userid"`
	WebsocketConnectionId string `json:"websocketconnectionid"`
	Username              string `json:"username"`
	FriendId              string `json:"friendid"`
}
