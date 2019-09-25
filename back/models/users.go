package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"time"
)

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

//a struct to rep user account
type User struct {
	gorm.Model
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"token";sql:"-"`
	FirstName string
	LastName string
	DateOfBirth time.Time
}