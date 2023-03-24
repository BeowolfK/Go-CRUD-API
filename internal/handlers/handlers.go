package handlers

import (
	"go_crud_api/internal/initializer"
	"go_crud_api/internal/models"
	// "log"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func PostsCreate(c *gin.Context) {
	// Extract data of request body

	// Create a new post
	post := models.Post{Title: "My first post", Body: "This is my first post"}
	result := initializer.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": &post,
	})
}