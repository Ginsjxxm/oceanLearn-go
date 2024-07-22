package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"wenzhang/common"
	"wenzhang/router"
)

func main() {
	db := common.GetDB()
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Println("Error closing database connection:", err)
		}
	}()
	r := gin.Default()
	r = router.CollectRouter(r)
	panic(r.Run(":8080"))
}
