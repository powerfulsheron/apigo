package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Token JWT claims struct
type Token struct {
	Uuid        *uuid.UUID
	AccessLevel int
	jwt.StandardClaims
}

//User a struct to rep user
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
