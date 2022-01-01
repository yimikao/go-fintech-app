package models

import (
	"github.com/jinzhu/gorm"
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

type UserResponse struct {
	ID       uint
	Username string
	Email    string
	Accounts []AccountResponse
}

type AccountResponse struct {
	ID      uint
	Name    string
	Balance uint
}

type Validation struct {
	Value string
	Valid string
}

type ErrResponse struct {
	Message string
}
