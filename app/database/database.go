package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var (
	DBConn *gorm.DB
)

func ConnectDb() {
	// Update these values with your local MySQL credentials
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC", username, password, host, port, dbName)

	db, err := gorm.Open("mysql", dbConnectionString)
	if err != nil {
		panic(err)
	}

	log.Println("Connected to local MySQL")

	DBConn = db
}
