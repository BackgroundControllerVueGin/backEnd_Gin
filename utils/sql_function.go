package utils

import (
	"backEnd_Gin/common"
	"backEnd_Gin/model"
	"fmt"
)

var db = common.GetDB()

func GetStudentData(studentID string) model.Student {
	sqlStr := "select * from student where studentid = ? ;"
	row := db.QueryRow(sqlStr, studentID)
	var student model.Student
	row.Scan(&student.StudentID, &student.Name, &student.Age, &student.Sex, &student.Department, &student.Studentclass)
	return student
}

func GetStudentCount() int {
	sqlStr := "select count(*) from student;"
	row := db.QueryRow(sqlStr)
	var number int
	fmt.Println(row.Scan(&number))
	return number
}
