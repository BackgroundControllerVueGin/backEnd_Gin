package model

import (
	"fmt"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&Student{}, &User{})

	if err != nil {
		fmt.Println(err)
		return
	}
}
