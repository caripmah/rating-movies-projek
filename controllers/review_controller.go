package controllers

import (
	"movie-api/config"
	"movie-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET all reviews
func GetReviews(c *gin.Context) {
	var reviews []models.Review
	config.DB.Preload("Movie").Find(&reviews)
	c.JSON(http.StatusOK, reviews)
}

// GET review by ID
func GetReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review
	if err := config.DB.Preload("Movie").First(&review, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "review not found"})
		return
	}
	c.JSON(http.StatusOK, review)
}

// CREATE new review
func CreateReview(c *gin.Context) {
	var input models.Review
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&input)
	c.JSON(http.StatusCreated, input)
}

// UPDATE review
func UpdateReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review
	if err := config.DB.First(&review, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "review not found"})
		return
	}
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&review)
	c.JSON(http.StatusOK, review)
}

// DELETE review
func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review
	if err := config.DB.First(&review, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "review not found"})
		return
	}
	config.DB.Delete(&review)
	c.JSON(http.StatusOK, gin.H{"message": "review deleted"})
}
