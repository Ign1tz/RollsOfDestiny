package SignUpLogic

import (
	"RollsOfDestiny/AccountServer/Database"
	"RollsOfDestiny/AccountServer/Encryption"
)

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginInfoLogic interface {
	CheckPassword()
	CheckUsername()
}

func (l *LoginInfo) CheckPassword() bool {
	account, err := Database.GetAccountByUsername(l.Username)

	if err == nil {
		return Encryption.CheckPasswordHash(l.Password, account.Password)
	} else {
		return false
	}
}

func (l *LoginInfo) CheckUsername() bool {
	_, err := Database.GetAccountByUsername(l.Username)
	return err == nil
}
