package main

import (
	"go_crud_api/internal/initializer"
	"go_crud_api/internal/models"
)

func init() {
	initializer.LoadEnvVar()
	initializer.ConnectDB()
}

func main() {
	initializer.DB.AutoMigrate(&models.Post{})
}