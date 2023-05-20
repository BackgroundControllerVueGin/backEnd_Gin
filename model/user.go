package model

//type User struct {
//	No       string `json:"no"`
//	Username string `json:"username"`
//	Password string `json:"password"`
//}

type User struct {
	No       int    `gorm:"column:No" db:"No" json:"No"`
	Username string `gorm:"column:username" db:"username" json:"username"`
	Password string `gorm:"column:password" db:"password" json:"password"`
}
