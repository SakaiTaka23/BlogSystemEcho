package model

import "time"

type Article struct {
	ID        uint   `gorm:"primaryKey"`
	title     string `gorm:"unique;not null"`
	body      string `gorm:"not null"`
	CreatedAt time.Time
}


