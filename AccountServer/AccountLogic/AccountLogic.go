package AccountLogic

import (
	"RollsOfDestiny/AccountServer/Database"
	"RollsOfDestiny/AccountServer/Encryption"
	"log"
	"strings"
)

type NewUsernameMessage struct {
	OldUsername string `json:"oldUsername"`
	NewUsername string `json:"newUsername"`
}

type NewPasswordMessage struct {
	OldPassword        string `json:"oldPassword"`
	NewPassword        string `json:"newPassword"`
	ConfirmNewPassword string `json:"confirmNewPassword"`
}

type NewProfilePicture struct {
	ProfilePicture string `json:"profilePicture"`
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

	err := Database.UpdateUsername(message.OldUsername, message.NewUsername)

	if err != nil {
		log.Println(err)
		return
	}
}

func ChangePasswprd(message NewPasswordMessage, userid string) {

	account, err := Database.GetAccountByUserID(userid)

	if err == nil && Encryption.CheckPasswordHash(message.OldPassword, account.Password) {
	} else {
		return
	}

	correct := 6 <= len(message.NewPassword) && len(message.NewPassword) <= 40

	for _, u := range message.NewPassword {
		if !strings.Contains("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_!$@&#+*-€", string(u)) {
			return
		}
	}

	correct = correct && message.NewPassword == message.ConfirmNewPassword
	hashedPassword, _ := Encryption.HashPassword(message.NewPassword)
	err = Database.UpdatePassword(userid, hashedPassword)

	if err != nil {
		log.Println(err)
		return
	}
}
