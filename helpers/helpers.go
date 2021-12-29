package helpers

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=user dbname=dbname password=password sslmode=disable")
	HandleErr(err)

	return db
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	HandleErr(err)

	return string(hash)
}
