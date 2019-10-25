package database

import (
	"apigo/back/models"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

//database declaration
var db *gorm.DB

//initialisation with credentials, etc..
func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print("godotenv:")
		fmt.Print(e)
	}

	// Get credentials from .env
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbHost := os.Getenv("db_host")
	dbName := os.Getenv("db_name")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) // construct db uri 

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print("gormopen:")
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&models.User{}, &models.Vote{}, &models.Ip{})
}

// GetDB : Getter db
func GetDB() *gorm.DB {
	return db
}
