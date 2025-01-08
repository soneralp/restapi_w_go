package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/middlewares"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// DB'yi başlat
	db.InitDB()

	// Gin sunucusunu başlat
	server := gin.Default()

	// Metrics middleware'ini ekle
	server.Use(middlewares.MetricsMiddleware())

	// Rotaları kaydet
	routes.RegisterRoutes(server)

	// Sunucuyu başlat
	server.Run(":8080")
}
