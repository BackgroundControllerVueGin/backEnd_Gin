package model

type Student struct {
	StudentID  string `json:"studentID"`
	Name       string `json:"name"`
	Age        string `json:"age"`
	Sex        string `json:"sex"`
	Department string `json:"department"`
}
