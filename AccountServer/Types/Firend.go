package Types

type Friend struct {
	UserID string
	Friend Account
}

type FriendInfo struct {
	FriendUsername string `json:"username"`
}
