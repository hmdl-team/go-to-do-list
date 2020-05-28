package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type MsSql struct {
	Host     string
	Port     string
	UserName string
	Password string
	DbName   string
}

//sql server
func (u *MsSql) SqlServeConnect() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%v;database=%s", u.Host, u.UserName, u.Password, u.Port, u.DbName)
	db, err := gorm.Open("mssql", connectionString)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return db, nil
}