package controllers

import (
	"movie-api/config"
	"movie-api/models"
	"movie-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register
func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := utils.HashPassword(input.Password)
	user := models.User{Username: input.Username, Email: input.Email, Password: hashedPassword}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "register success"})
}

// Login
func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	config.DB.Where("email = ?", input.Email).First(&user)

	if user.ID == 0 || !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, _ := utils.GenerateJWT(user.ID, user.Email)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
