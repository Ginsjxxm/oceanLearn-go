package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"wenzhang/model"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	driver := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.username")
	charset := viper.GetString("datasource.charset")
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
