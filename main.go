package main

import (
	"log"
	"os"

	"movie-api/config"
	"movie-api/routes"

	"github.com/joho/godotenv"
)

func main() {
	// load .env (optional untuk lokal)
	_ = godotenv.Load()

	// connect DB
	config.ConnectDatabase()

	// setup router
	r := routes.SetupRouter()

	// Railway pakai PORT dari environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default untuk lokal
	}

	log.Printf("Server running on port %s", port)
	r.Run("0.0.0.0:" + port) // ✅ penting: harus 0.0.0.0, bukan localhost
}
