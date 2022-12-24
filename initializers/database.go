package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")
	DBPORT := os.Getenv("DB_PORT")

	URL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", HOST, USER, PASS, DBNAME, DBPORT)
	fmt.Println(URL)
	var err error

	DB, err = gorm.Open(postgres.Open(URL), &gorm.Config{})
	if err != nil {
		log.Fatal("Can't connect to database")
	}
}
