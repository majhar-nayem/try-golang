package app

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	var err error
	dsn := "host=rain.db.elephantsql.com user=jkhznaer password=kPJwaAxLj4mscONpWt2ytpt_wd7Vbd61 dbname=jkhznaer port=5432"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err !=nil {
		log.Fatal("Faile to Connect Database!")
	}
	fmt.Printf("Database Connection Successfull!");

}