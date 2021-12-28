package migrations

import (
	"go-fintech-app/helpers"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=user dbname=dbname password=password sslmode=disable")
	helpers.HandleErr(err)

	return db
}

func Migrate() {
	db := connectDB()
	db.AutoMigrate(&User{}, &Account{})
	defer db.Close()

	createAccounts()
}

func createAccounts() {
	db := connectDB()
	users := [2]User{
		{Username: "Martin", Email: "martin@martin.com"},
		{Username: "Jake", Email: "jake@jake.com"},
	}

	for i, u := range users {
		pwd := helpers.HashAndSalt([]byte(u.Username))

		u := User{
			Username: u.Username,
			Email:    u.Email,
			Password: pwd,
		}

		db.Create(&u)

		acc := Account{
			Type:    "Daily Account",
			Name:    u.Username + "'s Account",
			Balance: uint(10000 * i),
			UserID:  u.ID,
		}
		db.Create(&acc)
	}
	defer db.Close()

}
