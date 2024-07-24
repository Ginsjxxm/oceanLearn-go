package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
	"wenzhang/common"
	"wenzhang/router"
)

func main() {
	InitConfig()
	db := common.GetDB()
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Println("Error closing database connection:", err)
		}
	}()
	r := gin.Default()
	r = router.CollectRouter(r)
	port := viper.GetString("server.port")
	panic(r.Run(":" + port))
}

func InitConfig() {
	workDir, _ := os.Getwd()

	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置读取失败 %s \n", err))
	}

}
