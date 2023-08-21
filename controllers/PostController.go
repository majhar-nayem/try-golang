package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/majhar-nayem/testproject/app"
	"github.com/majhar-nayem/testproject/models"
)

func PostCreate(ctx *gin.Context){
	var body struct {
		Title string
		Description string
	}
	ctx.Bind(&body)
	post := models.Post{Title: body.Title, Description: body.Description}
	result := app.DB.Create(&post)

	if result.Error != nil {
		ctx.Status(400)
		return;
	}
	ctx.JSON(200, gin.H{
		"post" : post,
	})
}

func PostIndex(ctx *gin.Context){
	var posts []models.Post
	app.DB.Find(&posts)

	ctx.JSON(200, gin.H{
		"posts" : posts,
	})
}

func PostShow(ctx *gin.Context){
	id := ctx.Param("id")

	var post models.Post
	app.DB.First(&post, id)

	ctx.JSON(200, gin.H{
		"post" : post,
	})

}

func PostUpdate(ctx *gin.Context){

	id := ctx.Param("id")
	
	var body struct {
		Title string
		Description string
	}
	ctx.Bind(&body)
	
	var post models.Post
	app.DB.First(&post, id)

	app.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Description: body.Description,
	})

	ctx.JSON(200, gin.H{
		"post" : post,
	})
}

func PostDelete(ctx *gin.Context){
	id := ctx.Param("id")

	var post models.Post
	app.DB.Delete(&post,id)

	ctx.JSON(200, gin.H{
		"message" : "Post Deleted!",
	})
}