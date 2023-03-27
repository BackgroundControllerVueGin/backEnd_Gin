package controller

import (
	"backEnd_Gin/model"
	"backEnd_Gin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StudentList(ctx *gin.Context) {
	fmt.Println("Student_show is running")
	rows, err := db.Query("SELECT * FROM student")
	if err != nil {
		fmt.Println("ERROR/Student_show is error")
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 405,
			"msg":  "数据发送错误",
			"err":  err.Error(),
		})
		return
	}
	students := make([]model.Student, 0)
	for rows.Next() {
		var student model.Student
		rows.Scan(&student.StudentID, &student.Name, &student.Age, &student.Sex, &student.Department, &student.Studentclass)
		students = append(students, student)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("ERROR/Student_show is error")
	}
	ctx.JSON(
		200,
		gin.H{
			"code": 200,
			"data": students,
		},
	)
}

func StudentIDSearch(ctx *gin.Context) {
	json := make(map[string]interface{})
	ctx.BindJSON(&json)
	fmt.Println(json["studentID"])
	rows, err := db.Query("SELECT * FROM student where studentid = ? ", json["studentID"])
	if err != nil {
		fmt.Println("ERROR/Student_show is error")
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 405,
			"msg":  "数据发送错误",
			"err":  err.Error(),
		})
		return
	}
	students := make([]model.Student, 0)
	for rows.Next() {
		var student model.Student
		rows.Scan(&student.StudentID, &student.Name, &student.Age, &student.Sex, &student.Department, &student.Studentclass)
		students = append(students, student)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("ERROR/StudentIDSearch is error")
	} else {
		str := fmt.Sprintf("找不到数据：%s", json["studentID"])
		if len(students) <= 0 {
			ctx.JSON(
				200,
				gin.H{
					"code": 403,
					"msg":  str,
				},
			)
		} else {
			ctx.JSON(
				200,
				gin.H{
					"code": 200,
					"data": students,
				},
			)
		}
	}
}

func StudentDelete(ctx *gin.Context) {
	json := make(map[string]interface{})
	ctx.BindJSON(&json)
	fmt.Println(json["studentID"])

	rs, err := db.Exec("delete from student where studentid= ? ", json["studentID"])
	if err != nil {
		fmt.Println("ERROR/StudentDelete is err:s", err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 405,
			"msg":  "请求发送失败",
			"err":  err.Error(),
		})
	} else {
		returnNumber, err := rs.RowsAffected()
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "注销失败",
			})
		}
		if returnNumber == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "注销失败,找不到该学生",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":   200,
				"msg":    "删除成功",
				"number": returnNumber,
			})
		}
	}
}

func StudentAdd(ctx *gin.Context) {
	var jsoninfo model.Student
	if err := ctx.ShouldBindJSON(&jsoninfo); err != nil {
		fmt.Println("ERROR/student_add is err")
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 405,
			"msg":  "读取数据失败",
			"err":  err.Error(),
		})
		return
	} else {
		fmt.Println(jsoninfo)
		sqlStr := "INSERT INTO student (studentid, name, age, sex, department,studentclass) VALUES (?, ?, ?, ?, ?,?)"
		_, err := db.Exec(sqlStr, jsoninfo.StudentID, jsoninfo.Name, jsoninfo.Age, jsoninfo.Sex, jsoninfo.Department, jsoninfo.Studentclass)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 405,
				"msg":  "写入失败，学号已经存在",
				"err":  err.Error(),
			})
			//log.Fatalln(err)
			return
		} else {
			student := utils.GetStudentData(jsoninfo.StudentID)
			ctx.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "注册成功",
				"data": student,
			})
		}
	}
}

func StudentModify(ctx *gin.Context) {
	var jsoninfo model.Student
	if err := ctx.ShouldBindJSON(&jsoninfo); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 405,
			"msg":  "读取数据失败",
			"err":  err.Error(),
		})
		//log.Fatalln(err)
		return
	} else {
		sqlStr := "update student set name = ?,age= ?,sex=?,department=?,studentclass=? where studentid = ?"
		stmt, err := db.Exec(sqlStr, jsoninfo.Name, jsoninfo.Age, jsoninfo.Sex, jsoninfo.Department, jsoninfo.Studentclass, jsoninfo.StudentID)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "更新失败",
				"err":  err.Error(),
			})
		} else {
			returnNumber, err := stmt.RowsAffected()
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 403,
					"msg":  "修改失败，写入数据报错",
					"err":  err.Error(),
				})
			}
			if returnNumber == 0 {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 403,
					"msg":  "修改失败,找不到该账号或密码错误",
					"data": jsoninfo,
				})
			} else {
				student := utils.GetStudentData(jsoninfo.StudentID)
				ctx.JSON(http.StatusOK, gin.H{
					"code": 200,
					"msg":  "修改成功",
					"data": student,
				})
			}
		}
	}
}
