package routes

import (
	"backEnd_Gin/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(gin_server *gin.Engine) *gin.Engine {
	gin_server.POST("/user_login", controller.UserLogin)
	gin_server.POST("/user_register", controller.UserRegister)
	gin_server.POST("/user_cancel", controller.UserCancel)

	gin_server.GET("/student_show", controller.StudentShow)
	gin_server.POST("/student_show/studentID_search", controller.StudentIDSearch)

	gin_server.GET("/test", controller.Test)
	return gin_server
}
