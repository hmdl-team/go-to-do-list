package main

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	data2 "to-do-list-service/data"
	"to-do-list-service/db"
	"to-do-list-service/models"
)

// hàm khởi tạo ban đầu
func init() {
	// load biến môi trường vào chương trình
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var (
	Db gorm.DB
)

func main() {
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

	Db = *db

	// Kiểm Tạo mới table
	db.AutoMigrate(
		&models.DanhSachCongViec{},
		&models.ChiTietCongViec{},
	)

	// Echo instance => khởi tạo
	e := echo.New()

	e.GET("/danh-sach-cong-viec/get-all", GetAllHandler)

	log.Fatal(e.Start(":8080"))
}


// Trả reponse cho client
func GetAllHandler(context echo.Context) error {

	data, err := GetDanhSachCongViec()

	// Nếu truy vẫn lỗi trả về thông tin lỗi - status code 500
	if err != nil {
		return context.JSON(http.StatusInternalServerError, &data2.ReposeData{
			Code:    500,
			Data:    nil,
			Message: err.Error(),
		})
	}

	// Trả về dữ liệu cho người dùng
	return context.JSON(http.StatusOK, &data2.ReposeData{
		Code:    200,
		Data:    data,
		Message: "Ok",
	})
}
// Hàm get danh sách công việc
func GetDanhSachCongViec() ([]models.DanhSachCongViec, error) {
	var data []models.DanhSachCongViec

	err := Db.Find(&data).Error
	// => select * from DanhSachCongViec

	if err != nil {
		return nil, err
	}

	return data, nil
}
