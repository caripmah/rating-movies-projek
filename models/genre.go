package models

import "time"

type Genre struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"type:varchar(100);unique;not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
