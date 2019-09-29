package data

import (
	"apigo/back/database"
	"apigo/back/models"
	u "apigo/back/utils"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	"golang.org/x/crypto/bcrypt"
)

// User : user model
type User models.User
type Token models.Token

// Validate incoming user details
func (user *User) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(user.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(user.Password) < 6 {
		return u.Message(false, "Password is too short"), false
	}

	//Email must be unique
	temp := &models.User{}

	//check for errors and duplicate emails
	err := database.GetDB().Table("users").Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

// Create an user
func (user *User) Create() map[string]interface{} {

	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	database.GetDB().Create(user)

	if user.ID <= 0 {
		return u.Message(false, "Failed to create user, connection error.")
	}

	user.Password = "" //delete password

	response := u.Message(true, "User has been created")
	response["user"] = user
	return response
}

// Update an user
func (user *User) Update() map[string]interface{} {

	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	database.GetDB().Save(user)

	if user.ID <= 0 {
		return u.Message(false, "Failed to create user, connection error.")
	}

	user.Password = "" //delete password

	response := u.Message(true, "User has been Updated")
	response["user"] = user
	return response
}

// Delete the user
func (user *User) Delete() map[string]interface{} {

	if user.ID <= 0 {
		return u.Message(false, "Failed to delete user, connection error.")
	}
	database.GetDB().Delete(user)

	response := u.Message(true, "User has been Deleted")
	response["user"] = user
	return response
}

// Login user
func Login(email, password, adress string) map[string]interface{} {

	user := &models.User{}
	err := database.GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var ip Ip
		ip.Adress = adress
		return ip.Increment()
	}
	//Worked! Logged In

	//Create new JWT token for the newly registered user
	tk := &models.Token{Uuid: user.Uuid.UUID, AccessLevel: user.AccessLevel}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	response := u.Message(true, "Logged In")
	response["jwt"] = tokenString
	return response

}

// GetUser getter
func GetUser(uuid uuid.UUID) *models.User {

	user := &models.User{}
	database.GetDB().Table("users").Where("uuid = ?", uuid).First(user)
	if user.Email == "" { //User not found!
		return nil
	}
	user.Password = ""
	return user
}

func GetUserWithPW(uuid uuid.UUID) *models.User {

	user := &models.User{}
	database.GetDB().Table("users").Where("uuid = ?", uuid).First(user)
	if user.Email == "" { //User not found!
		return nil
	}
	return user
}
