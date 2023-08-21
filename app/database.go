package app

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	var err error
	dsn := os.Getenv("DB_STRING")
	fmt.Println(dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err !=nil {
		log.Fatal("Failed to Connect Database!")
	}
	fmt.Printf("Database Connection Successfull!");

}