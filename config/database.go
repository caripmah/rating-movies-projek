package config

import (
	"fmt"
	"log"
	"movie-api/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal konek database: %v", err)
	}

	err = database.AutoMigrate(&models.User{}, &models.Genre{}, &models.Movie{}, &models.Review{})
	if err != nil {
		log.Fatalf("Gagal migrate tabel: %v", err)
	}

	DB = database
	log.Println("Database connected 🚀")
}
