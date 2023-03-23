package routes

import (
	"backEnd_Gin/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(gin_server *gin.Engine) *gin.Engine {
	gin_server.GET("/user_login", controller.User_login)
	return gin_server
}
