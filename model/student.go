package model

//type Student struct {
//	StudentID    string `json:"studentID"`
//	Name         string `json:"name"`
//	Age          string `json:"age"`
//	Sex          string `json:"sex"`
//	Department   string `json:"department"`
//	Studentclass string `json:"studentclass"`
//}

type Student struct {
	Studentid    int    `gorm:"column:studentid" db:"studentid" json:"studentid"`
	Name         string `gorm:"column:name" db:"name" json:"name"`
	Age          int    `gorm:"column:age" db:"age" json:"age"`
	Sex          string `gorm:"column:sex" db:"sex" json:"sex"`
	Department   string `gorm:"column:department" db:"department" json:"department"`
	Studentclass string `gorm:"column:studentclass" db:"studentclass" json:"studentclass"`
}
