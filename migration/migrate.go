package main

import (
	"github.com/majhar-nayem/testproject/app"
	"github.com/majhar-nayem/testproject/models"
)

func init(){
	app.LoadEnv()
	app.ConnectDB()
}

func main(){
	app.DB.AutoMigrate(&models.Post{})
}