package models

type DanhSachCongViec struct {
	Id              int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT; Column:Id" json:"id"`
	TenCongViec     string `gorm:"type:nvarchar(150);Column:TenCongViec" json:"ten_cong_viec"`
	NoiDungCongViec string `gorm:"type:nvarchar(200);Column:NoiDungCongViec" json:"noi_dung_cong_viec"`
	GhiChu          string `gorm:"type:nvarchar(200);Column:GhiChu" json:"ghi_chu"`
}

func (DanhSachCongViec) TableName() string {
	return "DanhSachCongViec"
}
