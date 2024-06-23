package SignUpLogic

import (
	"RollsOfDestiny/AccountServer/Database"
	"strings"
)

type SignUpInfo struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type SignUpInfoLogic interface {
	CheckUsername()
	CheckEmail()
	ComparePassword()
}

func (s *SignUpInfo) CheckUsername() bool {
	for _, u := range s.Username {
		if !strings.Contains("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_", string(u)) {
			return false
		}
	}

	_, err := Database.GetAccountByUsername(s.Username)

	if err == nil {
		return false
	}

	return true
}

func (s *SignUpInfo) CheckEmail() bool {

	_, err := Database.GetAccountByEmail(s.Email)

	if err == nil {
		return false
	}
	return true
}

func (s *SignUpInfo) ComparePassword() bool {
	correct := 6 <= len(s.Password) && len(s.Password) <= 40

	for _, u := range s.Password {
		if !strings.Contains("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_!$@&#+*-€", string(u)) {
			return false
		}
	}

	correct = correct && s.Password == s.ConfirmPassword

	return correct
}
