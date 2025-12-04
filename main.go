package main

import (
	"go-blog/config"
	"go-blog/models"
	"go-blog/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect DB
	config.ConnectDB()

	// Auto migrate model
	config.DB.AutoMigrate(&models.Blog{})

	// Blog Routes
	routes.BlogRoutes(r)

	// Start server
	r.Run(":8080")
}
