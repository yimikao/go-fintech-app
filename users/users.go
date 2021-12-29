package users

import (
	hp "go-fintech-app/helpers"
	md "go-fintech-app/models"
)

func Login(username string, pwd string) map[string]interface{} {
	db := hp.ConnectDB()
	defer db.Close()

	u := &md.User{}
	if db.Where("username = ? ", username).First(&u).RecordNotFound() {
		//user already filles at .First(&u)
		return map[string]interface{}{"message": "User not found"}
	}

	_ = hp.VerifyPwd(pwd, u.Password)

	accs := []md.AccountResponse{}
	db.Table("accounts").Select("id, name, balance").Where("user_id = ?", u.ID).Scan(&accs)

	uRes := md.UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Accounts: accs,
	}

	t, err := hp.SignToken(u.ID)
	hp.HandleErr(err)

	res := map[string]interface{}{
		"message": "user loggedin",
		"token":   t,
		"data":    uRes,
	}
	return res
}
