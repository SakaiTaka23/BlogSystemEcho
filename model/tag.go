package model

type Tag struct{
	ID uint `gorm:"primaryKey"`
	name string `gorm:"unique;not null"`
}