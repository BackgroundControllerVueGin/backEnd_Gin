package controller

import (
	"backEnd_Gin/model"
	"backEnd_Gin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserLogin(ctx *gin.Context) {
	fmt.Println("User_login is running yes")
	var jsoninfo model.User
	if err := ctx.ShouldBindJSON(&jsoninfo); err != nil {
		fmt.Printf("%s and %s\n", jsoninfo.Username, jsoninfo.Password)
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 405,
			"msg":  "error",
		})
		return
	}
	var user model.User
	result := db.Where("username = ? and password = ?", jsoninfo.Username, jsoninfo.Password).First(&user)
	if result.Error != nil {
		if result.RowsAffected != 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 405,
				"msg":  result.Error,
			})
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 405,
				"msg":  "登录失败，找不到该账号",
				"data": result.RowsAffected,
			})
		}
	} else {
		var usertoken model.UserToken
		usertoken.No = strconv.Itoa(user.No)
		usertoken.Username = user.Username
		usertoken.Password = user.Password
		token := utils.GenerateToken(&usertoken)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "登录成功",
			"data": []interface{}{
				user,
			},
			"token": token,
		})

		fmt.Println(jsoninfo)
	}

}

func UserRegister(ctx *gin.Context) {
	fmt.Println("User_register is running")
	json := make(map[string]interface{})
	ctx.BindJSON(&json)
	fmt.Println(json["username"])
	if json["password"] == json["passwordAgain"] {
		username := json["username"]
		password := json["password"]
		user := model.User{
			Username: username.(string),
			Password: password.(string),
		}
		result := db.Create(&user)
		if utils.Err_return_json(ctx, result, "注册失败，写入数据库失败，可能是存在同名账户") {
			result := db.Where("username = ? and password = ?", username, password).First(&user)
			if utils.Err_return_json(ctx, result, "注册失败，查询不到注册账号") {
				ctx.JSON(http.StatusOK, gin.H{
					"code":    200,
					"message": "注册成功",
					"data": []interface{}{
						user,
					},
				})
			}
		} else {
			fmt.Println("ERROR/User_register is err: Data conflict")
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 403,
			"msg":  "注册失败，请检查密码是否相同",
			"data": json,
		})
	}
}

func UserCancel(ctx *gin.Context) {
	fmt.Println("User_cancel is running")
	var jsoninfo model.User
	var user model.User
	if err := ctx.ShouldBindJSON(&jsoninfo); err == nil {
		result := db.Where("username = ? and password = ?", jsoninfo.Username, jsoninfo.Password).Delete(&user)
		if utils.Err_return_json(ctx, result, "数据发送错误：找不到该用户") {
			if result.RowsAffected == 0 {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 200,
					"msg":  "注销失败，找不到该账号",
					"data": []interface{}{
						gin.H{
							"username": jsoninfo.Username,
							"password": jsoninfo.Password,
						},
					},
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code":   200,
					"msg":    "注销成功",
					"number": result.RowsAffected,
					"data": []interface{}{
						gin.H{
							"username": jsoninfo.Username,
							"password": jsoninfo.Password,
						},
					},
				})
			}
		}
	} else {
		fmt.Println(err)
		fmt.Println(jsoninfo)
	}
}

//
//func Test(ctx *gin.Context) {
//	ctx.JSON(http.StatusOK, gin.H{
//		"msg": "测试成功",
//	})
//}
