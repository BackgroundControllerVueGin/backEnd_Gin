package routes

import (
	"backEnd_Gin/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(gin_server *gin.Engine) *gin.Engine {
	gin_server.POST("/user_login", controller.User_login)
	gin_server.GET("/test", controller.Test)
	return gin_server
}
