package utils

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Err_return_json(ctx *gin.Context, result *gorm.DB, message string) bool {
	if result.Error != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 405,
			"msg":  message,
			"err":  result.Error.Error(),
		})
		return false
	} else {
		return true
	}
}
