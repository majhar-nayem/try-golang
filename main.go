package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/majhar-nayem/testproject/app"
	"github.com/majhar-nayem/testproject/controllers"
)

func init(){
	app.LoadEnv()
	app.ConnectDB()
}

func main(){
	app := gin.Default()
	fmt.Println("runing...")
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Application is Running!",
		})
	})
	app.POST("/post", controllers.PostCreate)
	app.GET("/post", controllers.PostIndex)
	app.GET("/post/:id", controllers.PostShow)
	app.PUT("/post/:id", controllers.PostUpdate)
	app.DELETE("/post/:id", controllers.PostDelete)

	app.Run()
}