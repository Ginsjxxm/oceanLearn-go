package router

import (
	"github.com/gin-gonic/gin"
	"wenzhang/controller"
	"wenzhang/middleware"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleWare(), controller.Info)
	return r
}
