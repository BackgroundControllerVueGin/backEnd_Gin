package controller

import (
	"backEnd_Gin/common"
	"backEnd_Gin/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

var db = common.GetDB()

func StudentTest(ctx *gin.Context) {
	var Student []model.Student
	result := db.Find(&Student)
	if result.Error != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 405,
			"msg":  result.Error,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": Student,
		})
	}
}
