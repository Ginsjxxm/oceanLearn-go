package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"wenzhang/model"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	driver := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "root"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		username, password, host, port, database, charset)

	db, err := gorm.Open(driver, args)

	if err != nil {
		panic("failed to connect database：" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	return db
}

func GetDB() *gorm.DB {
	DB = InitDb()
	return DB
}
