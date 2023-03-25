package model

type Student struct {
	StudentID  int64  `json:"studentID"`
	Name       string `json:"name"`
	Age        int64  `json:"age"`
	Sex        string `json:"sex"`
	Department string `json:"department"`
}
