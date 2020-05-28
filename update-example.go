package main

import (
	"github.com/jinzhu/gorm"
	"log"
	"to-do-list-service/models"
)

func UpdateTest(db *gorm.DB) {

	// ----Tạo mới task------

	task := models.DanhSachCongViec{
		TenCongViec:     "Viêt code todo",
		NoiDungCongViec: "Đang viết chưa xong",
		GhiChu:          "abc",
	}

	err := db.Create(&task).Error

	taskUpdate := models.DanhSachCongViec{
		Id:              2,
		TenCongViec:     "Viêt code todo",
		NoiDungCongViec: "Đang viết chưa xong",
		GhiChu:          "abc",
	}

	// ----cập nhật task------

	// Cập nhật nguyên task
	err = db.Save(&taskUpdate).Error
	if err != nil {
		log.Fatal(err)
	}
	// Code Sql :
	////UPDATE [DanhSachCongViec] SET [TenCongViec] = "Viêt code todods ", [NoiDungCongViec] = 'Đang viết chưa xong', [GhiChu] = 'abc'  WHERE [DanhSachCongViec].[Id] = 2
	//

	// update từng column
	err = db.Model(&models.DanhSachCongViec{}).Where("Id = ? ", 2).Update("GhiChu", "abc").Error
	//
	////UPDATE [DanhSachCongViec] SET  [GhiChu] = 'abc'  WHERE [DanhSachCongViec].[Id] = 2
	//

	// update từng column
	err = db.Model(&models.DanhSachCongViec{}).Where("Id = ? ", 2).UpdateColumn(&models.DanhSachCongViec{GhiChu: "abc1"}).Error
	//
	////UPDATE [DanhSachCongViec] SET  [GhiChu] = 'abc1'  WHERE [DanhSachCongViec].[Id] = 2
	//

	// update all
	err = db.Model(&models.DanhSachCongViec{}).Update("GhiChu", "123").Error
	//
	////UPDATE [DanhSachCongViec] SET  [GhiChu] = 'abc1'  WHERE [DanhSachCongViec].[Id] = 2
	//
}
