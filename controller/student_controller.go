package controller

import (
	"backEnd_Gin/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StudentShow(ctx *gin.Context) {
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
		rows.Scan(&student.StudentID, &student.Name, &student.Age, &student.Sex, &student.Department)
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
