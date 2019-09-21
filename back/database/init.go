package database

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"os"
	"github.com/joho/godotenv"
	"fmt"
	"apigo/back/models"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print("godotenv:")
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbHost := os.Getenv("db_host")
	dbName := os.Getenv("db_name")


	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println("dburi")
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print("gormopen:")
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&models.Account{})
}

func GetDB() *gorm.DB {
	return db
}