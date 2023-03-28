package routes

import (
	"backEnd_Gin/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(gin_server *gin.Engine) *gin.Engine {
	gin_server.POST("/user_login", controller.UserLogin)
	gin_server.POST("/user_register", controller.UserRegister)
	gin_server.POST("/user_cancel", controller.UserCancel)

	gin_server.GET("/student_list", controller.StudentList)
	gin_server.POST("/student_searchID", controller.StudentIDSearch)
	gin_server.POST("/student_delete", controller.StudentDelete)
	gin_server.POST("/student_add", controller.StudentAdd)
	gin_server.POST("/student_modify", controller.StudentModify)

	gin_server.POST("/next_page", controller.NextPage)
	gin_server.GET("/test", controller.Test)
	return gin_server
}
