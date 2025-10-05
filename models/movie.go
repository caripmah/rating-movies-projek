package models

type Movie struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Title   string `json:"title"`
	Year    int    `json:"year"`
	GenreID uint   `json:"genre_id"`
	Genre   Genre  `gorm:"foreignKey:GenreID" json:"genre"`
}

