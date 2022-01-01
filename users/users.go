package users

import (
	hp "go-fintech-app/helpers"
	md "go-fintech-app/models"
)

type RegisterParams struct {
	Username string
	Email    string
	Password string
}

type LoginParams struct {
	Username string
	Password string
}

func Register(p *RegisterParams) map[string]interface{} {
	db := hp.ConnectDB()
	defer db.Close()

	hashedPwd := hp.HashAndSalt([]byte(p.Password))

	u := &md.User{
		Username: p.Username,
		Email:    p.Email,
		Password: hashedPwd,
	}
	db.Create(&u)

	a := md.Account{
		Type:    "savings",
		Name:    u.Username + "'s account.",
		Balance: 0,
		UserID:  u.ID,
	}
	db.Create(&a)

	accs := []md.AccountResponse{
		{ID: a.ID, Name: a.Name, Balance: a.Balance},
	}

	uRes := md.UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Accounts: accs,
	}

	t, err := hp.SignToken(u.ID)
	hp.HandleErr(err)

	res := map[string]interface{}{
		"message": "registration successful",
		"token":   t,
		"data":    uRes,
	}
	return res
}

func Login(p *LoginParams) map[string]interface{} {
	db := hp.ConnectDB()
	defer db.Close()

	u := &md.User{}
	if db.Where("username = ? ", p.Username).First(&u).RecordNotFound() {
		//user already filles at .First(&u)
		return map[string]interface{}{"message": "User not found"}
	}

	_ = hp.VerifyPwd(p.Password, u.Password)

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
		"message": "login successful",
		"token":   t,
		"data":    uRes,
	}
	return res
}

/*
If we would like to do any activity related to the user, the application needs to know if we are the owner of the account.
We can verify that by validating the JWT token.
*/
func GetUser(id string, jwt string) map[string]interface{} {
	isValid := hp.ValidateToken(id, jwt)
	if !isValid {
		return map[string]interface{}{"message": "invalid token"}
	}
	db := hp.ConnectDB()
	defer db.Close()

	u := &md.User{}

	if db.Where("id = ? ", id).First(&u).RecordNotFound() {
		return map[string]interface{}{"message": "User not found"}
	}
	accs := []md.AccountResponse{}
	db.Table("accounts").Select("id, name, balance").Where("user_id =  ?", u.ID).Scan(&accs)

	uRes := md.UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Accounts: accs,
	}

	res := map[string]interface{}{
		"message": "user fetched successfully",
		"data":    uRes,
	}

	return res
}
