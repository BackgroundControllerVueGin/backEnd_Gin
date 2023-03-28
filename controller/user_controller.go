package controller

import (
	"backEnd_Gin/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(ctx *gin.Context) {
	fmt.Println("User_login is running yes")

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
		fmt.Println("ERROR/User_login is err: No data")
		ctx.JSON(http.StatusOK, gin.H{
			"code":    403,
			"message": "账号密码错误",
		})
	}
	var user model.User
	row.Scan(&user.No, &user.Username, &user.Password)
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

func UserRegister(ctx *gin.Context) {
	fmt.Println("User_register is running")
	json := make(map[string]interface{})
	ctx.BindJSON(&json)
	fmt.Println(json["username"])
	if json["password"] == json["passwordAgain"] {
		sqlStr := "insert into user(username,password) VALUES (?,?)"
		_, err := db.Exec(sqlStr, json["username"], json["password"])
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 405,
				"msg":  "注册失败，写入数据库失败",
				"err":  err.Error(),
			})
			fmt.Println("ERROR/User_register is err: Data conflict")
			//log.Fatalln(err)
			return
		} else {
			sqlStr = "select * from user where username = ? and password = ?;"
			row := db.QueryRow(sqlStr, json["username"], json["password"])
			var user model.User
			row.Scan(&user.No, &user.Username, &user.Password)
			ctx.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "注册成功",
				"data": user,
			})
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
	if err := ctx.ShouldBindJSON(&jsoninfo); err == nil {
		rs, err := db.Exec("DELETE FROM user WHERE username = ? and password = ?;", jsoninfo.Username, jsoninfo.Password)
		if err != nil {
			fmt.Println("ERROR/User_cancel is err: No data")
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 405,
				"msg":  "数据发送错误：找不到该用户",
				"err":  err.Error(),
			})
			return
		} else {
			returnNumber, err := rs.RowsAffected()
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 403,
					"msg":  "注销失败",
					"err":  err.Error(),
					"data": []interface{}{
						jsoninfo,
					},
				})
			}
			if returnNumber == 0 {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 403,
					"msg":  "注销失败,找不到该账号或密码错误",
					"data": []interface{}{
						jsoninfo,
					},
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code":   200,
					"msg":    "注销成功",
					"number": returnNumber,
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

func Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "测试成功",
	})
}
