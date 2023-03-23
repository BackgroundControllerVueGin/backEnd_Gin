package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func User_login(ctx *gin.Context) {
	fmt.Println("User_login is running")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}
