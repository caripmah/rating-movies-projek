package controllers

import (
	"movie-api/config"
	"movie-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET all genres
func GetGenres(c *gin.Context) {
	var genres []models.Genre
	config.DB.Find(&genres)
	c.JSON(http.StatusOK, genres)
}

// GET genre by ID
func GetGenre(c *gin.Context) {
	id := c.Param("id")
	var genre models.Genre
	if err := config.DB.First(&genre, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "genre not found"})
		return
	}
	c.JSON(http.StatusOK, genre)
}

// CREATE new genre
func CreateGenre(c *gin.Context) {
	var input models.Genre
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&input)
	c.JSON(http.StatusCreated, input)
}

// UPDATE genre
func UpdateGenre(c *gin.Context) {
	id := c.Param("id")
	var genre models.Genre
	if err := config.DB.First(&genre, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "genre not found"})
		return
	}
	if err := c.ShouldBindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&genre)
	c.JSON(http.StatusOK, genre)
}

// DELETE genre
func DeleteGenre(c *gin.Context) {
	id := c.Param("id")
	var genre models.Genre
	if err := config.DB.First(&genre, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "genre not found"})
		return
	}
	config.DB.Delete(&genre)
	c.JSON(http.StatusOK, gin.H{"message": "genre deleted"})
}
