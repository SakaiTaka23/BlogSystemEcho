package model

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	if db != nil {
		return db
	}
	db,err := newDB()
	if err != nil{
		panic(err.Error())
	}
	return db
}

func newDB() (*gorm.DB, error) {
	_ = godotenv.Load("../.env")

	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := os.Getenv("DB")
	dsn := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db, err
}
