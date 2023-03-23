package controller

import (
	"backEnd_Gin/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func User_login(ctx *gin.Context) {
	fmt.Println("User_login is running")

	//json := make(map[string]interface{})
	//ctx.BindJSON(&json)
	//fmt.Println(json)

	var jsoninfo model.User
	if err := ctx.ShouldBindJSON(&jsoninfo); err != nil {
		fmt.Printf("%s and %s\n", jsoninfo.Username, jsoninfo.Password)
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 405,
			"msg":  "error",
		})
		return
	}
	sqlStr := "select * from user where username = ? and password = ?"
	row := db.QueryRow(sqlStr, jsoninfo.Username, jsoninfo.Password)
	if row == nil {
		fmt.Println("User_login is err: %s")
		log.Fatalln("User_login is err: %s")
	}
	var user model.User
	row.Scan(&user.Username, &user.Password)
	if user.Username != jsoninfo.Username && user.Password != jsoninfo.Password {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    403,
			"message": "账号密码错误",
			"data": []interface{}{
				jsoninfo,
			},
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "登录成功",
			"data": []interface{}{
				user,
			},
		})
	}
	fmt.Println(jsoninfo)
}

func Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "测试成功",
	})
}
