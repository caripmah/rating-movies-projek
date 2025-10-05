package controllers

import (
	"movie-api/config"
	"movie-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET all movies
func GetMovies(c *gin.Context) {
	var movies []models.Movie
	config.DB.Preload("Genre").Find(&movies)
	c.JSON(http.StatusOK, movies)
}

// GET movie by ID
func GetMovie(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie
	if err := config.DB.Preload("Genre").First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
		return
	}
	c.JSON(http.StatusOK, movie)
}

// CREATE new movie
func CreateMovie(c *gin.Context) {
	type MovieInput struct {
		Title   string `json:"title"`
		Year    int    `json:"year"`
		GenreID uint   `json:"genre_id"`
	}

	var input MovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie := models.Movie{
		Title:   input.Title,
		Year:    input.Year,
		GenreID: input.GenreID,
	}

	if err := config.DB.Create(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, movie)
}
// UPDATE movie
func UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie
	if err := config.DB.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
		return
	}
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&movie)
	c.JSON(http.StatusOK, movie)
}

// DELETE movie
func DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie
	if err := config.DB.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
		return
	}
	config.DB.Delete(&movie)
	c.JSON(http.StatusOK, gin.H{"message": "movie deleted"})
}
