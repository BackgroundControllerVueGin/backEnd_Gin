package main

import (
	"Gin/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(gin_server *gin.Engine) *gin.Engine {
	gin_server.POST("/user_login", controller.User_login)
	return gin_server
}
