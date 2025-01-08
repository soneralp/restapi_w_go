package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/middlewares"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.Use(middlewares.MetricsMiddleware())

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
