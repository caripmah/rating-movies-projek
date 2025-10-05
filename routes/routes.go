package routes

import (
	"movie-api/controllers"
	"movie-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// ==========================
	// 🔐 AUTH ROUTES
	// ==========================
	r.POST("/api/users/register", controllers.Register)
	r.POST("/api/users/login", controllers.Login)

	// ==========================
	// 🎬 MOVIE ROUTES (protected by JWT)
	// ==========================
	movie := r.Group("/api/movies")
	movie.Use(middlewares.JWTAuthMiddleware())
	{
		movie.GET("/", controllers.GetMovies)
		movie.GET("/:id", controllers.GetMovie)
		movie.POST("/", controllers.CreateMovie)
		movie.PUT("/:id", controllers.UpdateMovie)
		movie.DELETE("/:id", controllers.DeleteMovie)
	}

	// ==========================
	// 🎭 GENRE ROUTES (protected by JWT)
	// ==========================
	genre := r.Group("/api/genres")
	genre.Use(middlewares.JWTAuthMiddleware())
	{
		genre.GET("/", controllers.GetGenres)
		genre.GET("/:id", controllers.GetGenre)
		genre.POST("/", controllers.CreateGenre)
		genre.PUT("/:id", controllers.UpdateGenre)
		genre.DELETE("/:id", controllers.DeleteGenre)
	}

	// ==========================
	// 📝 REVIEW ROUTES (protected by JWT)
	// ==========================
	review := r.Group("/api/reviews")
	review.Use(middlewares.JWTAuthMiddleware())
	{
		review.GET("/", controllers.GetReviews)
		review.GET("/:id", controllers.GetReview)
		review.POST("/", controllers.CreateReview)
		review.PUT("/:id", controllers.UpdateReview)
		review.DELETE("/:id", controllers.DeleteReview)
	}

	return r
}
