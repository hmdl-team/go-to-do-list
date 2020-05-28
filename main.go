package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"to-do-list-service/db"
	"to-do-list-service/models"
)

func main() {
	// load biến môi trường vào chương trình
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Khai báo lớp kết nối
	dbConfig := db.MsSql{
		Host:     os.Getenv("SQL_HOST"),
		Port:     os.Getenv("SQL_PORT"),
		UserName: os.Getenv("SQL_USER"),
		Password: os.Getenv("SQL_PASSWORD"),
		DbName:   os.Getenv("SQL_DBNAME"),
	}

	db, err := dbConfig.SqlServeConnect()

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.DanhSachCongViec{})

}
