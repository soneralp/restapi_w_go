package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating user", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "", "error": err.Error()})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials", "error": err.Error()})
		return
	}

	token, error := utils.GenerateToken(user.Email, user.ID)

	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error generating token", "error": error.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Logged in", "user": user, "token": token})
}
