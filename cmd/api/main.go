package main

import (
	"go_crud_api/internal/handlers"
	"go_crud_api/internal/initializer"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVar()
	initializer.ConnectDB()
}



func main() {
	r := gin.Default()
	r.GET("/ping", handlers.Ping)
	r.POST("/posts", handlers.PostsCreate)
	r.Run()
}