package SignUpLogic

func LoginToAccount(info LoginInfo) bool {
	accountExists := info.CheckUsername()

	if accountExists {
		return info.CheckPassword()
	} else {
		return false
	}
}
