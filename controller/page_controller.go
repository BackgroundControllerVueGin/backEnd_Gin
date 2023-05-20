package controller

//
//import (
//	"backEnd_Gin/model"
//	"backEnd_Gin/utils"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func NextPage(ctx *gin.Context) {
//	fmt.Println("Student_show is running")
//	var jsoninfo model.NextPage
//	if err := ctx.ShouldBindJSON(&jsoninfo); err != nil {
//		fmt.Println(jsoninfo)
//		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
//			"code": 405,
//			"msg":  "读取数据失败",
//			"err":  err.Error(),
//		})
//	} else {
//		total_quantity := utils.GetStudentCount()
//		data_number := jsoninfo.DataNumber
//		next := (jsoninfo.CurrentPage - 1) * jsoninfo.DataNumber
//		fmt.Println("%s and %s", next, data_number)
//		rows, err := db.Query("SELECT * FROM student limit ?,?", next, data_number)
//		if err != nil {
//			fmt.Println("ERROR/Student_show is error")
//			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
//				"code": 405,
//				"msg":  "数据发送错误",
//				"err":  err.Error(),
//			})
//			return
//		}
//		students := make([]model.Student, 0)
//		for rows.Next() {
//			var student model.Student
//			rows.Scan(&student.StudentID, &student.Name, &student.Age, &student.Sex, &student.Department, &student.Studentclass)
//			students = append(students, student)
//		}
//		if err = rows.Err(); err != nil {
//			fmt.Println("ERROR/Student_show is error")
//		}
//		ctx.JSON(
//			200,
//			gin.H{
//				"code":  200,
//				"total": total_quantity,
//				"data":  students,
//			},
//		)
//	}
//}
