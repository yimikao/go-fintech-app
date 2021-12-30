package helpers

import (
	md "go-fintech-app/models"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func VerifyPwd(pwd string, pass string) map[string]interface{} {
	err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(pwd))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return map[string]interface{}{"message": "Wrong password"}
	}
	return nil
}

func SignToken(claim uint) (token string, err error) {
	tokenContent := jwt.MapClaims{
		"user_id": claim,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err = jwtToken.SignedString([]byte("Token"))

	return
}

func ValidateReq(v []md.Validation) bool {
	usernameVdn := regexp.MustCompile("^([A-Za-z0-9]{5,})+$")
	emailVdn := regexp.MustCompile("^[A-Za-z0-9]+[@]+[A-Za-z0-9]+[.]+[A-Za-z]+$")

	for _, vdn := range v {
		switch vdn.Valid {
		case "username":
			if !usernameVdn.MatchString(vdn.Value) {
				return false
			}
		case "email":
			if !emailVdn.MatchString(vdn.Value) {
				return false
			}
		case "password":
			if len(vdn.Value) < 5 {
				return false
			}
		}
	}
	return true
}
