package models

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
	"time"
)

/*
JWT claims struct
*/
type Token struct {
	Uuid *uuid.UUID 
	AccessLevel int
	jwt.StandardClaims
}

//a struct to rep user
type User struct {
	gorm.Model
	Uuid
	Email       string    `json:"email"`
	Password    string    `json:"pass"`
	AccessLevel int       `json:"access_level"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"` 
	DateOfBirth time.Time `json:"birth_date"`
}