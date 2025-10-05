package main

import (
	"log"
	"os"

	"movie-api/config"
	"movie-api/routes"

	"github.com/joho/godotenv"
)


func main() {
    // load env
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found or could not load it. Make sure env vars are set.")
    }

    // debug: cek env
    log.Println("DB_USER:", os.Getenv("DB_USER"))
    log.Println("DB_PASSWORD:", os.Getenv("DB_PASSWORD"))
    log.Println("DB_HOST:", os.Getenv("DB_HOST"))
    log.Println("DB_PORT:", os.Getenv("DB_PORT"))
    log.Println("DB_NAME:", os.Getenv("DB_NAME"))

    // connect db
    config.ConnectDatabase()

    r := routes.SetupRouter()
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    r.Run(":" + port)
}