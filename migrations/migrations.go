package migrations

import (
	"go-fintech-app/helpers"
	hp "go-fintech-app/helpers"
	md "go-fintech-app/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Migrate() {
	db := hp.ConnectDB()
	db.AutoMigrate(&md.User{}, &md.Account{})
	defer db.Close()

	createAccounts()
}

func createAccounts() {
	db := hp.ConnectDB()
	users := [2]md.User{
		{Username: "Martin", Email: "martin@martin.com"},
		{Username: "Jake", Email: "jake@jake.com"},
	}

	for i, u := range users {
		pwd := helpers.HashAndSalt([]byte(u.Username))

		u := md.User{
			Username: u.Username,
			Email:    u.Email,
			Password: pwd,
		}

		db.Create(&u)

		acc := md.Account{
			Type:    "Daily Account",
			Name:    u.Username + "'s Account",
			Balance: uint(10000 * i),
			UserID:  u.ID,
		}
		db.Create(&acc)
	}
	defer db.Close()

}
