package helpers

import "golang.org/x/crypto/bcrypt"

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	HandleErr(err)

	return string(hash)
}
