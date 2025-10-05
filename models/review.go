package models

import "time"

type Review struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`   // foreign key ke users
	MovieID   uint      `json:"movie_id"`  // foreign key ke movies
	Rating    float64       `json:"rating"`    // misalnya 1–5
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relasi
	User  User  `gorm:"foreignKey:UserID"`
	Movie Movie `gorm:"foreignKey:MovieID"`
}
