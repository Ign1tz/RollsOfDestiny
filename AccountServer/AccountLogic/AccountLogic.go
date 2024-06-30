package AccountLogic

import (
	"RollsOfDestiny/AccountServer/Database"
	"strings"
)

type NewUsernameMessage struct {
	OldUsername string `json:"oldUsername"`
	NewUsername string `json:"newUsername"`
}

type newPasswordMessage struct {
}

func ChangeUsername(message NewUsernameMessage) {
	if len(message.NewUsername) < 3 || len(message.NewUsername) > 20 {
		return
	}
	for _, u := range message.NewUsername {
		if !strings.Contains("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_", string(u)) {
			return
		}
	}

	_, err := Database.GetAccountByUsername(message.OldUsername)

	if err == nil {
		return
	}

	return
}
