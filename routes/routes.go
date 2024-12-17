package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authanticate)

	authenticated.POST("/events", createEvent)
	authenticated.POST("/events/delete", deleteAllEvents)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
}
